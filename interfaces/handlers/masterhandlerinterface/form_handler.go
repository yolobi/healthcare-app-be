package masterhandlerinterface

import "github.com/gin-gonic/gin"

type FormHandler interface {
	GetAllForms(c *gin.Context)
}
