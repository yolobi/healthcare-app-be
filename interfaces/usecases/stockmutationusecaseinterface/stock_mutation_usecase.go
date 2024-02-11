package stockmutationusecaseinterface

import (
	"context"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/models"
)

type StockMutationUsecase interface {
	CreateRequestStockMutation(ctx context.Context, stockMutation *models.StockMutation, uId uint64) (*models.StockMutation, error)
	FindByID(ctx context.Context, stockMutation_id uint64, uId uint64) (*models.StockMutation, error)
	FindAll(ctx context.Context, qry *requests.GlobalQuery, uId uint64, action string) ([]*models.StockMutation, *responses.Pagination, error)
	Update(ctx context.Context, stockMutation *models.StockMutation, uId uint64) (*models.StockMutation, error)
	Delete(ctx context.Context, stockMutation_id uint64, uId uint64) error
}
