package authhandlerinterface

import "github.com/gin-gonic/gin"

type AuthHandler interface {
	Register(ctx *gin.Context)
	Verify(ctx *gin.Context)
	Login(ctx *gin.Context)
}
