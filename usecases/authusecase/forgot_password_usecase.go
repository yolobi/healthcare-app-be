package authusecase

import (
	"context"
	"fmt"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/entities/models/transaction"
	"healthcare-capt-america/interfaces/repositories/authrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/userrepositoryinterface"
	"healthcare-capt-america/interfaces/usecases/authusecaseinterface"
	"healthcare-capt-america/services"
	"healthcare-capt-america/utils"
)

type forgotPasswordUsecase struct {
	fpp authrepositoryinterface.ForgotPasswordRepository
	ur  userrepositoryinterface.UserRepository
}

func NewForgotPasswordUsecase(fpp authrepositoryinterface.ForgotPasswordRepository, ur userrepositoryinterface.UserRepository) *forgotPasswordUsecase {
	return &forgotPasswordUsecase{
		fpp: fpp,
		ur:  ur,
	}
}

func (fpu *forgotPasswordUsecase) GetToken(ctx context.Context, email string) error {
	user, err := fpu.ur.FindByEmail(ctx, email)
	if err != nil {
		return err
	}
	if user == nil || !user.IsVerify {
		return apperror.NewClientError(fmt.Errorf("can't find user with email %s", email))
	}
	existing, err := fpu.fpp.FindByUserId(ctx, user.ID)
	if err != nil {
		return err
	}
	if existing != nil {
		err2 := services.SendResetPasswordEmail(user.Email, existing.Token)
		if err2 != nil {
			return apperror.NewServerError(err2)
		}
	}
	return fpu.fpp.Save(ctx, &models.ForgotPasswordToken{
		UserId: user.ID,
		Token:  utils.GenerateToken(15),
	})
}

func (fpu *forgotPasswordUsecase) ResetPassword(ctx context.Context, reset *transaction.ResetPassword) error {
	existing, err := fpu.fpp.FindByToken(ctx, reset.Token)
	if err != nil {
		return err
	}
	if existing == nil || !existing.IsValid {
		return apperror.NewClientError(fmt.Errorf("token is invalid"))
	}
	user, err := fpu.ur.FindById(ctx, existing.UserId)
	if err != nil {
		return err
	}
	if user == nil {
		return apperror.NewServerError(fmt.Errorf("can't find user"))
	}
	hashed, err := utils.HashPassword(reset.Password)
	if err != nil {
		return apperror.NewServerError(err)
	}
	user.Password = &hashed
	return fpu.fpp.ResetPassword(ctx, user)
}

var _ authusecaseinterface.ForgotPasswordUsecase = &forgotPasswordUsecase{}
