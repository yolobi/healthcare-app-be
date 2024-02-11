package pharmacyproducthandlerinterface

import "github.com/gin-gonic/gin"

type PharmacyDrugHandler interface {
	CreatePharmacyDrug(*gin.Context)
	EditPharmacyDrug(*gin.Context)
	DeletePharmacyDrug(*gin.Context)
	FindAllByPharmacyID(*gin.Context)
	FindByID(*gin.Context)
	GetAllProductsByAdminPharmacy(c *gin.Context)
	EditProductInPharmacy(c *gin.Context)
}
