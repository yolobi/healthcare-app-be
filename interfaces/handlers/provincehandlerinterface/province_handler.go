package provincehandlerinterface

import "github.com/gin-gonic/gin"

type ProvinceHandler interface {
	GetProvinceById(c *gin.Context)
	GetAllProvinces(c *gin.Context)
}
