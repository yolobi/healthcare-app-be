package checkoutusecase

import (
	"context"
	"fmt"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/dto/responses/checkoutres"
	"healthcare-capt-america/interfaces/repositories/cartrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/masterrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/orderrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/pharmacyproductrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/pharmacyrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/productrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/userrepositoryinterface"
	"healthcare-capt-america/interfaces/usecases/checkoutusecaseinterface"
)

type checkoutUsecase struct {
	cartRepo         cartrepositoryinterface.CartRepository
	userRepo         userrepositoryinterface.UserRepository
	productRepo      productrepositoryinterface.DrugRepository
	orderRepo        orderrepositoryinterface.OrderRepository
	pharmacyRepo     pharmacyrepositoryinterface.PharmacyRepository
	pharmacydrugRepo pharmacyproductrepositoryinterface.PharmacyDrugRepository
	addressRepo      masterrepositoryinterface.AddrressRepository
}

func (c checkoutUsecase) CheckoutUser(ctx context.Context, userId uint64, addressId uint64) (pharmacyId uint64, carts []checkoutres.CartResponse, err error) {
	cartModels, err := c.cartRepo.FindCartListByUserId(ctx, userId)
	if err != nil {
		return 0, nil, apperror.NewServerError(err)
	}
	address, err := c.addressRepo.FindById(ctx, addressId)
	if err != nil {
		return 0, nil, apperror.NewServerError(err)
	}
	if address == nil {
		return 0, nil, apperror.NewClientError(fmt.Errorf("user doesn't have default address"))
	}
	pharmacyId, err = c.pharmacydrugRepo.FindPharmacyIDCheckout(ctx, address, cartModels)
	if err != nil {
		return 0, nil, apperror.NewServerError(err)
	}
	if pharmacyId == 0 {
		return 0, nil, apperror.NewClientError(fmt.Errorf(`no pharmacy sell some product in nearest pharmacy`))
	}
	carts, err = c.cartRepo.FindCartListUserForCheckout(ctx, userId, pharmacyId)
	if err != nil {
		return 0, nil, apperror.NewServerError(err)
	}
	return pharmacyId, carts, nil
}

func NewCheckoutUsecase(cartRepo cartrepositoryinterface.CartRepository,
	userRepo userrepositoryinterface.UserRepository,
	productRepo productrepositoryinterface.DrugRepository,
	orderRepo orderrepositoryinterface.OrderRepository,
	pharmacyRepo pharmacyrepositoryinterface.PharmacyRepository,
	pharmacydrugRepo pharmacyproductrepositoryinterface.PharmacyDrugRepository,
	addressRepo masterrepositoryinterface.AddrressRepository,
) *checkoutUsecase {
	return &checkoutUsecase{
		cartRepo:         cartRepo,
		userRepo:         userRepo,
		productRepo:      productRepo,
		orderRepo:        orderRepo,
		pharmacyRepo:     pharmacyRepo,
		pharmacydrugRepo: pharmacydrugRepo,
		addressRepo:      addressRepo,
	}
}

var _ checkoutusecaseinterface.CheckoutUsecase = &checkoutUsecase{}
