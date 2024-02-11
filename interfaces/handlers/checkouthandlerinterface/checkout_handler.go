package checkouthandlerinterface

import "github.com/gin-gonic/gin"

type CheckoutHandler interface {
	Checkout(c *gin.Context)
}
