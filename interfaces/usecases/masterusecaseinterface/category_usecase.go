package masterusecaseinterface

import (
	"context"
	"healthcare-capt-america/entities/models"
)

type CategoryUsecase interface {
	GetAllCategories(ctx context.Context) ([]models.Category, error)
}
