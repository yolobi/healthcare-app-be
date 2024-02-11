package cartusecase

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/interfaces/repositories/cartrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/masterrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/pharmacyproductrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/userrepositoryinterface"
	"healthcare-capt-america/interfaces/usecases/cartusecaseinterface"
)

type cartUsecase struct {
	userRepo            userrepositoryinterface.UserRepository
	productPharmacyRepo pharmacyproductrepositoryinterface.PharmacyDrugRepository
	cartRepo            cartrepositoryinterface.CartRepository
	addressRepo         masterrepositoryinterface.AddrressRepository
}

func (usecase *cartUsecase) DeleteCartByUserId(ctx context.Context, id uint64) error {
	err := usecase.cartRepo.DeleteCartByUserID(ctx, id)
	if err != nil {
		return apperror.NewServerError(err)
	}
	return nil
}

func (usecase *cartUsecase) DeleteCartByCartId(ctx context.Context, id uint64) error {
	cart, err := usecase.cartRepo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperror.NewClientError(err)
		}
		return apperror.NewServerError(err)
	}
	err = usecase.cartRepo.Delete(ctx, cart)
	if err != nil {
		return apperror.NewServerError(err)
	}
	return nil
}

func (usecase *cartUsecase) AddCart(ctx context.Context, cart *models.Cart) (*models.Cart, error) {
	_, err := usecase.userRepo.FindById(ctx, cart.UserId)
	if err != nil {
		return nil, err
	}
	address, err := usecase.addressRepo.FindByUserId(ctx, cart.UserId)
	if err != nil {
		return nil, err
	}
	if address == nil {
		return nil, apperror.NewClientError(fmt.Errorf("user doesn't have default address"))
	}
	drugExists, err := usecase.productPharmacyRepo.CheckDrugIfExists(ctx, cart.DrugId)
	if err != nil {
		return nil, err
	}
	if !drugExists {
		return nil, apperror.NewClientError(errors.New("drug is not exists"))
	}
	drugInCart, err := usecase.cartRepo.CheckDrugInCartExists(ctx, cart.DrugId, cart.UserId)
	if err != nil {
		return nil, err
	}
	if drugInCart {
		return usecase.cartRepo.UpdateDrugCartQuantity(ctx, cart)
	}
	return usecase.cartRepo.Save(ctx, cart)
}

func (usecase *cartUsecase) GetUserCart(ctx context.Context, id uint64) ([]*models.Cart, error) {
	return usecase.cartRepo.FindCartListByUserId(ctx, id)
}

func NewCartUsecase(
	userRepo userrepositoryinterface.UserRepository,
	productPharmacyRepo pharmacyproductrepositoryinterface.PharmacyDrugRepository,
	cartRepo cartrepositoryinterface.CartRepository,
	addressRepo masterrepositoryinterface.AddrressRepository,
) *cartUsecase {
	return &cartUsecase{
		userRepo:            userRepo,
		productPharmacyRepo: productPharmacyRepo,
		cartRepo:            cartRepo,
		addressRepo:         addressRepo,
	}
}

var _ cartusecaseinterface.CartUsecase = &cartUsecase{}
