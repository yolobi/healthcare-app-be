package carthandler

import (
	"context"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/dto/requests/cartreq"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/dto/responses/cartres"
	"healthcare-capt-america/interfaces/handlers/carthandlerinterface"
	"healthcare-capt-america/interfaces/usecases/cartusecaseinterface"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CartHandler struct {
	usecase cartusecaseinterface.CartUsecase
}

func (h *CartHandler) DeleteAllDrugs(c *gin.Context) {
	var (
		requestId = uuid.NewString()
		ctx       = context.WithValue(c.Request.Context(), "request_id", requestId)
		userId, _ = c.Get("user-id")
	)

	uId := userId.(uint64)
	if err := h.usecase.DeleteCartByUserId(ctx, uId); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, responses.DefaultResponse{Message: "Success deleted cart"})
}

func (h *CartHandler) DeleteCartById(c *gin.Context) {
	var (
		id        = c.Param("id")
		requestId = uuid.NewString()
		ctx       = context.WithValue(c.Request.Context(), "request_id", requestId)
	)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.Error(apperror.NewClientError(err))
		return
	}
	if err = h.usecase.DeleteCartByCartId(ctx, uint64(idInt)); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, responses.DefaultResponse{Message: "Success deleted drug from cart"})
}

// AddCart implements carthandlerinterface.CartHandler.
func (handler *CartHandler) AddCart(c *gin.Context) {
	var (
		request   *cartreq.CartRequest
		requestId = uuid.NewString()
		ctx       = context.WithValue(c.Request.Context(), "request_id", requestId)
		userId, _ = c.Get("user-id")
	)
	if err := c.ShouldBind(&request); err != nil {
		c.Error(apperror.NewValidationError(err))
		return
	}
	cart := request.NewCart()
	cart.UserId = userId.(uint64)
	cart, err := handler.usecase.AddCart(ctx, cart)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, &responses.DefaultResponse{Data: cart})
}

func (handler *CartHandler) GetCartsByLoginUser(c *gin.Context) {
	var (
		requestId = uuid.NewString()
		ctx       = context.WithValue(c.Request.Context(), "request_id", requestId)
		userId, _ = c.Get("user-id")
		uId       = userId.(uint64)
	)
	carts, err := handler.usecase.GetUserCart(ctx, uId)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, &responses.DefaultResponse{Data: cartres.NewCartResponse(carts)})
}

func NewCartHandler(usecase cartusecaseinterface.CartUsecase) *CartHandler {
	return &CartHandler{usecase: usecase}
}

var _ carthandlerinterface.CartHandler = &CartHandler{}
