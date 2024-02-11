package masterhandlerinterface

import "github.com/gin-gonic/gin"

type ShipmentHandler interface {
	FindAll(ctx *gin.Context)
	CalculateDistance(ctx *gin.Context)
}
