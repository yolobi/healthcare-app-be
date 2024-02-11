package authrepository

import (
	"context"
	"errors"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/interfaces/repositories/authrepositoryinterface"

	"gorm.io/gorm"
)

type registerTokenRepository struct {
	db *gorm.DB
}

func (rt *registerTokenRepository) FindByUserId(ctx context.Context, user_id uint64) (*models.RegisterAccountToken, error) {
	var token *models.RegisterAccountToken
	err := rt.db.WithContext(ctx).
		Where("user_id", user_id).
		Where("is_valid", true).
		First(&token).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, apperror.NewServerError(err, "Database error")
	}
	return token, nil
}

func (rt *registerTokenRepository) FindByToken(ctx context.Context, token string) (*models.RegisterAccountToken, error) {
	var registerToken *models.RegisterAccountToken
	err := rt.db.WithContext(ctx).
		Where("token", token).
		First(&registerToken).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, apperror.NewClientError(err, "Database error")
	}
	return registerToken, nil
}

func NewRegisterTokenRepostiory(db *gorm.DB) *registerTokenRepository {
	return &registerTokenRepository{
		db: db,
	}
}

var _ authrepositoryinterface.RegisterTokenRepository = &registerTokenRepository{}
