package orderrepository

import (
	"context"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/dto/responses/checkoutres"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/repositories/orderrepositoryinterface"
	"healthcare-capt-america/utils"
	"time"
)

type orderRepository struct {
	db *gorm.DB
}

func (repo *orderRepository) FindByPaymentID(ctx context.Context, paymentId uint64) (*models.Order, error) {
	var order *models.Order
	err := repo.db.WithContext(ctx).Model(&models.Order{}).Where("payment_id = ? ", paymentId).First(&order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (repo *orderRepository) FindByID(ctx context.Context, id uint64) (*models.Order, error) {
	var order *models.Order
	err := repo.db.WithContext(ctx).Model(&models.Order{}).
		Preload("OrderDetails.PharmacyDrug.Drug").
		Preload("User").
		Preload("Address.Province").
		Preload("Address.City").
		Preload("Address.User").
		Preload("Pharmacy.AdminPharmacies").
		Preload("Pharmacy.Address.Province").
		Preload("Pharmacy.Address.City").
		Preload("Payment").
		Where("id = ?", id).
		First(&order).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, apperror.NewServerError(err)
	}
	return order, nil
}

func (repo *orderRepository) Delete(ctx context.Context, id uint64) error {
	return repo.db.WithContext(ctx).Model(&models.Order{}).Where("id = ?", id).Delete(id).Error
}

func (repo *orderRepository) FindAll(ctx context.Context, qry *requests.GlobalQuery) ([]*models.Order, *responses.Pagination, error) {
	orders := make([]*models.Order, 0)
	limit, offset := qry.GetPagination()
	preload := repo.db.WithContext(ctx).Model(&models.Order{}).
		Preload("OrderDetails.PharmacyDrug.Drug").
		Preload("User").
		Preload("Address.Province").
		Preload("Address.City").
		Preload("Address.User").
		Preload("Pharmacy.Address").
		Joins("Pharmacy.AdminPharmacies").
		Preload("Pharmacy.Address.Province").
		Preload("Pharmacy.Address.City").
		Preload("Payment")
	for _, condition := range qry.Conditions {
		sql := fmt.Sprintf("%s %s ?", condition.Field, condition.Operation)
		preload.Where(sql, condition.Value)
	}
	searchQuery := fmt.Sprintf("order_status ILIKE %s", qry.Search)
	preload.Where(searchQuery).Order(qry.OrderedBy)
	err := preload.Limit(limit).Offset(offset).Find(&orders).Error
	if err != nil {
		return nil, nil, apperror.NewServerError(err, "Database Error")
	}
	var totalItems, totalPages int64
	preload.Offset(-1).Count(&totalItems)
	totalPages = totalItems / int64(qry.PerPage)
	if totalItems%int64(qry.PerPage) != 0 {
		totalPages += 1
	}
	return orders, &responses.Pagination{
		TotalItems:  totalItems,
		TotalPages:  totalPages,
		CurrentPage: int64(qry.Page),
	}, nil
}

func (repo *orderRepository) CreateOrder(ctx context.Context, order *models.Order) (*models.Order, error) {
	err := repo.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		//1. create payment
		payment := &models.Payment{}
		payment.Status = enums.NotUploaded
		payment, err := utils.SaveQuery[models.Payment](ctx, tx, payment, enums.Create)
		if err != nil {
			return err
		}
		order.PaymentId = payment.ID
		//3. check cart quantity stock with pharmacy
		var carts []models.Cart
		err = tx.WithContext(ctx).Model(&models.Cart{}).Where("user_id = ?", order.UserId).Find(&carts).Error
		if err != nil {
			return err
		}
		productIdsInCart := []uint64{}
		for _, cart := range carts {
			productIdsInCart = append(productIdsInCart, cart.DrugId)
		}
		//3.1 locking for update product stock
		var pharmacyDrugs []models.PharmacyDrug
		err = tx.WithContext(ctx).Model(&models.PharmacyDrug{}).Where("pharmacy_id = ?", order.PharmacyId).Where("drug_id IN (?)", productIdsInCart).Find(&pharmacyDrugs).Error
		if err != nil {
			return err
		}
		for _, cart := range carts {
			err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).WithContext(ctx).Model(&models.PharmacyDrug{}).Where("pharmacy_id = ? AND drug_id = ?", order.PharmacyId, cart.DrugId).Update("stock", gorm.Expr("stock - ?", cart.Quantity)).Error
			if err != nil {
				return err
			}
		}
		//3.2.1 stock mutation(IF NEEDED)
		type DrugQuantity struct {
			DrugId   uint64
			Stock    int64
			Quantity int64
		}
		var drugQuantities []DrugQuantity
		err = tx.WithContext(ctx).Raw(enums.GetProductStockLessThanCartQuantity, order.PharmacyId, order.UserId).Scan(&drugQuantities).Error
		if err != nil {
			return err
		}

		if drugQuantities != nil || len(drugQuantities) != 0 {
			for _, drugQuantity := range drugQuantities {
				pharmacyId := 0
				//3.2.2 find pharmacy for stock mutation, select pharmacy with stock
				err = tx.WithContext(ctx).Raw(enums.GetPharmacyThatHasMoreStock, order.PharmacyId, drugQuantity.DrugId, drugQuantity.DrugId).Scan(&pharmacyId).Error
				if err != nil {
					return err
				}
				//3.2.3 create stock mutation
				var toPharmacy models.PharmacyDrug
				var fromPharmacy models.PharmacyDrug
				toPharmacyId := order.PharmacyId
				fromPharmacyId := pharmacyId
				tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("pharmacy_id = ? AND drug_id = ?", toPharmacyId, drugQuantity.DrugId).First(&toPharmacy)
				tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("pharmacy_id = ? AND drug_id = ?", fromPharmacyId, drugQuantity.DrugId).First(&fromPharmacy)
				err = tx.WithContext(ctx).Table("pharmacy_drugs").Where("pharmacy_id = ? AND drug_id = ?", fromPharmacy.PharmacyId, drugQuantity.DrugId).Update("stock", gorm.Expr("stock + ?", drugQuantity.Stock)).Scan(&fromPharmacy).Error
				if err != nil {
					return err
				}
				err = tx.WithContext(ctx).Table("pharmacy_drugs").Where("pharmacy_id = ? AND drug_id = ?", toPharmacy.PharmacyId, drugQuantity.DrugId).Update("stock", gorm.Expr("stock - ?", drugQuantity.Stock)).Scan(&toPharmacy).Error
				if err != nil {
					return err
				}
				//3.2.4 create stock mutation
				err = tx.WithContext(ctx).Exec(enums.InsertIntoStockMutation, fromPharmacyId, toPharmacyId, drugQuantity.DrugId, drugQuantity.Stock*(-1), enums.Accepted).Error
				if err != nil {
					return err
				}
				//3.2.5 create journal
				err = tx.WithContext(ctx).Exec(enums.InsertIntoJournals, drugQuantity.DrugId, fromPharmacyId, toPharmacyId, "Accepted", drugQuantity.Stock*(-1), fromPharmacy.Stock).Error
				if err != nil {
					return err
				}

			}
		}
		//4. create order
		order, err = utils.SaveQuery[models.Order](ctx, tx, order, enums.Create)
		if err != nil {
			return err
		}
		//5. create order detail
		var orderDetails []models.OrderDetail
		var cartCheckout []checkoutres.CartResponse
		err = repo.db.WithContext(ctx).Raw(enums.GetCartForCheckout, order.PharmacyId, order.UserId).Scan(&cartCheckout).Error
		if err != nil {
			return err
		}
		for _, cart := range cartCheckout {
			var updatedPharmacyDrug *models.PharmacyDrug
			err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).Model(&models.PharmacyDrug{}).Where("id = ?", cart.PharmacyDrugID).First(&updatedPharmacyDrug).Error
			if err != nil {
				return err
			}

			orderDetail := &models.OrderDetail{
				OrderId:        order.ID,
				PharmacyDrugId: cart.PharmacyDrugID,
				Quantity:       cart.Quantity,
				Amount:         decimal.NewFromFloat(cart.Price * float64(cart.Quantity)),
			}
			if err != nil {
				return err
			}
			orderDetails = append(orderDetails, *orderDetail)

		}
		err = tx.WithContext(ctx).Model(&models.OrderDetail{}).Create(&orderDetails).Error
		if err != nil {
			return err
		}
		//6. delete carts
		err = tx.WithContext(ctx).Unscoped().Delete(&carts).Error
		if err != nil {
			return err
		}
		return nil
	})
	return order, err
}

