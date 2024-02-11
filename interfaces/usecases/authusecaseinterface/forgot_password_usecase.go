package authusecaseinterface

import (
	"context"
	"healthcare-capt-america/entities/models/transaction"
)

type ForgotPasswordUsecase interface {
	GetToken(ctx context.Context, email string) error
	ResetPassword(ctx context.Context, reset *transaction.ResetPassword) error
}
