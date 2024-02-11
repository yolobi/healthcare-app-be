package orderhandler

import (
	"context"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/requests/orderreq"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/dto/responses/orderres"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/handlers/orderhandlerinterface"
	"healthcare-capt-america/interfaces/usecases/orderusecaseinterface"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type OrderHandler struct {
	usecase orderusecaseinterface.OrderUsecase
}

// UpdateOrderStatus implements orderhandlerinterface.OrderHandler.
func (*OrderHandler) UpdateOrderStatus(c *gin.Context) {
	panic("unimplemented")
}

func (handler *OrderHandler) GetOrderById(c *gin.Context) {
	var (
		requestId = uuid.NewString()
		uid       = c.Value(enums.UserIdKey.Key)
		ctx       = context.WithValue(c.Request.Context(), "request_id", requestId)
		ctx2      = context.WithValue(ctx, enums.UserIdKey, uid)
		id        = c.Param("id")
	)
	oId, err := strconv.Atoi(id)
	if err != nil {
		c.Error(err)
		return
	}
	order, err := handler.usecase.GetOrderById(ctx2, uint64(oId))
	if err != nil {
		c.Error(err)
		return
	}
	resp := orderres.NewOrderResponse(order)
	c.JSON(http.StatusOK, responses.DefaultResponse{
		Data: gin.H{"order": resp},
	})
}

func (handler *OrderHandler) DeleteOrderById(c *gin.Context) {
	var (
		requestId = uuid.NewString()
		uid       = c.Value(enums.UserIdKey.Key)
		ctx       = context.WithValue(c.Request.Context(), "request_id", requestId)
		ctx2      = context.WithValue(ctx, enums.UserIdKey, uid)
		id        = c.Param("id")
	)
	oId, err := strconv.Atoi(id)
	if err != nil {
		c.Error(err)
		return
	}
	err = handler.usecase.DeleteOrderById(ctx2, uint64(oId))
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, responses.DefaultResponse{Message: "Success to delete order"})
}

func (handler *OrderHandler) GetAllOrders(c *gin.Context) {
	var (
		requestId = uuid.NewString()
		uid       = c.Value(enums.UserIdKey.Key)
		ctx       = context.WithValue(c.Request.Context(), "request_id", requestId)
		ctx2      = context.WithValue(ctx, enums.UserIdKey, uid)
		request   requests.GlobalFilter
	)
	err := c.ShouldBind(&request)
	if err != nil {
		err = c.Error(apperror.NewValidationError(err))
		return
	}
	qry := requests.NewQuery(&request)
	if request.PharmacyId != 0 {
		qry.AddCondition("orders.pharmacy_id", requests.Equal, request.PharmacyId)
	}
	orders, pagination, err := handler.usecase.GetAllOrders(ctx2, qry)
	if err != nil {
		return
	}
	resp := make([]orderres.OrderResponse, 0)
	for _, order := range orders {
		resp = append(resp, *orderres.NewOrderResponse(order))
	}
	pagination.Items = resp
	c.JSON(http.StatusOK, responses.DefaultResponse{
		Data: responses.NewPagination(pagination, "orders"),
	})
}

func (handler *OrderHandler) CreateOrder(c *gin.Context) {
	var (
		requestId = uuid.NewString()
		ctx       = context.WithValue(c.Request.Context(), "request_id", requestId)
		request   *orderreq.OrderRequest
		userId, _ = c.Get("user-id")
	)
	uId := userId.(uint64)
	if err := c.ShouldBind(&request); err != nil {
		return
	}
	order := request.NewOrderModel()
	order.UserId = uId
	order, err := handler.usecase.CreateOrder(ctx, order)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, responses.DefaultResponse{Data: gin.H{"order": orderres.NewOrderResponse(order)}})
}

func (handler *OrderHandler) GetUserOrders(ctx *gin.Context) {
	uid := ctx.Value(enums.UserIdKey.Key).(uint64)
	ctx2 := context.WithValue(ctx.Request.Context(), enums.UserIdKey, uid)
	var request requests.GlobalFilter

	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.Error(apperror.NewValidationError(err))
		return
	}

	qry := requests.NewQuery(&request)
	orders, pagination, err := handler.usecase.GetUserOrder(ctx2, request.OrderStatus, qry)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := make([]orderres.OrderResponse, 0)
	for _, order := range orders {
		resp = append(resp, *orderres.NewOrderResponse(order))
	}
	pagination.Items = resp
	ctx.JSON(http.StatusOK, responses.DefaultResponse{
		Data: responses.NewPagination(pagination, "orders"),
	})
}

func (handler *OrderHandler) GetUserDetailOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	uid := ctx.Value(enums.UserIdKey.Key).(uint64)
	ctx2 := context.WithValue(ctx.Request.Context(), enums.UserIdKey, uid)

	orderId, err := strconv.Atoi(id)
	if err != nil {
		ctx.Error(apperror.NewValidationError(err))
		return
	}
	order, err := handler.usecase.GetUserDetailOrder(ctx2, uint64(orderId))
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := orderres.NewOrderResponse(order)
	ctx.JSON(http.StatusOK, responses.DefaultResponse{
		Data: gin.H{"order": resp},
	})
}

func (handler *OrderHandler) UpdateUserOrder(ctx *gin.Context) {
	uid := ctx.Value(enums.UserIdKey.Key).(uint64)
	ctx2 := context.WithValue(ctx.Request.Context(), enums.UserIdKey, uid)
	id := ctx.Param("id")
	var request orderreq.UpdateOrderStatusReq
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.Error(apperror.NewValidationError(err))
		return
	}
	orderId, err := strconv.Atoi(id)
	if err != nil {
		ctx.Error(apperror.NewValidationError(err))
	}
	err = handler.usecase.UpdateUserOrder(ctx2, uint64(orderId), request.Status)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, responses.DefaultResponse{
		Message: "Successfully Update Order Status",
	})
}

func (handler *OrderHandler) AdminUpdateOrder(ctx *gin.Context) {
	uid := ctx.Value(enums.UserIdKey.Key).(uint64)
	ctx2 := context.WithValue(ctx.Request.Context(), enums.UserIdKey, uid)
	id := ctx.Param("id")
	var request orderreq.UpdateOrderStatusReq
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.Error(apperror.NewValidationError(err))
		return
	}
	orderId, err := strconv.Atoi(id)
	if err != nil {
		ctx.Error(apperror.NewValidationError(err))
	}
	err = handler.usecase.AdminUpdateOrder(ctx2, uint64(orderId), request.Status)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, responses.DefaultResponse{
		Message: "Successfully Update Order Status",
	})
}

func NewOrderHandler(usecase orderusecaseinterface.OrderUsecase) *OrderHandler {
	return &OrderHandler{usecase: usecase}
}

var _ orderhandlerinterface.OrderHandler = &OrderHandler{}
