package authusecaseinterface

import (
	"context"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/entities/models/transaction"
)

type AuthUsecase interface {
	Register(ctx context.Context, user *models.User, role string) error
	VerifyAccount(ctx context.Context, token string, password string) error
	Login(ctx context.Context, auth transaction.Authentication) (*string, *string, error)
}
