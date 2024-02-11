package masterusecaseinterface

import (
	"context"
	"healthcare-capt-america/entities/models"
)

type ManufactureUsecase interface {
	GetAllManufactures(ctx context.Context) ([]models.Manufacture, error)
}
