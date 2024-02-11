package authusecase

import (
	"context"
	"errors"
	"fmt"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/entities/models/transaction"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/repositories/authrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/doctorrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/userrepositoryinterface"
	"healthcare-capt-america/interfaces/usecases/authusecaseinterface"
	"healthcare-capt-america/pkg/security"
	"healthcare-capt-america/services"
	"healthcare-capt-america/utils"
	"mime/multipart"

	"golang.org/x/crypto/bcrypt"
)

type authUsecase struct {
	ar  authrepositoryinterface.AuthRepository
	rtr authrepositoryinterface.RegisterTokenRepository
	ur  userrepositoryinterface.UserRepository
	dr  doctorrepositoryinterface.DoctorRepository
}

func (au *authUsecase) Register(ctx context.Context, user *models.User, role string) error {
	checkUser, err := au.ur.FindByEmail(ctx, user.Email)
	if err != nil {
		return apperror.NewServerError(err, "database error")
	}
	if checkUser != nil {
		if checkUser.Password != nil {
			return apperror.NewClientError(errors.New("email already used by another user"))
		}
		token, err := au.rtr.FindByUserId(ctx, checkUser.ID)
		if err != nil {
			return err
		}
		return au.ar.ResendVerification(ctx, checkUser.Email, token.Token)
	}
	err = au.ar.Register(ctx, user, role)
	return err
}

func (au *authUsecase) VerifyAccount(ctx context.Context, token string, password string) error {
	registerToken, err := au.rtr.FindByToken(ctx, token)
	if err != nil {
		return err
	}
	if registerToken == nil {
		return apperror.NewClientError(errors.New("token not found"))
	}
	if !registerToken.IsValid {
		return apperror.NewClientError(errors.New("token already used"))
	}
	userId := registerToken.UserId
	user, err := au.ur.FindById(ctx, userId)
	if err != nil {
		return err
	}
	if user == nil {
		return apperror.NewClientError(errors.New("can't find user from token"))
	}
	if ok, _ := services.Authority.CheckUserRole(user.ID, enums.Doctor.Slug); ok {
		a := ctx.Value(enums.CertificateKey).(*multipart.FileHeader)
		if a == nil {
			return apperror.NewClientError(fmt.Errorf("certificate is required"))
		}
		path, err := services.UploadCertificate(ctx, a, user.Name+"_"+user.Email)
		if err != nil {
			return apperror.NewServerError(err)
		}

		doctor, err := au.dr.FindByUserId(ctx, user.ID)
		if err != nil {
			return err
		}

		doctor.Certificate = *path
		_, err2 := au.dr.Update(ctx, doctor)
		if err2 != nil {
			return err2
		}
	}
	hashedPass, err := utils.HashPassword(password)
	if err != nil {
		return apperror.NewServerError(err)
	}
	user.Password = &hashedPass
	return au.ar.VerifyAccount(ctx, registerToken, user)
}

func (au *authUsecase) Login(ctx context.Context, auth transaction.Authentication) (*string, *string, error) {
	user, err := au.ur.FindByEmail(ctx, auth.Email)
	if err != nil {
		return nil, nil, err
	}
	if user == nil {
		return nil, nil, apperror.NewClientError(fmt.Errorf("can't find user with email %s", auth.Email))
	}
	if user.Password == nil {
		return nil, nil, apperror.NewClientError(fmt.Errorf("can't find user with email %s", auth.Email))
	}
	errPass := bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(auth.Password))
	if errPass != nil {
		return nil, nil, apperror.NewClientError(fmt.Errorf("invalid email or password"))
	}
	ok, err := services.Authority.CheckUserRole(user.ID, enums.Doctor.Slug)
	if err != nil {
		return nil, nil, apperror.NewServerError(err)
	}
	if ok && !user.IsVerify {
		return nil, nil, apperror.NewClientError(fmt.Errorf("your account is currently being reviewed, please wait until your account is verified by admin"))
	}
	role, err := services.Authority.GetUserRoles(user.ID)
	if err != nil {
		return nil, nil, apperror.NewServerError(err)
	}
	token := security.GenerateJWT(user, role[0].Name)
	return &token, &role[0].Name, nil
}

func NewAuthUsecase(ar authrepositoryinterface.AuthRepository, rtr authrepositoryinterface.RegisterTokenRepository, ur userrepositoryinterface.UserRepository, dr doctorrepositoryinterface.DoctorRepository) *authUsecase {
	return &authUsecase{
		ar:  ar,
		rtr: rtr,
		ur:  ur,
		dr:  dr,
	}
}

var _ authusecaseinterface.AuthUsecase = &authUsecase{}
