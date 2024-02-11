package masterrepositoryinterface

import (
	"context"
	"healthcare-capt-america/entities/models"
)

type ShipmentRepository interface {
	Save(ctx context.Context, shipment *models.Shipment) (*models.Shipment, error)
	FindAll(ctx context.Context) ([]*models.Shipment, error)
	FindOfficial(ctx context.Context) ([]*models.Shipment, error)
}
