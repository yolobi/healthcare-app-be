package authhandler

import (
	"context"
	"errors"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/dto/requests/authreq"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/handlers/authhandlerinterface"
	"healthcare-capt-america/interfaces/usecases/authusecaseinterface"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	au authusecaseinterface.AuthUsecase
}

func (ah *authHandler) Register(ctx *gin.Context) {
	var request authreq.RegisterRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.Error(apperror.NewValidationError(err))
		return
	}
	if request.Role != enums.User.Name && request.Role != enums.Doctor.Name {
		ctx.Error(apperror.NewValidationError(errors.New("role should be user or doctor")))
		return
	}
	err = ah.au.Register(ctx.Request.Context(), request.NewUser(), request.Role)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, responses.DefaultResponse{
		Message: "Succesfully Create User, Need to Verify Via Email",
	})
}

func (ah *authHandler) Verify(ctx *gin.Context) {
	var request authreq.VerifyAccountRequest
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.Error(apperror.NewValidationError(err))
		return
	}
	fh, _ := ctx.FormFile(enums.FileCertificate)
	ctx2 := ctx.Request.Context()
	ctx3 := context.WithValue(ctx2, enums.CertificateKey, fh)
	ctx.Request = ctx.Request.WithContext(ctx3)
	err = ah.au.VerifyAccount(ctx.Request.Context(), request.Token, request.Password)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, responses.DefaultResponse{
		Message: "Successfully verify account",
	})
}

func (ah *authHandler) Login(ctx *gin.Context) {
	var request authreq.LoginRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.Error(apperror.NewValidationError(err))
		return
	}
	token, role, err := ah.au.Login(ctx.Request.Context(), *request.NewAuth())
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, responses.DefaultResponse{
		Data: gin.H{
			"token": token,
			"role":  role,
		},
	})
}

func NewAuthHandler(au authusecaseinterface.AuthUsecase) *authHandler {
	return &authHandler{
		au: au,
	}
}

var _ authhandlerinterface.AuthHandler = &authHandler{}
