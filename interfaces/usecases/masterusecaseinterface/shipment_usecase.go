package masterusecaseinterface

import (
	"context"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/entities/models/transaction"
)

type ShipmentUsecase interface {
	FindAll(ctx context.Context) ([]*models.Shipment, error)
	CalculateDistance(ctx context.Context, address_id uint64, pharmacy_id uint64, weight int64) ([]*transaction.ShipmentFee, error)
}
