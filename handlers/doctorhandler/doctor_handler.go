package doctorhandler

import (
	"context"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/requests/doctorreq"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/dto/responses/doctorres"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/handlers/doctorhandlerinterface"
	"healthcare-capt-america/interfaces/usecases/doctorusecaseinterface"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type doctorHandler struct {
	du doctorusecaseinterface.DoctorUsecase
}

func (dh *doctorHandler) GetAllDoctor(ctx *gin.Context) {
	var request requests.GlobalFilter
	err := ctx.ShouldBindQuery(&request)
	if err != nil {
		ctx.Error(apperror.NewServerError(err))
		return
	}
	qry := requests.NewQuery(&request)
	doctors, pagination, err := dh.du.GetAllDoctor(ctx.Request.Context(), qry)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := make([]*doctorres.DoctorResponse, 0)
	for _, doctor := range doctors {
		resp = append(resp, doctorres.NewDoctorResponse(doctor))
	}
	if len(resp) == 0 {
		ctx.JSON(http.StatusOK, responses.DefaultResponse{Data: gin.H{"doctors": resp}})
		return
	}
	pagination.Items = &resp
	ctx.JSON(http.StatusOK, responses.DefaultResponse{
		Data: responses.NewPagination(pagination, "doctors"),
	})
}

func (dh *doctorHandler) GetCurrentDoctorDetail(ctx *gin.Context) {
	uId := ctx.Value(enums.UserIdKey.Key).(uint64)
	doctor, err := dh.du.GetCurrentDetailDoctor(ctx.Request.Context(), uId)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, responses.DefaultResponse{
		Data: gin.H{"doctor": doctorres.NewDoctorResponse(doctor)},
	})
}

func (dh *doctorHandler) UpdateProfile(ctx *gin.Context) {
	userId := ctx.Value(enums.UserIdKey.Key).(uint64)
	var request doctorreq.UpdateDoctorReq
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.Error(apperror.NewValidationError(err))
		return
	}
	log.Println(request)
	fh, _ := ctx.FormFile(enums.FilePhoto)
	ctx2 := ctx.Request.Context()
	ctx3 := context.WithValue(ctx2, enums.UserPhotoKey, fh)
	ctx.Request = ctx.Request.WithContext(ctx3)
	newDoctor := request.NewDoctor()
	newDoctor.UserId = userId
	updated, err := dh.du.UpdateProfile(ctx.Request.Context(), newDoctor)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, responses.DefaultResponse{
		Data: gin.H{"user": doctorres.NewDoctorResponse(updated)},
	})
}

func (dh *doctorHandler) GetDetailDoctor(ctx *gin.Context) {
	id := ctx.Param("id")
	doctor_id, err := strconv.Atoi(id)
	if err != nil {
		ctx.Error(apperror.NewServerError(err))
		return
	}
	doctor, err := dh.du.GetDetailDoctor(ctx.Request.Context(), uint64(doctor_id))
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, responses.DefaultResponse{
		Data: gin.H{"doctor": doctorres.NewDoctorResponse(doctor)},
	})
}

func (dh *doctorHandler) UpdateStatusDoctor(ctx *gin.Context) {
	var request doctorreq.UpdateStatusDoctor
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.Error(apperror.NewValidationError(err))
		return
	}
	id := ctx.Param("id")
	doctor_id, err := strconv.Atoi(id)
	if err != nil {
		ctx.Error(apperror.NewServerError(err))
		return
	}
	err = dh.du.UpdateStatusDoctor(ctx.Request.Context(), request.IsVerify, uint64(doctor_id))
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, responses.DefaultResponse{
		Message: "Successfully update verify value for the doctor",
	})
}

func NewDoctorHandler(du doctorusecaseinterface.DoctorUsecase) *doctorHandler {
	return &doctorHandler{
		du: du,
	}
}

var _ doctorhandlerinterface.DoctorHandler = &doctorHandler{}
