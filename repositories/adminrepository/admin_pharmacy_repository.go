package adminrepository

import (
	"context"
	"errors"
	"fmt"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/entities/models/transaction"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/repositories/adminrepositoryinterface"
	"healthcare-capt-america/services"
	"healthcare-capt-america/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type adminPharmacyRepository struct {
	db *gorm.DB
}

func (apr *adminPharmacyRepository) Create(ctx context.Context, user *models.User) (*models.AdminPharmacy, error) {
	var newAdmin *models.AdminPharmacy = &models.AdminPharmacy{}
	err := apr.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		newPass := utils.GenerateToken(8)
		enc, err := bcrypt.GenerateFromPassword([]byte(newPass), enums.DefaultCost)
		if err != nil {
			return apperror.NewServerError(err, "Server Error Encrypting Password")
		}
		pass := string(enc)
		user.Password = &pass
		err = tx.Create(&user).Error
		if err != nil {
			return apperror.NewServerError(err)
		}
		id := &user.ID
		newAdmin.UserId = *id
		err = tx.Create(&newAdmin).Error
		if err != nil {
			return apperror.NewServerError(err)
		}
		err = services.Authority.AssignRoleToUser(user.ID, enums.PharmacyAdmin.Slug)
		if err != nil {
			return apperror.NewServerError(err)
		}
		err = services.SendNewAdminPharmacyEmail(user.Email, newPass)
		if err != nil {
			return apperror.NewServerError(err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return newAdmin, nil
}

func (apr *adminPharmacyRepository) FindAll(ctx context.Context, qry *requests.GlobalQuery) ([]models.AdminPharmacy, *responses.Pagination, error) {
	admins := make([]models.AdminPharmacy, 0)
	limit, offset := qry.GetPagination()
	var role transaction.Role
	err := apr.db.WithContext(ctx).Table("authority_roles").
		Where("name", enums.PharmacyAdmin.Name).
		First(&role).Error
	if err != nil {
		return nil, nil, apperror.NewServerError(err)
	}
	preload := apr.db.WithContext(ctx).
		Model(&models.AdminPharmacy{}).
		Preload("Pharmacies.Address.Province").
		Preload("Pharmacies.Address.City").
		Preload("User")
	for _, condition := range qry.Conditions {
		sql := fmt.Sprintf("%s %s ?", condition.Field, condition.Operation)
		preload.Where(sql, condition.Value)
	}
	err = preload.Limit(limit).Offset(offset).Order(qry.OrderedBy).Find(&admins).Error
	if err != nil {
		return nil, nil, apperror.NewServerError(err)
	}
	var totalItems, totalPages int64
	preload.Offset(-1).Count(&totalItems)
	totalPages = totalItems / int64(qry.PerPage)
	if totalItems%int64(qry.PerPage) != 0 {
		totalPages += 1
	}
	return admins, &responses.Pagination{
		TotalItems:  totalItems,
		TotalPages:  totalPages,
		CurrentPage: int64(qry.Page),
	}, nil
}

func (apr *adminPharmacyRepository) FindByUserId(ctx context.Context, user_id uint64) (*models.AdminPharmacy, error) {
	var admin *models.AdminPharmacy
	err := apr.db.WithContext(ctx).
		Model(&models.AdminPharmacy{}).
		Preload("User").
		Preload("Pharmacies").
		Where("user_id", user_id).
		First(&admin).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, apperror.NewServerError(err)
	}
	return admin, nil
}

func (apr *adminPharmacyRepository) FindBydId(ctx context.Context, id uint64) (*models.AdminPharmacy, error) {
	var admin *models.AdminPharmacy
	err := apr.db.WithContext(ctx).Model(&models.AdminPharmacy{}).
		Where("id", id).
		Preload("Pharmacies.Address.Province").
		Preload("Pharmacies.Address.City").
		Preload("User").
		First(&admin).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, apperror.NewServerError(err)
	}
	return admin, nil
}

func (apr *adminPharmacyRepository) Delete(ctx context.Context, admin *models.AdminPharmacy) error {
	err := apr.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		admin, err := apr.FindByUserId(ctx, admin.UserId)
		if err != nil {
			return err
		}
		if admin == nil {
			return apperror.NewClientError(fmt.Errorf("can't find admin with user id %d", admin.UserId))
		}
		err = tx.Delete(&admin).Error
		if err != nil {
			return apperror.NewServerError(err)
		}
		err = tx.Delete(&models.User{}, admin.UserId).Error
		if err != nil {
			return apperror.NewServerError(err)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func NewAdminPharmacyRepository(db *gorm.DB) *adminPharmacyRepository {
	return &adminPharmacyRepository{
		db: db,
	}
}

var _ adminrepositoryinterface.AdminPharmacyRepository = &adminPharmacyRepository{}
