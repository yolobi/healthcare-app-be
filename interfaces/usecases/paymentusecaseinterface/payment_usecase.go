package paymentusecaseinterface

import (
	"context"
	"healthcare-capt-america/entities/models"
)

type PaymentUsecase interface {
	UploadPaymentFile(ctx context.Context, payment *models.Payment) (*models.Payment, error)
	UpdatePaymentStatus(ctx context.Context, id uint64, status string) (*models.Payment, error)
}
