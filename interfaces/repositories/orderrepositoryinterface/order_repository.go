package orderrepositoryinterface

import (
	"context"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/models"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *models.Order) (*models.Order, error)
	FindAll(ctx context.Context, qry *requests.GlobalQuery) ([]*models.Order, *responses.Pagination, error)
	FindByID(ctx context.Context, id uint64) (*models.Order, error)
	Delete(ctx context.Context, id uint64) error
	FindByPaymentID(ctx context.Context, paymentId uint64) (*models.Order, error)
	Update(ctx context.Context, order *models.Order) error
	SetConfirmed(ctx context.Context) error
	SetCanceled(ctx context.Context) error
}
