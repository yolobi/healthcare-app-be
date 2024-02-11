package pharmacyhandler

import (
	"context"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/requests/pharmacyreq"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/dto/responses/adminpharmacyres"
	"healthcare-capt-america/entities/dto/responses/pharmacies"
	"healthcare-capt-america/entities/dto/responses/pharmacyres"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/handlers/pharmacyhandlerinterface"
	"healthcare-capt-america/interfaces/usecases/pharmacyusecaseinterface"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PharmacyHandler struct {
	usecase pharmacyusecaseinterface.PharmacyUsecase
}

func (h *PharmacyHandler) GetPharmaciesByLoginAdminPharmacy(c *gin.Context) {
	var (
		userId  = c.Value(enums.UserIdKey.Key).(uint64)
		ctx     = context.WithValue(c.Request.Context(), enums.UserIdKey, userId)
		request requests.GlobalFilter
	)
	if err := c.ShouldBindQuery(&request); err != nil {
		c.Error(apperror.NewServerError(err))
		return
	}
	qry := requests.NewQuery(&request)
	pharmacies, pagination, err := h.usecase.GetPharmacyByLoginAdminPharmacy(ctx, userId, qry)
	if err != nil {
		c.Error(err)
		return
	}
	var pharmaciesRes []adminpharmacyres.PharmacyRes
	for _, pharmacy := range pharmacies {
		pharmacyRes := adminpharmacyres.PharmacyRes{}
		pharmaciesRes = append(pharmaciesRes, pharmacyRes.NewPharmacyRes(pharmacy))
	}
	if len(pharmaciesRes) == 0 {
		pagination.Items = nil
	}
	pagination.Items = pharmaciesRes
	c.JSON(http.StatusOK,
		responses.DefaultResponse{
			Data: responses.NewPagination(pagination, "pharmacies"),
		},
	)
}

func (h *PharmacyHandler) GetPharmacyById(c *gin.Context) {
	var (
		requestId = uuid.NewString()
		id, err   = strconv.Atoi(c.Param("id"))
		ctx       = context.WithValue(c.Request.Context(), "request_id", requestId)
	)
	if err != nil {
		c.Error(apperror.NewClientError(err))
		return
	}
	res, err := h.usecase.GetPharmacyById(ctx, uint64(id))
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, responses.DefaultResponse{Data: gin.H{"pharmacy": pharmacies.NewPharmacyResponse(res)}})
}

func (h *PharmacyHandler) DeletePharmacy(c *gin.Context) {
	var (
		id  = c.Param("id")
		uId = c.Value(enums.UserIdKey.Key).(uint64)
		ctx = context.WithValue(c.Request.Context(), enums.UserIdKey.Key, uId)
	)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.Error(apperror.NewClientError(err))
		return
	}
	err = h.usecase.DeletePharmacyById(ctx, uint64(idInt))
	if err != nil {
		c.Error(err)
		return
	}
	defaultResponse := &responses.DefaultResponse{
		Message: "Success to delete",
	}
	c.JSON(http.StatusOK, defaultResponse)
}

func (h *PharmacyHandler) UpdatePharmacy(c *gin.Context) {
	var (
		request *pharmacyreq.PharmacyRequest
		ctx     = context.WithValue(c.Request.Context(), enums.UserIdKey.Key, c.Value(enums.UserIdKey.Key).(uint64))
		id      = c.Param("id")
	)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.Error(apperror.NewClientError(err))
		return
	}
	if err := c.ShouldBind(&request); err != nil {
		c.Error(err)
		return
	}
	pharmacy := request.NewPharmacy()
	address := request.NewAddress()
	pharmacy, err = h.usecase.UpdatePharmacyById(ctx, uint64(idInt), pharmacy, address)
	if err != nil {
	c.Error(err)
		return
	}
	c.JSON(http.StatusOK, responses.DefaultResponse{Data: gin.H{"pharmacy": pharmacyres.NewPharmacyResponse(pharmacy)}})
}

// AddPharmacy implements handlers.PharmacyHandler.
func (h *PharmacyHandler) AddPharmacy(c *gin.Context) {
	var (
		uId     = c.Value(enums.UserIdKey.Key).(uint64)
		request *pharmacyreq.PharmacyRequest
		ctx     = context.WithValue(c.Request.Context(), enums.UserIdKey.Key, uId)
	)
	if err := c.ShouldBind(&request); err != nil {
		c.Error(apperror.NewClientError(err))
		return
	}
	address := request.NewAddress()
	address.IsDefault = false
	pharmacy := request.NewPharmacy()
	pharmacy, err := h.usecase.AddPharmacy(ctx, address, pharmacy)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, responses.DefaultResponse{Data: gin.H{"pharmacy": pharmacyres.NewPharmacyResponse(pharmacy)}})
}

// GetPharmacies implements handlers.PharmacyHandler.
func (h *PharmacyHandler) GetPharmacies(c *gin.Context) {
	var (
		requestId = uuid.NewString()
		ctx       = context.WithValue(c.Request.Context(), "request_id", requestId)
		request   requests.GlobalFilter
	)
	if err := c.ShouldBindQuery(&request); err != nil {
		c.Error(apperror.NewServerError(err))
		return
	}
	qry := requests.NewQuery(&request)
	pharmacies, pagination, err := h.usecase.GetAllPharmacies(ctx, qry)
	if err != nil {
		c.Error(err)
		return
	}
	var pharmaciesRes []adminpharmacyres.PharmacyRes
	for _, pharmacy := range pharmacies {
		pharmacyRes := adminpharmacyres.PharmacyRes{}
		pharmaciesRes = append(pharmaciesRes, pharmacyRes.NewPharmacyRes(pharmacy))
	}
	pagination.Items = pharmaciesRes
	c.JSON(http.StatusOK,
		responses.DefaultResponse{
			Data: responses.NewPagination(pagination, "pharmacies"),
		},
	)
}

func NewPharmacyHandler(usecase pharmacyusecaseinterface.PharmacyUsecase) *PharmacyHandler {
	return &PharmacyHandler{usecase: usecase}
}

var _ pharmacyhandlerinterface.PharmacyHandler = &PharmacyHandler{}
