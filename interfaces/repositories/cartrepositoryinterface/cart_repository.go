package cartrepositoryinterface

import (
	"context"
	"healthcare-capt-america/entities/dto/responses/checkoutres"
	"healthcare-capt-america/entities/models"
)

type CartRepository interface {
	Save(ctx context.Context, cart *models.Cart) (*models.Cart, error)
	FindCartListByUserId(ctx context.Context, userId uint64) ([]*models.Cart, error)
	FindByID(ctx context.Context, cartId uint64) (*models.Cart, error)
	CheckDrugInCartExists(ctx context.Context, drugId, userId uint64) (bool, error)
	UpdateDrugCartQuantity(ctx context.Context, cart *models.Cart) (*models.Cart, error)
	Delete(ctx context.Context, cart *models.Cart) error
	DeleteCartByUserID(ctx context.Context, userId uint64) error
	FindCartListUserForCheckout(ctx context.Context, userId uint64, pharmacyId uint64) ([]checkoutres.CartResponse, error)
}
