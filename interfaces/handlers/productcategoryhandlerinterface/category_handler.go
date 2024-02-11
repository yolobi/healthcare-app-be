package productcategoryhandlerinterface

import "github.com/gin-gonic/gin"

type CategoryHandler interface {
	CreateCategory(*gin.Context)
	EditCategory(*gin.Context)
	DeleteCategory(*gin.Context)
	FindAllCategory(*gin.Context)
	FindByID(*gin.Context)
	FindAllCategories(c *gin.Context)
}
