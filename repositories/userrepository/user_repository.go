package userrepository

import (
	"context"
	"errors"
	"fmt"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/interfaces/repositories/userrepositoryinterface"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func (ur *userRepository) Save(ctx context.Context, user *models.User) (*models.User, error) {
	err := ur.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return nil, apperror.NewServerError(err, "Database Error")
	}
	return user, nil
}

func (ur *userRepository) FindAll(ctx context.Context, qry *requests.GlobalQuery) ([]*models.User, *responses.Pagination, error) {
	users := make([]*models.User, 0)
	limit, offset := qry.GetPagination()
	preload := ur.db.WithContext(ctx).Model(&models.User{}).
		Joins(`JOIN authority_user_roles aur ON "users"."id" = CAST(aur.user_id AS BIGINT) `).Where("aur.role_id = 3")
	for _, condition := range qry.Conditions {
		sql := fmt.Sprintf("%s %s ?", condition.Field, condition.Operation)
		preload.Where(sql, condition.Value)
	}
	searchQuery := fmt.Sprintf("name ILIKE %s OR email ILIKE %s OR phone_number ILIKE %s", qry.Search, qry.Search, qry.Search)
	preload.Where(searchQuery).Order(qry.OrderedBy)
	err := preload.Limit(limit).Offset(offset).Find(&users).Error
	if err != nil {
		return nil, nil, apperror.NewServerError(err, "Database Error")
	}
	var totalItems, totalPages int64
	preload.Offset(-1).Count(&totalItems)
	totalPages = totalItems / int64(qry.PerPage)
	if totalItems%int64(qry.PerPage) != 0 {
		totalPages += 1
	}
	return users, &responses.Pagination{
		TotalItems:  totalItems,
		TotalPages:  totalPages,
		CurrentPage: int64(qry.Page),
	}, nil
}

func (ur *userRepository) FindById(ctx context.Context, id uint64) (*models.User, error) {
	var user *models.User
	err := ur.db.WithContext(ctx).
		Preload("Addresses.Province").
		Preload("Addresses.City").
		Where("id", id).
		First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, apperror.NewServerError(err, "Database Error")
	}
	return user, nil
}

func (ur *userRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user *models.User
	err := ur.db.WithContext(ctx).
		Where("email", email).
		First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, apperror.NewServerError(err, "Database Error")
	}
	return user, nil
}

func (ur *userRepository) Update(ctx context.Context, user *models.User) (*models.User, error) {
	err := ur.db.WithContext(ctx).Save(&user).Error
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	return user, nil
}

func (ur *userRepository) Delete(ctx context.Context, id uint64) error {
	err := ur.db.WithContext(ctx).Delete(&models.User{}, id).Error
	if err != nil {
		return nil
	}
	return nil
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

var _ userrepositoryinterface.UserRepository = &userRepository{}
