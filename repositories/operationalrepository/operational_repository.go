package operationalrepository

import (
	"context"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/interfaces/repositories/operationalrepositoryinterface"

	"gorm.io/gorm"
)

type operationalRepository struct {
	db *gorm.DB
}

func (repo *operationalRepository) Delete(ctx context.Context, operationalId uint64) error {
	return repo.db.WithContext(ctx).Table("operationals").Where("id = ?", operationalId).Delete(operationalId).Error
}

func (repo *operationalRepository) FindByPharmacyID(ctx context.Context, pharmacyId uint64) ([]models.Operational, error) {
	var operationals []models.Operational
	err := repo.db.WithContext(ctx).
		Preload("Pharmacy.Address").
		Preload("Pharmacy.Address.City").
		Preload("Pharmacy.Address.Province").Where("pharmacy_id = ?", pharmacyId).
		Find(&operationals).Error
	if err != nil {
		return nil, err
	}
	return operationals, nil
}

func (repo *operationalRepository) SaveMultiple(ctx context.Context, operationals []models.Operational) ([]models.Operational, error) {
	err := repo.db.WithContext(ctx).Create(&operationals).Error
	if err != nil {
		return nil, err
	}
	return operationals, nil
}

func NewOperationalRepository(db *gorm.DB) *operationalRepository {
	return &operationalRepository{db: db}
}

var _ operationalrepositoryinterface.OperationalRepository = &operationalRepository{}
