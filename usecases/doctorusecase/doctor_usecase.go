package doctorusecase

import (
	"context"
	"fmt"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/entities/models/transaction"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/repositories/doctorrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/userrepositoryinterface"
	"healthcare-capt-america/interfaces/usecases/doctorusecaseinterface"
	"healthcare-capt-america/services"
	"mime/multipart"

	"golang.org/x/crypto/bcrypt"
)

type doctorUsecase struct {
	dr doctorrepositoryinterface.DoctorRepository
	ur userrepositoryinterface.UserRepository
}

func (du *doctorUsecase) UpdateProfile(ctx context.Context, doctor *transaction.UpdateDoctor) (*models.Doctor, error) {
	oldUser, err := du.ur.FindById(ctx, doctor.UserId)
	if err != nil {
		return nil, err
	}
	if oldUser == nil {
		return nil, apperror.NewClientError(fmt.Errorf("user not found"))
	}
	oldDoctor, err := du.dr.FindByUserId(ctx, doctor.UserId)
	if err != nil {
		return nil, err
	}
	if oldDoctor == nil {
		return nil, apperror.NewClientError(fmt.Errorf("doctor not found"))
	}
	if doctor.OldPassword != nil {
		err := bcrypt.CompareHashAndPassword([]byte(*oldUser.Password), []byte(*doctor.OldPassword))
		if err != nil {
			return nil, apperror.NewClientError(fmt.Errorf("incorrect old password"))
		}
		hashed, err := bcrypt.GenerateFromPassword([]byte(*doctor.NewPassword), enums.DefaultCost)
		if err != nil {
			return nil, apperror.NewServerError(err)
		}
		*oldUser.Password = string(hashed)
	}
	fh := ctx.Value(enums.UserPhotoKey).(*multipart.FileHeader)
	if fh != nil {
		s, err := services.UploadUserImage(ctx, fh, oldUser.Email)
		if err != nil {
			return nil, apperror.NewServerError(err)
		}
		oldUser.Photo = *s
	}
	oldUser.Name = doctor.Name
	oldUser.PhoneNumber = &doctor.PhoneNumber
	oldDoctor.Certificate = doctor.Certificate
	oldDoctor.YearsOfExperience = doctor.YearsOfExperience
	if oldDoctor.SpecializationId == nil {
		oldDoctor.SpecializationId = &doctor.SpecializationId
	} else {
		*oldDoctor.SpecializationId = doctor.SpecializationId
	}
	d, err := du.dr.UpdateProfile(ctx, oldUser, oldDoctor)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (du *doctorUsecase) GetAllDoctor(ctx context.Context, qry *requests.GlobalQuery) ([]*models.Doctor, *responses.Pagination, error) {
	doctors, pagination, err := du.dr.FindAll(ctx, qry)
	if err != nil {
		return nil, nil, err
	}
	return doctors, pagination, nil
}

func (du *doctorUsecase) GetDetailDoctor(ctx context.Context, id uint64) (*models.Doctor, error) {
	doctor, err := du.dr.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	if doctor == nil {
		return nil, apperror.NewClientError(fmt.Errorf("can't find doctor with id %d", id))
	}
	return doctor, nil
}

func (du *doctorUsecase) GetCurrentDetailDoctor(ctx context.Context, id uint64) (*models.Doctor, error) {
	doctor, err := du.dr.FindByUserId(ctx, id)
	if err != nil {
		return nil, err
	}
	if doctor == nil {
		return nil, apperror.NewClientError(fmt.Errorf("can't find doctor with id %d", id))
	}
	return doctor, nil
}

func (du *doctorUsecase) UpdateStatusDoctor(ctx context.Context, status bool, id uint64) error {
	doctor, err := du.dr.FindById(ctx, id)
	if err != nil {
		return err
	}
	if doctor == nil {
		return apperror.NewClientError(fmt.Errorf("can't find admin with id %d", id))
	}
	user, err := du.ur.FindById(ctx, doctor.UserId)
	if err != nil {
		return err
	}
	if user == nil {
		return apperror.NewClientError(fmt.Errorf("can't find user with id %d", user.ID))
	}
	user.IsVerify = status
	_, err = du.ur.Update(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func NewDoctorUsecase(dr doctorrepositoryinterface.DoctorRepository, ur userrepositoryinterface.UserRepository) *doctorUsecase {
	return &doctorUsecase{
		dr: dr,
		ur: ur,
	}
}

var _ doctorusecaseinterface.DoctorUsecase = &doctorUsecase{}
