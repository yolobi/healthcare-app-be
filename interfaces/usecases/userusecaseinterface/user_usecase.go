package userusecaseinterface

import (
	"context"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/entities/models/transaction"
)

type UserUsecase interface {
	GetAllUsers(ctx context.Context, qry *requests.GlobalQuery) ([]*models.User, *responses.Pagination, error)
	GetUserDetail(ctx context.Context, user_id uint64) (*models.User, error)
	UpdateProfile(ctx context.Context, user *transaction.UpdateUser) (*models.User, error)
	AddAddress(ctx context.Context, address *models.Address) (*models.Address, error)
	FindAllUserAddress(ctx context.Context, uid uint64) ([]*models.Address, error)
	SetDefaultAddress(ctx context.Context, address_id uint64) error
	DeleteAddress(ctx context.Context, address_id uint64) error
	UpdateAddress(ctx context.Context, address *models.Address) (*models.Address, error)
}
