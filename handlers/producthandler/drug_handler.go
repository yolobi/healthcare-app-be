package producthandler

import (
	"context"
	"fmt"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/requests/drugsreq"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/dto/responses/drugsres"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/handlers/producthandlerinterface"
	"healthcare-capt-america/interfaces/usecases/productusecaseinterface"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DrugHandler struct {
	usecase productusecaseinterface.DrugUsecase
}

func NewDrugHandler(usecase productusecaseinterface.DrugUsecase) *DrugHandler {
	return &DrugHandler{usecase: usecase}
}

func (handler *DrugHandler) CreateDrug(ctx *gin.Context) {
	var request drugsreq.CreateDrugRequest
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.Error(apperror.NewValidationError(err))
		return
	}
	fh, err := ctx.FormFile(enums.ProductImage)
	if err != nil {
		ctx.Error(apperror.NewValidationError(err))
		return
	}
	if fh == nil {
		ctx.Error(apperror.NewClientError(fmt.Errorf("image is required")))
		return
	}
	ctx2 := ctx.Request.Context()
	ctx3 := context.WithValue(ctx2, enums.ProductImageKey.Key, fh)
	ctx.Request = ctx.Request.WithContext(ctx3)
	drug := request.ToDrug()
	createdDrug, err := handler.usecase.CreateDrug(ctx.Request.Context(), &drug)
	if err != nil {
		ctx.Error(err)
		return
	}
	responseDrug := drugsres.DrugResponse{}
	responseDrug.Set(createdDrug)
	response := responses.DefaultResponse{Data: gin.H{"product": responseDrug}}
	ctx.JSON(http.StatusCreated, response)
}

func (handler *DrugHandler) EditDrug(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.Error(apperror.NewClientError(err))
		return
	}
	var request drugsreq.CreateDrugRequest
	err = ctx.ShouldBind(&request)
	if err != nil {
		ctx.Error(apperror.NewValidationError(err))
		return
	}
	fh, _ := ctx.FormFile(enums.ProductImage)
	ctx2 := ctx.Request.Context()
	ctx3 := context.WithValue(ctx2, enums.ProductImageKey, fh)
	ctx.Request = ctx.Request.WithContext(ctx3)
	drug := request.ToDrug()
	drug.ID = id
	editedDrug, err := handler.usecase.EditDrug(ctx.Request.Context(), &drug)
	if err != nil {
		ctx.Error(err)
		return
	}
	responseDrug := drugsres.DrugResponse{}
	responseDrug.Set(editedDrug)
	response := responses.DefaultResponse{Data: gin.H{"product": responseDrug}}
	ctx.JSON(http.StatusCreated, response)
}

func (handler *DrugHandler) FindAllDrug(ctx *gin.Context) {
	var request requests.GlobalFilter
	err := ctx.ShouldBindQuery(&request)
	if err != nil {
		ctx.Error(apperror.NewServerError(err))
		return
	}
	qry := requests.NewQuery(&request)
	drugs, pagination, err := handler.usecase.FindAllDrug(ctx, qry)
	if err != nil {
		ctx.Error(err)
		return
	}
	var drugs2 []drugsres.DrugMaster
	for _, drug := range drugs {
		image := drug.Image
		drug.Image = image
		drugs2 = append(drugs2, drug)
	}
	pagination.Items = drugs2
	ctx.JSON(http.StatusOK, responses.DefaultResponse{
		Data: responses.NewPagination(pagination, "products"),
	})
}

func (handler *DrugHandler) FindByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.Error(apperror.NewClientError(err))
		return
	}
	drug, err := handler.usecase.FindByID(ctx, id)
	if err != nil {
		ctx.Error(err)
		return
	}
	response := responses.DefaultResponse{Data: gin.H{"product": drug}}
	ctx.JSON(http.StatusCreated, response)
}

func (dh *DrugHandler) FindAllWithDist(ctx *gin.Context) {
	var (
		userId = ctx.Value(enums.UserIdKey.Key).(uint64)
		ctx2   = context.WithValue(ctx.Request.Context(), enums.UserIdKey, userId)
	)
	drugs, err := dh.usecase.FindAllWithDist(ctx2)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, responses.DefaultResponse{
		Data: gin.H{"products": drugs},
	})
}

var _ producthandlerinterface.DrugHandler = &DrugHandler{}
