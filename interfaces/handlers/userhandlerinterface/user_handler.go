package userhandlerinterface

import "github.com/gin-gonic/gin"

type UserHandler interface {
	GetAllUsers(ctx *gin.Context)
	GetUserDetail(ctx *gin.Context)
	GetCurrentUserDetail(ctx *gin.Context)
	UpdateProfile(ctx *gin.Context)
	AddAddress(ctx *gin.Context)
	FindAllUserAddress(ctx *gin.Context)
	SetDefaultAddress(ctx *gin.Context)
	DeleteAddress(ctx *gin.Context)
	UpdateAddress(ctx *gin.Context)
}
