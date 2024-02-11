package provincehandler

import (
	"context"
	"fmt"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/dto/responses/provinceres"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/handlers/provincehandlerinterface"
	"healthcare-capt-america/interfaces/usecases/provinceusecaseinterface"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProvinceHandler struct {
	usecase provinceusecaseinterface.ProvinceUsecase
}

func (handler *ProvinceHandler) GetAllProvinces(c *gin.Context) {
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(c.Request.Context(), enums.RequestIdKey, rID)
		request requests.GlobalFilter
	)
	err := c.ShouldBindQuery(&request)
	if err != nil {
		c.Error(apperror.NewServerError(err))
		return
	}
	qry := requests.NewQuery(&request)
	results, pagination, err := handler.usecase.GetAllProvinces(ctx, qry)
	if err != nil {
		c.Error(err)
		return
	}
	resp := make([]provinceres.ProvinceResponse, 0)
	for _, result := range results {
		resp = append(resp, *provinceres.NewProvinceRes(&result))
	}
	pagination.Items = resp
	c.JSON(http.StatusOK, responses.DefaultResponse{
		Data: responses.NewPagination(pagination, "provinces"),
	})
}

func (handler *ProvinceHandler) GetProvinceById(c *gin.Context) {
	var (
		rID   = uuid.NewString()
		ctx   = context.WithValue(c.Request.Context(), enums.RequestIdKey, rID)
		getId = c.Param("id")
	)
	id, err := strconv.Atoi(getId)
	if err != nil {
		c.Error(apperror.NewClientError(fmt.Errorf("invalid id")))
		return
	}
	result, err := handler.usecase.GetProvinceById(ctx, uint64(id))
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, responses.DefaultResponse{Data: gin.H{"province": provinceres.NewProvinceRes(result)}})
}

func NewProvinceHandler(usecase provinceusecaseinterface.ProvinceUsecase) *ProvinceHandler {
	return &ProvinceHandler{
		usecase: usecase,
	}
}

var _ provincehandlerinterface.ProvinceHandler = &ProvinceHandler{}
