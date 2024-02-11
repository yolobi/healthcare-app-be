package stockmutationhandlerinterface

import (
	"github.com/gin-gonic/gin"
)

type StockMutationHandler interface {
	CreateRequestStockMutation(*gin.Context)
	FindAll(*gin.Context)
	FindById(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}
