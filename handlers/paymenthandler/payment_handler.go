package paymenthandler

import (
	"context"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/dto/requests/paymentreq"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/handlers/paymenthandlerinterface"
	"healthcare-capt-america/interfaces/usecases/paymentusecaseinterface"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PaymentHandler struct {
	usecase paymentusecaseinterface.PaymentUsecase
}

func (handler *PaymentHandler) UpdatePaymentStatus(c *gin.Context) {
	ctx := context.WithValue(c.Request.Context(), enums.RequestIdKey, uuid.NewString())
	paymentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(apperror.NewClientError(err))
		return
	}
	var request *paymentreq.PaymentStatusRequest
	if err := c.ShouldBind(&request); err != nil {
		c.Error(apperror.NewValidationError(err))
		return
	}
	payment, err := handler.usecase.UpdatePaymentStatus(ctx, uint64(paymentId), enums.PaymentApproved)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, responses.DefaultResponse{Data: payment})
}

func (handler *PaymentHandler) UploadPayment(c *gin.Context) {
	var (
		request *paymentreq.PaymentRequest
	)
	if err := c.ShouldBind(&request); err != nil {
		c.Error(apperror.NewValidationError(err))
		return
	}
	userId := c.Value(enums.UserIdKey.Key).(uint64)
	fh, err := c.FormFile(enums.PaymentFile)
	if err != nil {
		c.Error(apperror.NewServerError(err))
		return
	}
	ctx2 := context.WithValue(c.Request.Context(), enums.PaymentFileKey.Key, fh)
	ctx := context.WithValue(ctx2, enums.UserIdKey.Key, userId)
	c.Request = c.Request.WithContext(ctx)
	payment := request.ToPayment()
	payment, err = handler.usecase.UploadPaymentFile(c.Request.Context(), payment)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, responses.DefaultResponse{Data: gin.H{"payment": payment}})
}

func NewPaymentHandler(usecase paymentusecaseinterface.PaymentUsecase) *PaymentHandler {
	return &PaymentHandler{usecase: usecase}
}

var _ paymenthandlerinterface.PaymentHandler = &PaymentHandler{}
