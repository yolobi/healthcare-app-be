package stockmutationhandler

import (
	"fmt"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/requests/stockmutationreq"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/dto/responses/stockmutationres"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/handlers/stockmutationhandlerinterface"
	"healthcare-capt-america/interfaces/usecases/stockmutationusecaseinterface"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type StockMutationHandler struct {
	usecase stockmutationusecaseinterface.StockMutationUsecase
}

func (handler *StockMutationHandler) FindById(ctx *gin.Context) {
	uId := ctx.Value(enums.UserIdKey.Key).(uint64)
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		err = ctx.Error(apperror.NewClientError(fmt.Errorf("invalid id")))
		return
	}
	stockMutation, err := handler.usecase.FindByID(ctx.Request.Context(), id, uId)
	if err != nil {
		err = ctx.Error(err)
		return
	}
	responseStockMutation := stockmutationres.StockMutationResponse{}
	responseStockMutation.Set(stockMutation)
	response := responses.DefaultResponse{Data: gin.H{"stock_mutation": responseStockMutation}}
	ctx.JSON(http.StatusOK, response)
}

func (handler *StockMutationHandler) CreateRequestStockMutation(ctx *gin.Context) {
	uId := ctx.Value(enums.UserIdKey.Key).(uint64)
	var request stockmutationreq.CreateStockMutationRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		err = ctx.Error(apperror.NewValidationError(err))
		return
	}
	stockMutation := request.ToStockMutation()
	createdRequest, err := handler.usecase.CreateRequestStockMutation(ctx.Request.Context(), &stockMutation, uId)
	if err != nil {
		err = ctx.Error(err)
		return
	}
	responseRequest := stockmutationres.StockMutationResponse{}
	responseRequest.Set(createdRequest)
	response := responses.DefaultResponse{Data: responseRequest}
	ctx.JSON(http.StatusCreated, response)
}

func (handler *StockMutationHandler) Delete(ctx *gin.Context) {
	uId := ctx.Value(enums.UserIdKey.Key).(uint64)
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		err = ctx.Error(err)
		return
	}
	err = handler.usecase.Delete(ctx.Request.Context(), id, uId)
	if err != nil {
		err = ctx.Error(err)
		return
	}
	response := responses.DefaultResponse{Message: "success delete stock mutation request"}
	ctx.JSON(http.StatusOK, response)
}

func (handler *StockMutationHandler) FindAll(ctx *gin.Context) {
	var request requests.GlobalFilter
	var err error
	uId := ctx.Value(enums.UserIdKey.Key).(uint64)
	request.StartDate, err = time.Parse(enums.DDMMYY, time.Time{}.UTC().Format(enums.DDMMYY))
	if err != nil {
		err = ctx.Error(err)
		return
	}
	request.EndDate, err = time.Parse(enums.DDMMYY, time.Now().UTC().Format(enums.DDMMYY))
	if err != nil {
		err = ctx.Error(err)
		return
	}
	err = ctx.ShouldBindQuery(&request)
	if err != nil {
		err = ctx.Error(err)
		return
	}
	qry := requests.NewQuery(&request)
	qry.AddCondition("stock_mutations.updated_at", requests.GreaterEqual, request.StartDate)
	qry.AddCondition("stock_mutations.updated_at", requests.Less, request.EndDate.Add(24*time.Hour))
	stockMutations, pagination, err := handler.usecase.FindAll(ctx.Request.Context(), qry, uId, request.Pharmacy)
	if err != nil {
		err = ctx.Error(err)
		return
	}
	resp := make([]*stockmutationres.StockMutationResponse, 0)
	for _, stockMutation := range stockMutations {
		respStockMutation := stockmutationres.StockMutationResponse{}
		respStockMutation.Set(stockMutation)
		resp = append(resp, &respStockMutation)
	}
	if len(resp) == 0 {
		resp = nil
	}
	pagination.Items = resp
	ctx.JSON(http.StatusOK, responses.DefaultResponse{
		Data: responses.NewPagination(pagination, "stock_mutations"),
	})
}

func (handler *StockMutationHandler) Update(ctx *gin.Context) {
	uId := ctx.Value(enums.UserIdKey.Key).(uint64)
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		err = ctx.Error(err)
		return
	}
	var request stockmutationreq.CreateStockMutationRequest
	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		err = ctx.Error(err)
		return
	}
	stockMutation := request.ToStockMutation()
	stockMutation.ID = id
	switch request.Action {
	case "accept":
		stockMutation.StatusMutationId = enums.Accepted
	case "cancel":
		stockMutation.StatusMutationId = enums.Canceled
	case "reject":
		stockMutation.StatusMutationId = enums.Rejected
	case "edit":
		stockMutation.StatusMutationId = enums.Pending
	}
	edited, err := handler.usecase.Update(ctx.Request.Context(), &stockMutation, uId)
	if err != nil {
		err = ctx.Error(err)
		return
	}
	responseStockMutation := stockmutationres.StockMutationResponse{}
	responseStockMutation.Set(edited)
	response := responses.DefaultResponse{Data: responseStockMutation}
	ctx.JSON(http.StatusOK, response)
}

func NewStockMutationHandler(usecase stockmutationusecaseinterface.StockMutationUsecase) *StockMutationHandler {
	return &StockMutationHandler{usecase: usecase}
}

var _ stockmutationhandlerinterface.StockMutationHandler = &StockMutationHandler{}
