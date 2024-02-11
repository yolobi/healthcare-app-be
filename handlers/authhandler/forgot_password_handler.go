package authhandler

import (
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/dto/requests/authreq"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/interfaces/handlers/authhandlerinterface"
	"healthcare-capt-america/interfaces/usecases/authusecaseinterface"
	"net/http"

	"github.com/gin-gonic/gin"
)

type forgotPasswordHandler struct {
	fpu authusecaseinterface.ForgotPasswordUsecase
}

func NewForgotPasswordHandler(fpu authusecaseinterface.ForgotPasswordUsecase) *forgotPasswordHandler {
	return &forgotPasswordHandler{
		fpu: fpu,
	}
}

func (fph *forgotPasswordHandler) GetToken(ctx *gin.Context) {
	var request authreq.ForgotPasswordRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.Error(apperror.NewValidationError(err))
		return
	}
	err = fph.fpu.GetToken(ctx, request.Email)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, responses.DefaultResponse{
		Message: "Successfully send reset password email",
	})
}

func (fph *forgotPasswordHandler) ResetPassword(ctx *gin.Context) {
	var request *authreq.ResetPasswordRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.Error(apperror.NewValidationError(err))
		return
	}
	err = fph.fpu.ResetPassword(ctx.Request.Context(), request.NewResetPassword())
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, responses.DefaultResponse{
		Message: "Succesfully reset password",
	})
}

var _ authhandlerinterface.ForgotPasswordHandler = &forgotPasswordHandler{}