func (repo *orderRepository) Update(ctx context.Context, order *models.Order) error {
	err := repo.db.WithContext(ctx).
		Model(&models.Order{}).
		Where("id", order.ID).
		Update("order_status", order.OrderStatus).
		Error
	if err != nil {
		return apperror.NewServerError(err)
	}
	return nil
}

func (repo *orderRepository) SetConfirmed(ctx context.Context) error {
	err := repo.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		beforeTime := time.Now().Add(time.Duration(-1*24*7) * time.Hour).Format(time.RFC3339)
		err := tx.Model(&models.Order{}).
			Where("order_status", enums.Sent).
			Where("updated_at < ?", beforeTime).
			Update("order_status", enums.OrderConfirmed).Error
		return err
	})
	if err != nil {
		return apperror.NewServerError(err)
	}
	return nil
}

func (repo *orderRepository) SetCanceled(ctx context.Context) error {
	beforeTime := time.Now().Add(time.Duration(-24) * time.Hour).Format(time.RFC3339)
	err := repo.db.WithContext(ctx).
		Model(&models.Order{}).
		Where("order_status", enums.WaitingForPayment).
		Where("updated_at < ?", beforeTime).
		Update("order_status", enums.OrderCanceled).Error
	if err != nil {
		return apperror.NewServerError(err)
	}
	return nil
}

func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{db: db}
}

var _ orderrepositoryinterface.OrderRepository = &orderRepository{}
