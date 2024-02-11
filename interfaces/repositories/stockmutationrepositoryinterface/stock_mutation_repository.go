package stockmutationrepositoryinterface

import (
	"context"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/models"
)

type StockMutationRepository interface {
	CreateRequestStockMutation(ctx context.Context, stockMutation *models.StockMutation) (*models.StockMutation, error)
	FindAll(ctx context.Context, qry *requests.GlobalQuery, uId uint64, action string) ([]*models.StockMutation, *responses.Pagination, error)
	FindByID(ctx context.Context, stockMutation_id uint64) (*models.StockMutation, error)
	AcceptRequestStockMutation(ctx context.Context, stockMutation *models.StockMutation, journal *models.Journal) (*models.StockMutation, error)
	Update(ctx context.Context, stockMutation *models.StockMutation) (*models.StockMutation, error)
	Delete(ctx context.Context, stockMutation *models.StockMutation) error
}
