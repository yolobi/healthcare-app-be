package userrepositoryinterface

import (
	"context"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/models"
)

type UserRepository interface {
	Save(ctx context.Context, user *models.User) (*models.User, error)
	FindAll(ctx context.Context, qry *requests.GlobalQuery) ([]*models.User, *responses.Pagination, error)
	FindById(ctx context.Context, id uint64) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	Update(ctx context.Context, user *models.User) (*models.User, error)
	Delete(ctx context.Context, id uint64) error
}
