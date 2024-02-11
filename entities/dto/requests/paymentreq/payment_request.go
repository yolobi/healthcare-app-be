package paymentreq

import (
	"healthcare-capt-america/entities/models"
	"mime/multipart"
)

type PaymentRequest struct {
	ID   uint64               `form:"id" binding:"required"`
	File multipart.FileHeader `form:"file" binding:"required"`
}

func (req *PaymentRequest) ToPayment() *models.Payment {
	return &models.Payment{ID: req.ID}
}

type PaymentStatusRequest struct {
	ID     uint64 `json:"id"`
	Status string `json:"status"`
}

func (req *PaymentStatusRequest) ToPayment() *models.Payment {
	return &models.Payment{ID: req.ID, Status: req.Status}
}
