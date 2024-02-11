package userusecase

import (
	"context"
	"fmt"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/entities/models/transaction"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/repositories/masterrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/userrepositoryinterface"
	"healthcare-capt-america/interfaces/usecases/userusecaseinterface"
	"healthcare-capt-america/services"
	"mime/multipart"

	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepo    userrepositoryinterface.UserRepository
	addressRepo masterrepositoryinterface.AddrressRepository
}

func (uu *userUsecase) GetAllUsers(ctx context.Context, qry *requests.GlobalQuery) ([]*models.User, *responses.Pagination, error) {
	users, pagination, err := uu.userRepo.FindAll(ctx, qry)
	if err != nil {
		return nil, nil, err
	}
	return users, pagination, nil
}

func (uu *userUsecase) GetUserDetail(ctx context.Context, user_id uint64) (*models.User, error) {
	user, err := uu.userRepo.FindById(ctx, user_id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, apperror.NewClientError(fmt.Errorf("can't find user with id %d", user_id))
	}
	return user, nil
}

func (uu *userUsecase) UpdateProfile(ctx context.Context, user *transaction.UpdateUser) (*models.User, error) {
	oldUser, err := uu.userRepo.FindById(ctx, user.ID)
	if err != nil {
		return nil, err
	}
	if oldUser == nil {
		return nil, apperror.NewClientError(fmt.Errorf("this user is not found"))
	}
	if user.OldPassword != nil {
		err := bcrypt.CompareHashAndPassword([]byte(*oldUser.Password), []byte(*user.OldPassword))
		if err != nil {
			return nil, apperror.NewClientError(fmt.Errorf("incorrect old password"))
		}
		hashed, err := bcrypt.GenerateFromPassword([]byte(*user.NewPassword), enums.DefaultCost)
		if err != nil {
			return nil, apperror.NewServerError(err)
		}
		*oldUser.Password = string(hashed)
	}
	oldUser.Name = user.Name
	oldUser.PhoneNumber = &user.PhoneNumber
	fh := ctx.Value(enums.UserPhotoKey).(*multipart.FileHeader)
	if fh != nil {
		s, err := services.UploadUserImage(ctx, fh, oldUser.Email)
		if err != nil {
			return nil, apperror.NewServerError(err)
		}
		oldUser.Photo = *s
	}
	updatedUser, err := uu.userRepo.Update(ctx, oldUser)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func (uu *userUsecase) AddAddress(ctx context.Context, address *models.Address) (*models.Address, error) {
	addresses, err := uu.addressRepo.FindAllByUserId(ctx, *address.UserId)
	if err != nil {
		return nil, err
	}
	if len(addresses) >= 5 {
		return nil, apperror.NewClientError(fmt.Errorf("you only can add 5 address"))
	}
	if len(addresses) == 0 {
		address.IsDefault = true
	}
	newAddress, err := uu.addressRepo.Save(ctx, address)
	if err != nil {
		return nil, err
	}
	return newAddress, nil
}

func (uu *userUsecase) FindAllUserAddress(ctx context.Context, uid uint64) ([]*models.Address, error) {
	addresses, err := uu.addressRepo.FindAllByUserId(ctx, uid)
	if err != nil {
		return nil, err
	}
	return addresses, nil
}

func (uu *userUsecase) SetDefaultAddress(ctx context.Context, address_id uint64) error {
	address, err := uu.addressRepo.FindById(ctx, address_id)
	if err != nil {
		return err
	}
	if address == nil {
		return apperror.NewClientError(fmt.Errorf("can't find address with id %d", address_id))
	}
	uid := ctx.Value(enums.UserIdKey).(uint64)
	if *address.UserId != uid {
		return apperror.NewClientError(fmt.Errorf("this address is not yours"))
	}
	err = uu.addressRepo.SetDefault(ctx, address_id)
	return err
}

func (uu *userUsecase) DeleteAddress(ctx context.Context, address_id uint64) error {
	address, err := uu.addressRepo.FindById(ctx, address_id)
	if err != nil {
		return err
	}
	if address == nil {
		return apperror.NewClientError(fmt.Errorf("can't find address with id %d", address_id))
	}
	uid := ctx.Value(enums.UserIdKey).(uint64)
	if *address.UserId != uid {
		return apperror.NewClientError(fmt.Errorf("this address is not yours"))
	}
	if address.IsDefault {
		return apperror.NewClientError(fmt.Errorf("can't delete default addresss"))
	}
	err = uu.addressRepo.Delete(ctx, address_id)
	return err
}

func (uu *userUsecase) UpdateAddress(ctx context.Context, address *models.Address) (*models.Address, error) {
	oldAddress, err := uu.addressRepo.FindById(ctx, address.ID)
	if err != nil {
		return nil, err
	}
	if oldAddress == nil {
		return nil, apperror.NewClientError(fmt.Errorf("can't find address with id %d", address.ID))
	}
	uid := ctx.Value(enums.UserIdKey).(uint64)
	if *oldAddress.UserId != uid {
		return nil, apperror.NewClientError(fmt.Errorf("this is not your address"))
	}
	oldAddress.Detail = address.Detail
	oldAddress.ProvinceID = address.ProvinceID
	oldAddress.CityID = address.CityID
	oldAddress.Longtitude = address.Longtitude
	oldAddress.Latitude = address.Latitude
	a, err := uu.addressRepo.Update(ctx, oldAddress)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func NewUserUsecase(userRepo userrepositoryinterface.UserRepository, addressRepo masterrepositoryinterface.AddrressRepository) *userUsecase {
	return &userUsecase{
		userRepo:    userRepo,
		addressRepo: addressRepo,
	}
}

var _ userusecaseinterface.UserUsecase = &userUsecase{}
