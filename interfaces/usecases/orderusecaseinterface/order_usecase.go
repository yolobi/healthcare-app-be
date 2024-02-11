package orderusecaseinterface

import (
	"context"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/models"
)

type OrderUsecase interface {
	CreateOrder(ctx context.Context, order *models.Order) (*models.Order, error)
	GetAllOrders(ctx context.Context, qry *requests.GlobalQuery) ([]*models.Order, *responses.Pagination, error)
	DeleteOrderById(ctx context.Context, id uint64) error
	GetOrderById(ctx context.Context, id uint64) (*models.Order, error)
	UpdateOrderStatus(ctx context.Context, id uint64) (*models.Order, error)
	GetUserOrder(ctx context.Context, status string, qry *requests.GlobalQuery) ([]*models.Order, *responses.Pagination, error)
	GetUserDetailOrder(ctx context.Context, id uint64) (*models.Order, error)
	UpdateUserOrder(ctx context.Context, order_id uint64, status string) error
	AdminUpdateOrder(ctx context.Context, order_id uint64, status string) error
}
