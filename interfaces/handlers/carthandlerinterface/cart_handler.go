package carthandlerinterface

import "github.com/gin-gonic/gin"

type CartHandler interface {
	AddCart(c *gin.Context)
	GetCartsByLoginUser(c *gin.Context)
	DeleteCartById(c *gin.Context)
	DeleteAllDrugs(c *gin.Context)
}
