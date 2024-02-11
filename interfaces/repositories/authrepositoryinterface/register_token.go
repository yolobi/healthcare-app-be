package authrepositoryinterface

import (
	"context"
	"healthcare-capt-america/entities/models"
)

type RegisterTokenRepository interface {
	FindByUserId(ctx context.Context, user_id uint64) (*models.RegisterAccountToken, error)
	FindByToken(ctx context.Context, token string) (*models.RegisterAccountToken, error)
}
