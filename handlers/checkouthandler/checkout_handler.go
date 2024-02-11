package checkouthandler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/dto/requests/checkoutreq"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/dto/responses/checkoutres"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/handlers/checkouthandlerinterface"
	"healthcare-capt-america/interfaces/usecases/checkoutusecaseinterface"
	"net/http"
)

type CheckoutHandler struct {
	usecase checkoutusecaseinterface.CheckoutUsecase
}

func (handler *CheckoutHandler) Checkout(c *gin.Context) {
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(c.Request.Context(), enums.RequestIdKey.Key, rID)
		id, _   = c.Get("user-id")
		request *checkoutreq.CheckoutRequest
	)
	if err := c.ShouldBind(&request); err != nil {
		c.Error(apperror.NewValidationError(err))
		return
	}
	uID := id.(uint64)
	pharmacyID, carts, err := handler.usecase.CheckoutUser(ctx, uID, request.AddressID)
	if err != nil {
		c.Error(err)
		return
	}
	checkoutResponse := &checkoutres.CheckoutResponse{PharmacyID: pharmacyID, Carts: carts, TotalPrice: checkoutres.GetTotalPrice(carts)}
	c.JSON(http.StatusOK, responses.DefaultResponse{Data: gin.H{"checkout": checkoutResponse}})
}

func NewCheckoutHandler(usecase checkoutusecaseinterface.CheckoutUsecase) *CheckoutHandler {
	return &CheckoutHandler{usecase: usecase}
}

var _ checkouthandlerinterface.CheckoutHandler = &CheckoutHandler{}
