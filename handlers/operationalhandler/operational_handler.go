package operationalhandler

import (
	"context"
	"healthcare-capt-america/entities/dto/requests/operationalreq"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/dto/responses/operationalres"
	"healthcare-capt-america/interfaces/handlers/operationalhandlerinterface"
	"healthcare-capt-america/interfaces/usecases/operationalusecaseinterface"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type OperationalHandler struct {
	usecase operationalusecaseinterface.OperationalUsecase
}

// UpdatePharmacyOperationalDay implements operationalhandlerinterface.OperationalHandler.
func (handler *OperationalHandler) UpdatePharmacyOperationalDay(c *gin.Context) {
	// var (
	// 	requestId = uuid.NewString()
	// 	ctx       = context.WithValue(c.Request.Context(), "request_id", requestId)
	// 	request   *operationalreq.OperationRequest
	// 	id, err   = strconv.Atoi(c.Param("id"))
	// )
	// if err != nil {
	// 	return
	// }
}

func (handler *OperationalHandler) AddOperationalDays(c *gin.Context) {
	var (
		requestId = uuid.NewString()
		ctx       = context.WithValue(c.Request.Context(), "request_id", requestId)
		request   *operationalreq.OperationRequest
	)
	if err := c.ShouldBind(&request); err != nil {
		return
	}
	req, err := request.OperationalModels()
	if err != nil {
		return
	}
	res, err := handler.usecase.AddOperationalDays(ctx, req)
	if err != nil {
		return
	}
	c.JSON(http.StatusCreated, responses.DefaultResponse{Data: operationalres.OperationalResponseConvert(res)})
}

func (handler *OperationalHandler) DeletePharmacyOperationalDay(c *gin.Context) {
	var (
		requestId = uuid.NewString()
		ctx       = context.WithValue(c.Request.Context(), "request_id", requestId)
		id, err   = strconv.Atoi(c.Param("id"))
	)
	if err != nil {
		return
	}
	err = handler.usecase.DeleteOperationalDay(ctx, uint64(id))
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, responses.DefaultResponse{Message: "Success delete operational day"})
}

func (handler *OperationalHandler) GetPharmacyOperationalDays(c *gin.Context) {
}

func NewOperationalHandler(usecase operationalusecaseinterface.OperationalUsecase) *OperationalHandler {
	return &OperationalHandler{
		usecase: usecase,
	}
}

var _ operationalhandlerinterface.OperationalHandler = &OperationalHandler{}
