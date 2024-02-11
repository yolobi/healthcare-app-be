package pharmacyproducthandler

import (
	"context"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/requests/pharmacydrugreq"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/dto/responses/pharmacydrugres"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/handlers/pharmacyproducthandlerinterface"
	"healthcare-capt-america/interfaces/usecases/pharmacyproductusecaseinterface"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PharmacyDrugHandler struct {
	usecase pharmacyproductusecaseinterface.PharmacyDrugUsecase
}

func (handler *PharmacyDrugHandler) EditProductInPharmacy(c *gin.Context) {
	requestId := uuid.NewString()
	ctx := context.WithValue(c.Request.Context(), enums.RequestIdKey, requestId)
	var request *pharmacydrugreq.EditPharmacysDrugRequest
	if err := c.ShouldBind(&request); err != nil {
		c.Error(apperror.NewValidationError(err))
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(apperror.NewClientError(err))
		return
	}
	result, err := handler.usecase.EditPharmacyDrug(ctx, &models.PharmacyDrug{DrugId: uint64(id), PharmacyId: request.PharmacyID, Stock: request.Stock, SellingUnit: request.SellingUnit, Status: request.Status})
	if err != nil {
		c.Error(apperror.NewServerError(err))
		return
	}
	var res pharmacydrugres.PharmacyDrugResponse
	res.Set(result)
	c.JSON(http.StatusOK, responses.DefaultResponse{Data: gin.H{"product": res}})
}

func (handler *PharmacyDrugHandler) GetAllProductsByAdminPharmacy(c *gin.Context) {
	var request requests.GlobalFilter
	err := c.ShouldBindQuery(&request)
	if err != nil {
		c.Error(apperror.NewServerError(err))
		return
	}
	qry := requests.NewQuery(&request)
	userId := c.Value(enums.UserIdKey.Key).(uint64)
	requestId := uuid.NewString()
	ctx := context.WithValue(c.Request.Context(), enums.RequestIdKey, requestId)
	results, pagination, err := handler.usecase.FindAllByLoginAdminId(ctx, qry, userId)
	if err != nil {
		c.Error(err)
		return
	}
	var responseDrugs []pharmacydrugres.PharmacyDrugResponse
	for _, result := range results {
		var respDrug pharmacydrugres.PharmacyDrugResponse
		respDrug.Set(&result)
		responseDrugs = append(responseDrugs, respDrug)
	}
	pagination.Items = responseDrugs
	c.JSON(http.StatusOK, responses.DefaultResponse{Data: responses.NewPagination(pagination, "products")})
}

func NewPharmacyDrugHandler(usecase pharmacyproductusecaseinterface.PharmacyDrugUsecase) *PharmacyDrugHandler {
	return &PharmacyDrugHandler{usecase: usecase}
}

func (handler *PharmacyDrugHandler) CreatePharmacyDrug(ctx *gin.Context) {
	pharmacy_id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		err = ctx.Error(err)
		return
	}
	var request pharmacydrugreq.CreatePharmacyDrugRequest
	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		err = ctx.Error(err)
		return
	}
	pharmacyDrug := request.ToPharmacyDrug()
	pharmacyDrug.PharmacyId = pharmacy_id
	createdPharmacyDrug, err := handler.usecase.CreatePharmacyDrug(ctx, &pharmacyDrug)
	if err != nil {
		err = ctx.Error(err)
		return
	}
	responsePharmacyDrug := pharmacydrugres.PharmacyDrugResponse{}
	responsePharmacyDrug.Set(createdPharmacyDrug)
	response := responses.DefaultResponse{Data: responsePharmacyDrug}
	ctx.JSON(http.StatusCreated, response)
}

func (handler *PharmacyDrugHandler) DeletePharmacyDrug(ctx *gin.Context) {
	pharmacy_id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		err = ctx.Error(err)
		return
	}
	drug_id, err := strconv.ParseUint(ctx.Param("drug_id"), 10, 64)
	if err != nil {
		err = ctx.Error(err)
		return
	}
	err = handler.usecase.DeletePharmacyDrug(ctx, pharmacy_id, drug_id)
	if err != nil {
		err = ctx.Error(err)
		return
	}
	response := responses.DefaultResponse{Message: "success delete drug"}
	ctx.JSON(http.StatusOK, response)
}

func (handler *PharmacyDrugHandler) EditPharmacyDrug(ctx *gin.Context) {
	pharmacy_id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		err = ctx.Error(err)
		return
	}
	drug_id, err := strconv.ParseUint(ctx.Param("drug_id"), 10, 64)
	if err != nil {
		err = ctx.Error(err)
		return
	}
	var request pharmacydrugreq.EditPharmacyDrugRequest
	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		err = ctx.Error(err)
		return
	}
	pharmacyDrug := request.ToPharmacyDrug()
	pharmacyDrug.PharmacyId = pharmacy_id
	pharmacyDrug.DrugId = drug_id
	editedPharmacyDrug, err := handler.usecase.EditPharmacyDrug(ctx, &pharmacyDrug)
	if err != nil {
		err = ctx.Error(err)
		return
	}
	responsePharmacyDrug := pharmacydrugres.PharmacyDrugResponse{}
	responsePharmacyDrug.Set(editedPharmacyDrug)
	response := responses.DefaultResponse{Data: responsePharmacyDrug}
	ctx.JSON(http.StatusOK, response)
}

func (handler *PharmacyDrugHandler) FindAllByPharmacyID(ctx *gin.Context) {
	var request requests.GlobalFilter
	var err error
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
		ctx.Error(apperror.NewServerError(err))
		return
	}
	qry := requests.NewQuery(&request)
	qry.AddCondition("pharmacy_drugs.updated_at", requests.GreaterEqual, request.StartDate)
	qry.AddCondition("pharmacy_drugs.updated_at", requests.Less, request.EndDate.Add(24*time.Hour))
	pharmacy_id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		err = ctx.Error(err)
		return
	}
	pharmacyDrugs, pagination, err := handler.usecase.FindAllByPharmacyID(ctx, qry, pharmacy_id)
	if err != nil {
		err = ctx.Error(err)
		return
	}
	resp := []*pharmacydrugres.PharmacyDrugResponse{}
	for _, pharmacyDrug := range pharmacyDrugs {
		respPharmacyDrug := pharmacydrugres.PharmacyDrugResponse{}
		respPharmacyDrug.Set(&pharmacyDrug)
		resp = append(resp, &respPharmacyDrug)
	}
	pagination.Items = resp
	if len(resp) == 0 {
		pagination.Items = nil
	}
	ctx.JSON(http.StatusOK, responses.DefaultResponse{Data: responses.NewPagination(pagination, "products")})
}

func (handler *PharmacyDrugHandler) FindByID(ctx *gin.Context) {
	pharmacy_id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		err = ctx.Error(err)
		return
	}
	drug_id, err := strconv.ParseUint(ctx.Param("drug_id"), 10, 64)
	if err != nil {
		err = ctx.Error(err)
		return
	}
	pharmacyDrug, err := handler.usecase.FindByID(ctx, pharmacy_id, drug_id)
	if err != nil {
		err = ctx.Error(err)
		return
	}
	responsePharmacyDrug := pharmacydrugres.PharmacyDrugResponse{}
	responsePharmacyDrug.Set(pharmacyDrug)
	ctx.JSON(http.StatusOK, responses.DefaultResponse{
		Data: gin.H{"product": responsePharmacyDrug},
	})
}

var _ pharmacyproducthandlerinterface.PharmacyDrugHandler = &PharmacyDrugHandler{}
