package masterhandlerinterface

import "github.com/gin-gonic/gin"

type ManufactureHandler interface {
	GetAllManufactures(c *gin.Context)
}
