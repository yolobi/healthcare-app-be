package masterrepositoryinterface

import (
	"context"
	"healthcare-capt-america/entities/models"
)

type PaymentRepository interface {
	Save(ctx context.Context, payment *models.Payment) (*models.Payment, error)
	FindByID(ctx context.Context, id uint64) (*models.Payment, error)
	ValidateUserPayment(ctx context.Context, paymentId uint64, userId uint64) (bool, error)
	Update(ctx context.Context, payment *models.Payment) (*models.Payment, error)
}
