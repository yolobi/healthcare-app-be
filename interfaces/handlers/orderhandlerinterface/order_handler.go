package orderhandlerinterface

import "github.com/gin-gonic/gin"

type OrderHandler interface {
	CreateOrder(c *gin.Context)
	GetAllOrders(c *gin.Context)
	GetOrderById(c *gin.Context)
	DeleteOrderById(c *gin.Context)
	UpdateOrderStatus(c *gin.Context)
	GetUserOrders(ctx *gin.Context)
	GetUserDetailOrder(ctx *gin.Context)
	UpdateUserOrder(ctx *gin.Context)
	AdminUpdateOrder(ctx *gin.Context)
}
