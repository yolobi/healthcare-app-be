package authhandlerinterface

import "github.com/gin-gonic/gin"

type ForgotPasswordHandler interface {
	GetToken(ctx *gin.Context)
	ResetPassword(ctx *gin.Context)
}
