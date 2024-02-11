package productcategoryusecaseinterface

import (
	"context"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/models"
)

type CategoryUsecase interface {
	CreateCategory(context.Context, *models.Category) (*models.Category, error)
	EditCategory(context.Context, *models.Category) (*models.Category, error)
	DeleteCategory(context.Context, uint64) error
	FindAllCategory(context.Context, *requests.GlobalQuery) ([]models.Category, *responses.Pagination, error)
	FindByID(context.Context, uint64) (*models.Category, error)
	FindAllCategories(ctx context.Context) ([]models.Category, error)
}
