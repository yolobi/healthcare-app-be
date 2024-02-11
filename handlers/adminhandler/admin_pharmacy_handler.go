package adminhandler

import (
	"context"
	"fmt"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/requests/adminpharmacyreq"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/dto/responses/adminpharmacyres"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/usecases/adminusecaseinterface"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type adminPharmacyHandler struct {
	adminPharmacyUsecase adminusecaseinterface.AdminPharmacyUsecase
}

func NewAdminPharmacyHandler(apu adminusecaseinterface.AdminPharmacyUsecase) *adminPharmacyHandler {
	return &adminPharmacyHandler{
		adminPharmacyUsecase: apu,
	}
}

func (aph *adminPharmacyHandler) CreateAdminPharmacy(ctx *gin.Context) {
	var req adminpharmacyreq.CreateAdminPharmacyRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.Error(apperror.NewValidationError(err))
		return
	}
	admin, err := aph.adminPharmacyUsecase.CreateAdminPharmacy(ctx.Request.Context(), req.ToUser())
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, responses.DefaultResponse{
		Data: adminpharmacyres.CreateAdminResponse{
			AdminId: admin.ID,
		},
	})
}

func (aph *adminPharmacyHandler) GetAllAdmin(ctx *gin.Context) {
	var request requests.GlobalFilter
	err := ctx.ShouldBindQuery(&request)
	if err != nil {
		ctx.Error(apperror.NewServerError(err))
		return
	}
	qry := requests.NewQuery(&request)
	admins, pagination, err := aph.adminPharmacyUsecase.GetAllAdmin(ctx.Request.Context(), qry)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := make([]adminpharmacyres.AdminResponse, 0)
	for _, admin := range admins {
		resp = append(resp, adminpharmacyres.NewAdminResponse(&admin))
	}
	pagination.Items = resp
	ctx.JSON(http.StatusOK, responses.DefaultResponse{
		Data: responses.NewPagination(pagination, "admins"),
	})
}

func (aph *adminPharmacyHandler) DeleteAdminPharmacy(ctx *gin.Context) {
	id := ctx.Param("id")
	admin_id, err := strconv.Atoi(id)
	if err != nil {
		ctx.Error(apperror.NewClientError(fmt.Errorf("invalid id")))
		return
	}
	err = aph.adminPharmacyUsecase.Delete(ctx, uint64(admin_id))
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, responses.DefaultResponse{
		Message: "successfully delete admin pharmcy",
	})
}

func (aph *adminPharmacyHandler) GetDetailAdmin(ctx *gin.Context) {
	id := ctx.Param("id")
	admin_id, err := strconv.Atoi(id)
	if err != nil {
		ctx.Error(apperror.NewClientError(fmt.Errorf("invalid id")))
		return
	}
	admin, err := aph.adminPharmacyUsecase.GetDetailAdmin(ctx, uint64(admin_id))
	if err != nil {
		ctx.Error(err)
		return
	}
	var resp adminpharmacyres.DetailAdmin
	ctx.JSON(http.StatusOK, responses.DefaultResponse{
		Data: gin.H{"admin": resp.NewDetailAdmin(admin)},
	})
}

func (aph *adminPharmacyHandler) UpdateAdmin(ctx *gin.Context) {
	var request adminpharmacyreq.UpdateRequest
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.Error(apperror.NewValidationError(err))
		return
	}
	id := ctx.Param("id")
	admin_id, err := strconv.Atoi(id)
	if err != nil {
		ctx.Error(apperror.NewClientError(fmt.Errorf("invalid id")))
		return
	}
	fh, _ := ctx.FormFile(enums.FilePhoto)
	ctx2 := ctx.Request.Context()
	ctx3 := context.WithValue(ctx2, enums.UserPhotoKey, fh)
	ctx.Request = ctx.Request.WithContext(ctx3)
	err = aph.adminPharmacyUsecase.UpdateAdmin(ctx.Request.Context(), uint64(admin_id), request.NewUser())
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, responses.DefaultResponse{
		Message: "Succesfullly update admin pharmacy data",
	})
}
