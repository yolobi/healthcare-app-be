package masterrepositoryinterface

import (
	"context"
	"healthcare-capt-america/entities/models"
)

type ManufactureRepository interface {
	Save(ctx context.Context, manufacture *models.Manufacture) (*models.Manufacture, error)
	FindByID(ctx context.Context, manufacture_id uint64) (*models.Manufacture, error)
	FindAll(ctx context.Context) ([]models.Manufacture, error)
}
