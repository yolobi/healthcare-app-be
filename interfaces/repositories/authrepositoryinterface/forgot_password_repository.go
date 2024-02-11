package authrepositoryinterface

import (
	"context"
	"healthcare-capt-america/entities/models"
)

type ForgotPasswordRepository interface {
	Save(ctx context.Context, token *models.ForgotPasswordToken) error
	FindByUserId(ctx context.Context, user_id uint64) (*models.ForgotPasswordToken, error)
	ResetPassword(ctx context.Context, user *models.User) error
	FindByToken(ctx context.Context, token string) (*models.ForgotPasswordToken, error)
}
