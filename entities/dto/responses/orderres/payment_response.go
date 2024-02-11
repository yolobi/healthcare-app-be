package orderres

import "healthcare-capt-america/entities/models"

type PaymentOrder struct {
	ID          uint64 `json:"id"`
	PaymentFile string `json:"payment_file"`
	Status      string `json:"status"`
}

func newPaymentOrder(payment *models.Payment) *PaymentOrder {
	return &PaymentOrder{
		ID:          payment.ID,
		PaymentFile: "/" + payment.File,
		Status:      payment.Status,
	}
}

func getFileDir(file string) string {
	if file != "" {
		file = file[len("/varmasea"):]
		file = file + "?authuser=2"
		return file
	}
	return ""
}
