package pharmacyproductrepository

import (
	"context"
	"errors"
	"fmt"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/repositories/pharmacyproductrepositoryinterface"
	"healthcare-capt-america/utils"

	"github.com/jackc/pgconn"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type pharmacyDrugRepository struct {
	db *gorm.DB
}

func (repo *pharmacyDrugRepository) FindPharmacyIDCheckout(ctx context.Context, address *models.Address, carts []*models.Cart) (uint64, error) {
	var pharmacyID uint64
	var drugIDs []uint64
	for _, cart := range carts {
		drugIDs = append(drugIDs, cart.DrugId)
	}
	err := repo.db.WithContext(ctx).
		Model(&models.Pharmacy{}).
		Raw(enums.QueryToGetPharmacyThatHasProducts, drugIDs, address.Longtitude, address.Latitude).Scan(&pharmacyID).Error
	if err != nil {
		return 0, err
	}
	return pharmacyID, nil
}

func (repo *pharmacyDrugRepository) FindByAdminPharmacyID(ctx context.Context, qry *requests.GlobalQuery, adminId uint64) ([]models.PharmacyDrug, *responses.Pagination, error) {
	var results = make([]models.PharmacyDrug, 0)
	limit, offset := qry.GetPagination()
	// query := enums.GetAdminPharmacyProducts
	preload := repo.db.WithContext(ctx).Model(&models.PharmacyDrug{})
	preload.Preload("Pharmacy.Address.Province").Preload("Pharmacy.Address.City").
		Preload("Drug.Manufacture").Preload("Drug.Form").Preload("Drug.Category").Preload("Drug").
		Joins("left join pharmacies on pharmacies.id = pharmacy_drugs.pharmacy_id")
	// Raw(query).
	for _, condition := range qry.Conditions {
		sql := fmt.Sprintf("%s %s ?", condition.Field, condition.Operation)
		preload.Where(sql, condition.Value)
	}
	preload.Where("pharmacies.admin_pharmacy_id = ?", adminId)
	err := preload.Limit(limit).Offset(offset).Order(qry.OrderedBy).Find(&results).Error
	if err != nil {
		return nil, nil, err
	}
	var totalItems, totalPages int64
	preload.Offset(-1).Count(&totalItems)
	totalPages = totalItems / int64(qry.PerPage)
	if totalItems%int64(qry.PerPage) != 0 {
		totalPages += 1
	}
	return results, &responses.Pagination{
		TotalItems:  totalItems,
		TotalPages:  totalPages,
		CurrentPage: int64(qry.Page),
	}, nil
}

func (repo *pharmacyDrugRepository) CheckDrugIfExists(ctx context.Context, drugId uint64) (bool, error) {
	var exist bool
	err := repo.db.WithContext(ctx).Raw(enums.QueryCheckProductExistsInPharmacyProduct, drugId).Scan(&exist).Error
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (repo *pharmacyDrugRepository) CountDrugStock(ctx context.Context, drugId uint64) (uint64, error) {
	var totalStock uint64
	err := repo.db.WithContext(ctx).Exec("SELECT SUM(stock) FROM pharmacy_drugs WHERE id = ?", drugId).Scan(&totalStock).Error
	if err != nil {
		return 0, err
	}
	return totalStock, nil
}

func NewPharmacyDrugRepository(db *gorm.DB) *pharmacyDrugRepository {
	return &pharmacyDrugRepository{db: db}
}

func (repo *pharmacyDrugRepository) FindByPharmacyID(ctx context.Context, qry *requests.GlobalQuery, pharmacy_id uint64) ([]models.PharmacyDrug, *responses.Pagination, error) {
	pharmacyDrugs := make([]models.PharmacyDrug, 0)
	limit, offset := qry.GetPagination()
	preload := repo.db.WithContext(ctx).Model(&models.PharmacyDrug{}).
		Preload("Drug.Form").Preload("Drug.Manufacture").Preload("Drug.Category").Preload("Pharmacy").
		Where("pharmacy_id = ?", pharmacy_id)
	for _, condition := range qry.Conditions {
		sql := fmt.Sprintf("%s %s ?", condition.Field, condition.Operation)
		preload.Where(sql, condition.Value)
	}
	err := preload.Limit(limit).Offset(offset).Order(qry.OrderedBy).Find(&pharmacyDrugs).Error
	if err != nil {
		return nil, nil, err
	}
	var totalItems, totalPages int64
	preload.Offset(-1).Count(&totalItems)
	totalPages = totalItems / int64(qry.PerPage)
	if totalItems%int64(qry.PerPage) != 0 {
		totalPages += 1
	}
	return pharmacyDrugs, &responses.Pagination{
		TotalItems:  totalItems,
		TotalPages:  totalPages,
		CurrentPage: int64(qry.Page),
	}, nil
}

func (repo *pharmacyDrugRepository) FindByID(ctx context.Context, pharmacy_id uint64, drug_id uint64) (pharmacyDrug *models.PharmacyDrug, err error) {
	err = repo.db.WithContext(ctx).Model(&pharmacyDrug).Where("pharmacy_id = ? AND drug_id = ?", pharmacy_id, drug_id).Preload("Drug.Form").Preload("Drug.Manufacture").Preload("Drug.Category").Preload("Pharmacy").First(&pharmacyDrug).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, apperror.NewServerError(err)
	}
	return pharmacyDrug, nil
}

func (repo *pharmacyDrugRepository) Update(ctx context.Context, pharmacyDrug *models.PharmacyDrug) (*models.PharmacyDrug, error) {
	err := repo.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var old *models.PharmacyDrug
		journal := &models.Journal{
			DrugId:       pharmacyDrug.DrugId,
			ToPharmacyId: pharmacyDrug.PharmacyId,
			Quantity:     pharmacyDrug.Stock,
		}
		tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("pharmacy_id = ? AND drug_id = ?", pharmacyDrug.PharmacyId, pharmacyDrug.DrugId).First(&old)
		quantityChanged := old.Stock - pharmacyDrug.Stock
		tx.Model(&pharmacyDrug).Preload("Drug").Preload("Pharmacy").Where("pharmacy_id = ? AND drug_id = ?", pharmacyDrug.PharmacyId, pharmacyDrug.DrugId).Updates(&pharmacyDrug)
		if quantityChanged > 0 {
			journal.Status = "Addition"
		} else {
			journal.Status = "Subtraction"
			journal.Quantity *= -1
		}
		tx.Create(&journal)
		return nil
	})
	return pharmacyDrug, err
}

func (repo *pharmacyDrugRepository) Save(ctx context.Context, pharmacyDrug *models.PharmacyDrug) (*models.PharmacyDrug, error) {
	err := repo.db.Model(&pharmacyDrug).Preload("Drug").Preload("Pharmacy").Create(&pharmacyDrug).Error
	var duplicateEntryError = &pgconn.PgError{Code: "23505"}
	if errors.As(err, &duplicateEntryError) {
		return nil, apperror.NewClientError(errors.New("product already exist on the pharmacy"))
	}
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	journal := &models.Journal{
		DrugId:       pharmacyDrug.DrugId,
		ToPharmacyId: pharmacyDrug.PharmacyId,
		Quantity:     pharmacyDrug.Stock,
		Status:       "Addition",
	}
	err = repo.db.Create(&journal).Error
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	return pharmacyDrug, nil
}

func (repo *pharmacyDrugRepository) Delete(ctx context.Context, pharmacyDrug *models.PharmacyDrug) error {
	return utils.Delete[models.PharmacyDrug](ctx, repo.db, pharmacyDrug)
}

var _ pharmacyproductrepositoryinterface.PharmacyDrugRepository = &pharmacyDrugRepository{}
