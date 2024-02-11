package producthandlerinterface

import (
	"github.com/gin-gonic/gin"
)

type DrugHandler interface {
	CreateDrug(*gin.Context)
	EditDrug(*gin.Context)
	FindAllDrug(*gin.Context)
	FindByID(*gin.Context)
	FindAllWithDist(*gin.Context)
}
