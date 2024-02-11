package masterhandlerinterface

import "github.com/gin-gonic/gin"

type SpecializationHandler interface {
	GetAllSpecializations(c *gin.Context)
}
