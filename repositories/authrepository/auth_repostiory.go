package authrepository

import (
	"context"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/repositories/authrepositoryinterface"
	"healthcare-capt-america/services"
	"healthcare-capt-america/utils"

	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func (ar *authRepository) Register(ctx context.Context, user *models.User, role string) error {
	err := ar.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		newUser, userErr := utils.SaveQuery[models.User](ctx, tx, user, enums.Create)
		if userErr != nil {
			return apperror.NewClientError(userErr)
		}
		salt := "u"
		if role == enums.Doctor.Name {
			salt = "d"
		}
		token, regErr := utils.SaveQuery[models.RegisterAccountToken](ctx, tx, &models.RegisterAccountToken{
			UserId: newUser.ID,
			Token:  utils.GenerateToken(15) + salt,
		}, enums.Create)
		if regErr != nil {
			return apperror.NewClientError(regErr)
		}
		assignRole := enums.User
		if role == enums.Doctor.Name {
			_, errDoc := utils.SaveQuery[models.Doctor](ctx, tx, &models.Doctor{
				UserId:            user.ID,
				YearsOfExperience: 0,
			}, enums.Create)
			if errDoc != nil {
				return apperror.NewClientError(errDoc)
			}
			assignRole = enums.Doctor
		}
		errEmail := services.SendVerificationEmail(newUser.Email, token.Token)
		if errEmail != nil {
			return apperror.NewServerError(errEmail)
		}
		errAuth := services.Authority.AssignRoleToUser(newUser.ID, assignRole.Slug)
		if errAuth != nil {
			return apperror.NewClientError(errAuth)
		}
		return nil
	})
	return err
}

func (ar *authRepository) ResendVerification(ctx context.Context, email string, token string) error {
	return services.SendVerificationEmail(email, token)
}

func (ar *authRepository) VerifyAccount(ctx context.Context, token *models.RegisterAccountToken, user *models.User) error {
	err := ar.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		token.IsValid = false
		errToken := tx.Save(&token).Error
		if errToken != nil {
			return apperror.NewClientError(errToken)
		}

		ok, err := services.Authority.CheckUserRole(user.ID, enums.User.Slug)
		if err != nil {
			return apperror.NewServerError(err)
		}
		if ok {
			user.IsVerify = true
		}
		errUser := tx.Save(&user).Error
		if errUser != nil {
			apperror.NewClientError(errUser)
		}
		return nil
	})
	return err
}

func NewAuthrepository(db *gorm.DB) *authRepository {
	return &authRepository{
		db: db,
	}
}

var _ authrepositoryinterface.AuthRepository = &authRepository{}
