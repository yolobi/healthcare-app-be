package operationalrepositoryinterface

import (
	"context"
	"healthcare-capt-america/entities/models"
)

type OperationalRepository interface {
	SaveMultiple(ctx context.Context, operational []models.Operational) ([]models.Operational, error)
	FindByPharmacyID(ctx context.Context, pharmacyId uint64) ([]models.Operational, error)
	Delete(ctx context.Context, operationalId uint64) error
}
