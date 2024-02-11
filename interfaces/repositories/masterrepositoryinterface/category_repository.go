package masterrepositoryinterface

import (
	"context"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/models"
)

type CategoryRepository interface {
	Save(context.Context, *models.Category) (*models.Category, error)
	FindByID(context.Context, uint64) (*models.Category, error)
	Update(context.Context, *models.Category) (*models.Category, error)
	FindAll(context.Context, *requests.GlobalQuery) ([]models.Category, *responses.Pagination, error)
	Delete(context.Context, *models.Category) error
	Find(ctx context.Context) ([]models.Category, error)
}
