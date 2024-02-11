package doctorhandlerinterface

import "github.com/gin-gonic/gin"

type DoctorHandler interface {
	GetAllDoctor(ctx *gin.Context)
	GetDetailDoctor(ctx *gin.Context)
	GetCurrentDoctorDetail(ctx *gin.Context)
	UpdateStatusDoctor(ctx *gin.Context)
	UpdateProfile(ctx *gin.Context)
}
