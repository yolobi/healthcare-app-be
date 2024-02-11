package pharmacyhandlerinterface

import "github.com/gin-gonic/gin"

type PharmacyHandler interface {
	AddPharmacy(c *gin.Context)
	GetPharmacies(c *gin.Context)
	GetPharmacyById(c *gin.Context)
	UpdatePharmacy(c *gin.Context)
	DeletePharmacy(c *gin.Context)
	GetPharmaciesByLoginAdminPharmacy(c *gin.Context)
}
