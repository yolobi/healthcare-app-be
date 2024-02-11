package pharmacyrepository

import (
	"context"
	"errors"
	"fmt"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/repositories/pharmacyrepositoryinterface"
	"healthcare-capt-america/utils"

	"gorm.io/gorm"
)

type pharmacyRepository struct {
	db *gorm.DB
}

func (repo *pharmacyRepository) FindByAdminPharmacyId(ctx context.Context, adminPharmacyId uint64, qry *requests.GlobalQuery) ([]models.Pharmacy, *responses.Pagination, error) {
	var pharmacies = make([]models.Pharmacy, 0)
	limit, offset := qry.GetPagination()
	preload := repo.db.WithContext(ctx).Model(&models.Pharmacy{})
	preload.Preload("Address.City").
		Preload("Address.Province")
	for _, condition := range qry.Conditions {
		sql := fmt.Sprintf("%s %s ?", condition.Field, condition.Operation)
		preload.Where(sql, condition.Value)
	}
	searchQuery := fmt.Sprintf(`
		name ILIKE %s OR 
		pharmaciest_name ILIKE %s OR 
		license_number ILIKE %s OR 
		phone_number ILIKE %s
	`, qry.Search, qry.Search, qry.Search, qry.Search)
	preload.Where(searchQuery).Where("admin_pharmacy_id = ?", adminPharmacyId)
	err := preload.Limit(limit).Offset(offset).Order(qry.OrderedBy).Find(&pharmacies).Error
	if err != nil {
		return nil, nil, err
	}
	var totalItems, totalPages int64
	preload.Offset(-1).Count(&totalItems)
	totalPages = totalItems / int64(qry.PerPage)
	if totalItems%int64(qry.PerPage) != 0 {
		totalPages += 1
	}
	return pharmacies, &responses.Pagination{
		TotalItems:  totalItems,
		TotalPages:  totalPages,
		CurrentPage: int64(qry.Page),
	}, nil
}

func (repo *pharmacyRepository) Delete(ctx context.Context, pharmacy *models.Pharmacy) error {
	repo.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		pharmacyDrug := &models.PharmacyDrug{}
		err := repo.db.Delete(&pharmacy).Error
		if err != nil {
			return apperror.NewServerError(err)
		}
		err = repo.db.Model(&pharmacyDrug).Where("pharmacy_id = ?", pharmacy.ID).Error
		if err != nil {
			return apperror.NewServerError(err)
		}
		return nil
	})
	return nil
}

func (repo *pharmacyRepository) FindById(ctx context.Context, id uint64) (*models.Pharmacy, error) {
	var pharmacy *models.Pharmacy
	err := repo.db.
		Preload("Operationals").
		Preload("Address.City").
		Preload("AdminPharmacies.User").
		Preload("Address.Province").Where("id = ?", id).First(&pharmacy).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, apperror.NewServerError(err)
	}
	return pharmacy, nil
}

func (repo *pharmacyRepository) Update(ctx context.Context, pharmacy *models.Pharmacy) (*models.Pharmacy, error) {
	return utils.SaveQuery[models.Pharmacy](ctx, repo.db, pharmacy, enums.Update)
}

func (repo *pharmacyRepository) FindAll(ctx context.Context, qry *requests.GlobalQuery) ([]models.Pharmacy, *responses.Pagination, error) {
	var pharmacies = make([]models.Pharmacy, 0)
	limit, offset := qry.GetPagination()
	preload := repo.db.WithContext(ctx).Model(&models.Pharmacy{})
	preload.Preload("Address.City").
		Preload("Address.Province")
	for _, condition := range qry.Conditions {
		sql := fmt.Sprintf("%s %s ?", condition.Field, condition.Operation)
		preload.Where(sql, condition.Value)
	}
	searchQuery := fmt.Sprintf(`
		name ILIKE %s OR 
		pharmaciest_name ILIKE %s OR 
		license_number ILIKE %s OR 
		phone_number ILIKE %s
	`, qry.Search, qry.Search, qry.Search, qry.Search)
	preload.Where(searchQuery)
	err := preload.Limit(limit).Offset(offset).Find(&pharmacies).Error
	if err != nil {
		return nil, nil, err
	}
	var totalItems, totalPages int64
	preload.Offset(-1).Count(&totalItems)
	totalPages = totalItems / int64(qry.PerPage)
	if totalItems%int64(qry.PerPage) != 0 {
		totalPages += 1
	}
	return pharmacies, &responses.Pagination{
		TotalItems:  totalItems,
		TotalPages:  totalPages,
		CurrentPage: int64(qry.Page),
	}, nil
}

func (repo *pharmacyRepository) AddPharmacy(ctx context.Context, address *models.Address, pharmacy *models.Pharmacy) (*models.Pharmacy, error) {
	err := repo.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		address, err := utils.SaveQuery[models.Address](ctx, tx, address, enums.Create)
		if err != nil {
			return err
		}

		pharmacy.AddressID = address.ID
		pharmacy, err = utils.SaveQuery[models.Pharmacy](ctx, tx, pharmacy, enums.Create)
		if err != nil {
			return err
		}

		return nil
	})
	return pharmacy, err
}

func NewPharmacyRepository(db *gorm.DB) *pharmacyRepository {
	return &pharmacyRepository{db: db}
}

var _ pharmacyrepositoryinterface.PharmacyRepository = &pharmacyRepository{}
