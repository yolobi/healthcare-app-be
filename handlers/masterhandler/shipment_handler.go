package masterhandler

import (
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/dto/requests/shipmentreq"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/interfaces/handlers/masterhandlerinterface"
	"healthcare-capt-america/interfaces/usecases/masterusecaseinterface"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ShipmentHandler struct {
	shipUsecase masterusecaseinterface.ShipmentUsecase
}

func (sh *ShipmentHandler) FindAll(ctx *gin.Context) {
	shipments, err := sh.shipUsecase.FindAll(ctx.Request.Context())
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := make([]shipmentreq.GetAllShipmentReq, 0)
	for _, ship := range shipments {
		resp = append(resp, shipmentreq.NewGetAllReq(ship))
	}
	ctx.JSON(http.StatusOK, responses.DefaultResponse{
		Data: resp,
	})
}

func (sh *ShipmentHandler) CalculateDistance(ctx *gin.Context) {
	var request shipmentreq.CalculateDistanceReq
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.Error(apperror.NewValidationError(err))
		return
	}

	shipments, err := sh.shipUsecase.CalculateDistance(ctx.Request.Context(), request.AddressId, request.PharmacyId, request.Weight)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"shipments": shipments})
}

func NewShipmentHandler(su masterusecaseinterface.ShipmentUsecase) *ShipmentHandler {
	return &ShipmentHandler{
		shipUsecase: su,
	}
}

var _ masterhandlerinterface.ShipmentHandler = &ShipmentHandler{}
