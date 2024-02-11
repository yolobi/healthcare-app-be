package stockmutationrepository

import (
	"context"
	"errors"
	"fmt"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/interfaces/repositories/stockmutationrepositoryinterface"
	"healthcare-capt-america/utils"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type stockMutationRepository struct {
	db *gorm.DB
}

func NewStockMutationRepository(db *gorm.DB) *stockMutationRepository {
	return &stockMutationRepository{db: db}
}

func (repo *stockMutationRepository) FindByID(ctx context.Context, stockMutation_id uint64) (*models.StockMutation, error) {
	var stocks_mutations *models.StockMutation
	err := repo.db.Model(&stocks_mutations).Preload("Drug").Preload("FromPharmacy.Address").Preload("ToPharmacy.Address").Preload("StatusMutation").Where("id = ?", stockMutation_id).First(&stocks_mutations).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, apperror.NewServerError(err)
	}
	return stocks_mutations, nil
}

func (repo *stockMutationRepository) AcceptRequestStockMutation(ctx context.Context, stockMutation *models.StockMutation, journal *models.Journal) (*models.StockMutation, error) {
	err := repo.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var toPharmacy models.PharmacyDrug
		var fromPharmacy models.PharmacyDrug
		tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("pharmacy_id = ? AND drug_id = ?", stockMutation.ToPharmacyId, stockMutation.DrugId).First(&toPharmacy)
		tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("pharmacy_id = ? AND drug_id = ?", stockMutation.FromPharmacyId, stockMutation.DrugId).First(&fromPharmacy)
		if fromPharmacy.Stock-stockMutation.Quantity < 0 {
			return errors.New("insufficient stock")
		}
		err := tx.Create(&journal).Error
		if err != nil {
			return err
		}
		fromPharmacy.Stock = fromPharmacy.Stock - stockMutation.Quantity
		tx.Save(&fromPharmacy)
		toPharmacy.Stock = toPharmacy.Stock + stockMutation.Quantity
		tx.Save(&toPharmacy)
		stockMutation.UpdatedAt = time.Now()
		tx.Preload(clause.Associations).Save(&stockMutation)
		return nil
	})
	return stockMutation, err
}

func (repo *stockMutationRepository) CreateRequestStockMutation(ctx context.Context, stockMutation *models.StockMutation) (*models.StockMutation, error) {
	err := repo.db.WithContext(ctx).
		Preload(clause.Associations).
		Create(&stockMutation).First(&stockMutation).Error
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	return stockMutation, nil
}

func (repo *stockMutationRepository) Delete(ctx context.Context, stockMutation *models.StockMutation) error {
	return utils.Delete[models.StockMutation](ctx, repo.db, stockMutation)
}

func (repo *stockMutationRepository) FindAll(ctx context.Context, qry *requests.GlobalQuery, uId uint64, action string) ([]*models.StockMutation, *responses.Pagination, error) {
	// action = (strings.Split(action, ","))[0]
	stockMutations := make([]*models.StockMutation, 0)
	limit, offset := qry.GetPagination()
	preload := repo.db.WithContext(ctx).Model(&stockMutations)
	if action == "from" {
		preload.Joins("JOIN pharmacies ON pharmacies.id=stock_mutations.from_pharmacy_id AND pharmacies.admin_pharmacy_id = ?", uId)
	} else if action == "to" {
		preload.Joins("JOIN pharmacies ON pharmacies.id=stock_mutations.to_pharmacy_id AND pharmacies.admin_pharmacy_id = ?", uId)
	} else {
		preload.Joins("JOIN pharmacies ON pharmacies.id=stock_mutations.from_pharmacy_id OR pharmacies.id=stock_mutations.to_pharmacy_id AND pharmacies.admin_pharmacy_id = ?", uId)
	}
	for _, condition := range qry.Conditions {
		sql := fmt.Sprintf("%s %s ?", condition.Field, condition.Operation)
		preload.Where(sql, condition.Value)
	}
	order := fmt.Sprintf("stock_mutations.%s", qry.OrderedBy)
	preload.Offset(offset).Limit(limit).Order(order)
	// err := preload.Preload(clause.Associations).Distinct("stock_mutations.id,stock_mutations.drug_id,stock_mutations.to_pharmacy_id,stock_mutations.from_pharmacy_id,stock_mutations.updated_at,pharmacies.address_id").Find(&stockMutations).Error
	err := preload.Preload(clause.Associations).Find(&stockMutations).Error
	if err != nil {
		return nil, nil, apperror.NewServerError(err, "Database Error")
	}
	var totalItems, totalPages int64
	preload.Offset(-1).Count(&totalItems)
	totalPages = totalItems / int64(qry.PerPage)
	if totalItems%int64(qry.PerPage) != 0 {
		totalPages += 1
	}
	return stockMutations, &responses.Pagination{
		TotalItems:  totalItems,
		TotalPages:  totalPages,
		CurrentPage: int64(qry.Page),
	}, nil
}

func (repo *stockMutationRepository) Update(ctx context.Context, stockMutation *models.StockMutation) (*models.StockMutation, error) {
	err := repo.db.WithContext(ctx).Model(&stockMutation).
		Updates(&stockMutation).Error
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	repo.db.WithContext(ctx).Preload(clause.Associations).First(&stockMutation)
	return stockMutation, nil
}

var _ stockmutationrepositoryinterface.StockMutationRepository = &stockMutationRepository{}
