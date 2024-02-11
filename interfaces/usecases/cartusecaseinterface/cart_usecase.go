package cartusecaseinterface

import (
	"context"
	"healthcare-capt-america/entities/models"
)

type CartUsecase interface {
	AddCart(ctx context.Context, cart *models.Cart) (*models.Cart, error)
	GetUserCart(ctx context.Context, id uint64) ([]*models.Cart, error)
	DeleteCartByCartId(ctx context.Context, id uint64) error
	DeleteCartByUserId(ctx context.Context, id uint64) error
}
