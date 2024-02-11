package checkoutusecaseinterface

import (
	"context"
	"healthcare-capt-america/entities/dto/responses/checkoutres"
)

type CheckoutUsecase interface {
	CheckoutUser(ctx context.Context, userId uint64, addressId uint64) (pharmacyId uint64, carts []checkoutres.CartResponse, err error)
}
