package userhandler

import (
	"context"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/requests/addressreq"
	"healthcare-capt-america/entities/dto/requests/userreq"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/dto/responses/adminpharmacyres"
	"healthcare-capt-america/entities/dto/responses/userres"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/handlers/userhandlerinterface"
	"healthcare-capt-america/interfaces/usecases/userusecaseinterface"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase userusecaseinterface.UserUsecase
}

func (uh *UserHandler) GetAllUsers(ctx *gin.Context) {
	var request requests.GlobalFilter
	err := ctx.ShouldBindQuery(&request)
	if err != nil {
		ctx.Error(apperror.NewServerError(err))
		return
	}
	qry := requests.NewQuery(&request)
	users, pagination, err := uh.userUsecase.GetAllUsers(ctx.Request.Context(), qry)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := make([]userres.UserDetailResponse, 0)
	for _, user := range users {
		resp = append(resp, userres.NewUserDetail(user))
	}
	pagination.Items = resp
	ctx.JSON(http.StatusOK, responses.DefaultResponse{
		Data: responses.NewPagination(pagination, "users"),
	})
}

func (uh *UserHandler) GetUserDetail(ctx *gin.Context) {
	id := ctx.Param("id")
	user_id, err := strconv.Atoi(id)
	if err != nil {
		ctx.Error(apperror.NewServerError(err))
		return
	}
	user, err := uh.userUsecase.GetUserDetail(ctx.Request.Context(), uint64(user_id))
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, responses.DefaultResponse{
		Data: gin.H{"user": userres.NewUserDetail(user)},
	})
}

func (uh *UserHandler) GetCurrentUserDetail(ctx *gin.Context) {
	userId := ctx.Value(enums.UserIdKey.Key).(uint64)
	user, err := uh.userUsecase.GetUserDetail(ctx.Request.Context(), userId)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, responses.DefaultResponse{
		Data: gin.H{"user": userres.NewUserDetail(user)},
	})
}

func (uh *UserHandler) UpdateProfile(ctx *gin.Context) {
	userId := ctx.Value(enums.UserIdKey.Key).(uint64)
	var request userreq.UpdateUserReq
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.Error(apperror.NewValidationError(err))
		return
	}
	fh, _ := ctx.FormFile(enums.FilePhoto)
	ctx2 := ctx.Request.Context()
	ctx3 := context.WithValue(ctx2, enums.UserPhotoKey, fh)
	ctx.Request = ctx.Request.WithContext(ctx3)
	newUser := request.NewUser()
	newUser.ID = userId
	updated, err := uh.userUsecase.UpdateProfile(ctx.Request.Context(), newUser)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, responses.DefaultResponse{
		Data: gin.H{"user": userres.NewUserDetail(updated)},
	})
}

func (uh *UserHandler) AddAddress(ctx *gin.Context) {
	var request addressreq.CreateAddressReq
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.Error(apperror.NewValidationError(err))
		return
	}
	uid := ctx.Value(enums.UserIdKey.Key).(uint64)
	request.UserId = uid
	address, err := uh.userUsecase.AddAddress(ctx, request.NewAddress())
	if err != nil {
		ctx.Error(err)
		return
	}
	add := adminpharmacyres.AddressRes{}
	ctx.JSON(http.StatusOK, responses.DefaultResponse{
		Data: gin.H{"address": add.NewAddressRes(*address)},
	})
}

func (uh *UserHandler) FindAllUserAddress(ctx *gin.Context) {
	uid := ctx.Value(enums.UserIdKey.Key).(uint64)
	adrresses, err := uh.userUsecase.FindAllUserAddress(ctx, uid)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := make([]*adminpharmacyres.AddressRes, 0)
	for _, address := range adrresses {
		add := adminpharmacyres.AddressRes{}
		ar := add.NewAddressRes(*address)
		resp = append(resp, &ar)
	}
	ctx.JSON(http.StatusOK, responses.DefaultResponse{
		Data: gin.H{"addresses": resp},
	})
}

func (uh *UserHandler) SetDefaultAddress(ctx *gin.Context) {
	var (
		userId  = ctx.Value(enums.UserIdKey.Key).(uint64)
		ctx2    = context.WithValue(ctx.Request.Context(), enums.UserIdKey, userId)
		request addressreq.SetDefaultReq
	)

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.Error(apperror.NewValidationError(err))
		return
	}

	err = uh.userUsecase.SetDefaultAddress(ctx2, request.AddressId)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, responses.DefaultResponse{
		Message: "Succesfully set default address",
	})
}

func (uh *UserHandler) DeleteAddress(ctx *gin.Context) {
	var (
		userId  = ctx.Value(enums.UserIdKey.Key).(uint64)
		ctx2    = context.WithValue(ctx.Request.Context(), enums.UserIdKey, userId)
		request addressreq.SetDefaultReq
	)
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.Error(apperror.NewValidationError(err))
		return
	}
	err = uh.userUsecase.DeleteAddress(ctx2, request.AddressId)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, responses.DefaultResponse{
		Message: "Successfully delete address",
	})
}

func (uh *UserHandler) UpdateAddress(ctx *gin.Context) {
	var (
		userId  = ctx.Value(enums.UserIdKey.Key).(uint64)
		ctx2    = context.WithValue(ctx.Request.Context(), enums.UserIdKey, userId)
		request addressreq.UpdateAddressReq
	)
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.Error(apperror.NewValidationError(err))
		return
	}
	a, err := uh.userUsecase.UpdateAddress(ctx2, request.NewAddress())
	if err != nil {
		ctx.Error(err)
		return
	}
	add := adminpharmacyres.AddressRes{}
	resp := add.NewAddressRes(*a)
	ctx.JSON(http.StatusOK, responses.DefaultResponse{
		Data: gin.H{"address": resp},
	})
}

var _ userhandlerinterface.UserHandler = &UserHandler{}

func NewUserHandler(uu userusecaseinterface.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: uu,
	}
}
