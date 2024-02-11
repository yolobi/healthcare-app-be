package adminusecase

import (
	"context"
	"fmt"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/repositories/adminrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/userrepositoryinterface"
	"healthcare-capt-america/interfaces/usecases/adminusecaseinterface"
	"healthcare-capt-america/services"
	"mime/multipart"
)

type adminPharmacyUsecase struct {
	AdminPharmacyRepo adminrepositoryinterface.AdminPharmacyRepository
	UserRepo          userrepositoryinterface.UserRepository
}

func NewAdminPharmacyUsecase(apr adminrepositoryinterface.AdminPharmacyRepository, ur userrepositoryinterface.UserRepository) *adminPharmacyUsecase {
	return &adminPharmacyUsecase{
		AdminPharmacyRepo: apr,
		UserRepo:          ur,
	}
}

func (apu *adminPharmacyUsecase) CreateAdminPharmacy(ctx context.Context, adminPharmacy *models.User) (*models.AdminPharmacy, error) {
	user, err := apu.UserRepo.FindByEmail(ctx, adminPharmacy.Email)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return nil, apperror.NewClientError(fmt.Errorf("email already used"))
	}
	adminPharmacy.IsVerify = true
	admin, err := apu.AdminPharmacyRepo.Create(ctx, adminPharmacy)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (apu *adminPharmacyUsecase) GetAllAdmin(ctx context.Context, qry *requests.GlobalQuery) ([]models.AdminPharmacy, *responses.Pagination, error) {
	admins, pagination, err := apu.AdminPharmacyRepo.FindAll(ctx, qry)
	if err != nil {
		return nil, nil, err
	}
	return admins, pagination, nil
}

func (apu *adminPharmacyUsecase) Delete(ctx context.Context, admin_id uint64) error {
	admin, err := apu.AdminPharmacyRepo.FindBydId(ctx, admin_id)
	if err != nil {
		return err
	}
	if admin == nil {
		return apperror.NewClientError(fmt.Errorf("can't find admin with id %d", admin_id))
	}
	err = apu.AdminPharmacyRepo.Delete(ctx, admin)
	if err != nil {
		return err
	}
	return nil
}

func (apu *adminPharmacyUsecase) GetDetailAdmin(ctx context.Context, id uint64) (*models.AdminPharmacy, error) {
	admin, err := apu.AdminPharmacyRepo.FindBydId(ctx, id)
	if err != nil {
		return nil, err
	}
	if admin == nil {
		return nil, apperror.NewClientError(fmt.Errorf("can't find admin with id %d", id))
	}
	return admin, nil
}

func (apu *adminPharmacyUsecase) UpdateAdmin(ctx context.Context, admin_id uint64, user *models.User) error {
	admin, err := apu.AdminPharmacyRepo.FindBydId(ctx, admin_id)
	if err != nil {
		return err
	}
	if admin == nil {
		return apperror.NewClientError(fmt.Errorf("can't find admin with id %d", admin_id))
	}
	oldUser, err := apu.UserRepo.FindById(ctx, admin.UserId)
	if err != nil {
		return err
	}
	if oldUser == nil {
		return apperror.NewClientError(fmt.Errorf("can't find user with id %d", admin.UserId))
	}
	fh := ctx.Value(enums.UserPhotoKey).(*multipart.FileHeader)
	if fh != nil {
		s, err := services.UploadUserImage(ctx, fh, oldUser.Email)
		if err != nil {
			return apperror.NewServerError(err)
		}
		oldUser.Photo = *s
	}
	updatedUser := oldUser.UpdateUserData(user)
	_, err = apu.UserRepo.Update(ctx, updatedUser)
	if err != nil {
		return err
	}
	return nil
}

var _ adminusecaseinterface.AdminPharmacyUsecase = &adminPharmacyUsecase{}
