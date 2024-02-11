package masterrepository

import (
	"context"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/repositories/masterrepositoryinterface"
	"healthcare-capt-america/utils"

	"gorm.io/gorm"
)

type shipmentRepository struct {
	db *gorm.DB
}

func (repo *shipmentRepository) Save(ctx context.Context, shipment *models.Shipment) (*models.Shipment, error) {
	return utils.SaveQuery[models.Shipment](ctx, repo.db, shipment, enums.Create)
}

func (sr *shipmentRepository) FindAll(ctx context.Context) ([]*models.Shipment, error) {
	shipments := make([]*models.Shipment, 0)
	err := sr.db.WithContext(ctx).
		Model(&models.Shipment{}).
		Find(&shipments).Error
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	return shipments, nil
}

func (sr *shipmentRepository) FindOfficial(ctx context.Context) ([]*models.Shipment, error) {
	shipments := make([]*models.Shipment, 0)
	err := sr.db.WithContext(ctx).Model(&models.Shipment{}).
		Where("name = 'instant' OR name = 'same day'").
		Find(&shipments).Error
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	return shipments, nil
}

func NewShipmentRepository(db *gorm.DB) *shipmentRepository {
	return &shipmentRepository{db: db}
}

var _ masterrepositoryinterface.ShipmentRepository = &shipmentRepository{}
