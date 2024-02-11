package masterhandler

import (
	"github.com/gin-gonic/gin"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/dto/responses/manufactureres"
	"healthcare-capt-america/interfaces/handlers/masterhandlerinterface"
	"healthcare-capt-america/interfaces/usecases/masterusecaseinterface"
	"net/http"
)

type ManufactureHandler struct {
	usecase masterusecaseinterface.ManufactureUsecase
}

func (m ManufactureHandler) GetAllManufactures(c *gin.Context) {
	manufactures, err := m.usecase.GetAllManufactures(c.Request.Context())
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, responses.DefaultResponse{Data: gin.H{"manufactures": manufactureres.NewManufactures(manufactures)}})
}

var _ masterhandlerinterface.ManufactureHandler = &ManufactureHandler{}

func NewManufactureHandler(usecase masterusecaseinterface.ManufactureUsecase) *ManufactureHandler {
	return &ManufactureHandler{usecase: usecase}
}
