package authrepositoryinterface

import (
	"context"
	"healthcare-capt-america/entities/models"
)

type AuthRepository interface {
	Register(ctx context.Context, user *models.User, role string) error
	VerifyAccount(ctx context.Context, token *models.RegisterAccountToken, user *models.User) error
	ResendVerification(ctx context.Context, email string, token string) error
}
