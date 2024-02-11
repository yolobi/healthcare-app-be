package authrepository

import (
	"context"
	"errors"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/repositories/authrepositoryinterface"
	"healthcare-capt-america/services"
	"healthcare-capt-america/utils"

	"gorm.io/gorm"
)

type forgotPasswordRepository struct {
	db *gorm.DB
}

func (fpp *forgotPasswordRepository) Save(ctx context.Context, token *models.ForgotPasswordToken) error {
	err := fpp.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		token, err := utils.SaveQuery[models.ForgotPasswordToken](ctx, tx, token, enums.Create)
		if err != nil {
			return apperror.NewServerError(err)
		}
		var user *models.User
		err = tx.Model(&models.User{}).
			Where("id", token.UserId).
			First(&user).Error
		if err != nil {
			return apperror.NewServerError(err)
		}
		err = services.SendResetPasswordEmail(user.Email, token.Token)
		if err != nil {
			return apperror.NewServerError(err)
		}
		return nil
	})
	return err
}

func (fpp *forgotPasswordRepository) FindByUserId(ctx context.Context, user_id uint64) (*models.ForgotPasswordToken, error) {
	var token *models.ForgotPasswordToken
	err := fpp.db.WithContext(ctx).
		Where("user_id", user_id).
		Where("is_valid", true).
		First(&token).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, apperror.NewServerError(err)
	}
	return token, nil
}

func (fpp *forgotPasswordRepository) FindByToken(ctx context.Context, token string) (*models.ForgotPasswordToken, error) {
	var existringToken *models.ForgotPasswordToken
	err := fpp.db.WithContext(ctx).
		Where("token", token).
		First(&existringToken).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return existringToken, nil
}

func (fpp *forgotPasswordRepository) ResetPassword(ctx context.Context, user *models.User) error {
	err := fpp.db.WithContext(ctx).Save(&user).Error
	if err != nil {
		return apperror.NewServerError(err)
	}
	return nil
}

func NewForgotPasswordRepository(db *gorm.DB) *forgotPasswordRepository {
	return &forgotPasswordRepository{
		db: db,
	}
}

var _ authrepositoryinterface.ForgotPasswordRepository = &forgotPasswordRepository{}
