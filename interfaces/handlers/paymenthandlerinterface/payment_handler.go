package paymenthandlerinterface

import "github.com/gin-gonic/gin"

type PaymentHandler interface {
	UploadPayment(c *gin.Context)
	UpdatePaymentStatus(c *gin.Context)
}
