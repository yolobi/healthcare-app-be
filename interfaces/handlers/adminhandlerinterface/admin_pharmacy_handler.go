package adminhandlerinterface

import "github.com/gin-gonic/gin"

type AdminPharmacyHandler interface {
	CreateAdminPharmacy(ctx *gin.Context)
	GetAllAdmin(ctx *gin.Context)
	DeleteAdminPharmacy(ctx *gin.Context)
	GetDetailAdmin(ctx *gin.Context)
	UpdateAdmin(ctx *gin.Context)
}
