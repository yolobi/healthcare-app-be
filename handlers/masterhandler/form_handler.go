package masterhandler

import (
	"github.com/gin-gonic/gin"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/dto/responses/formres"
	"healthcare-capt-america/interfaces/handlers/masterhandlerinterface"
	"healthcare-capt-america/interfaces/usecases/masterusecaseinterface"
	"net/http"
)

type FormHandler struct {
	usecase masterusecaseinterface.FormUsecase
}

func (f FormHandler) GetAllForms(c *gin.Context) {
	forms, err := f.usecase.GetAllForms(c.Request.Context())
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, responses.DefaultResponse{Data: gin.H{"drug_forms": formres.NewForms(forms)}})
}

func NewFormHandler(usecase masterusecaseinterface.FormUsecase) *FormHandler {
	return &FormHandler{usecase: usecase}
}

var _ masterhandlerinterface.FormHandler = &FormHandler{}
