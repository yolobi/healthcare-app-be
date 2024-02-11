package masterhandler

import (
	"github.com/gin-gonic/gin"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/dto/responses/specializationres"
	"healthcare-capt-america/interfaces/handlers/masterhandlerinterface"
	"healthcare-capt-america/interfaces/usecases/masterusecaseinterface"
	"net/http"
)

type SpecializationHandler struct {
	usecase masterusecaseinterface.SpecializationUsecase
}

func (s *SpecializationHandler) GetAllSpecializations(c *gin.Context) {
	specializations, err := s.usecase.GetAllSpecializations(c.Request.Context())
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, responses.DefaultResponse{Data: gin.H{"specializations": specializationres.NewSpecializations(specializations)}})
}

func NewSpecializationHandler(usecase masterusecaseinterface.SpecializationUsecase) *SpecializationHandler {
	return &SpecializationHandler{
		usecase: usecase,
	}
}

var _ masterhandlerinterface.SpecializationHandler = &SpecializationHandler{}
