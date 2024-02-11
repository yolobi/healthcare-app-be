package operationalusecaseinterface

import (
	"context"
	"healthcare-capt-america/entities/models"
)

type OperationalUsecase interface {
	AddOperationalDays(ctx context.Context, operationals []models.Operational) ([]models.Operational, error)
	GetPharmacyOperationalsDay(ctx context.Context, pharmacyId uint64) ([]models.Pharmacy, error)
	DeleteOperationalDay(ctx context.Context, operationalId uint64) error
}
