package operationalhandlerinterface

import "github.com/gin-gonic/gin"

type OperationalHandler interface {
	AddOperationalDays(c *gin.Context)
	GetPharmacyOperationalDays(c *gin.Context)
	DeletePharmacyOperationalDay(c *gin.Context)
	UpdatePharmacyOperationalDay(c *gin.Context)
}
