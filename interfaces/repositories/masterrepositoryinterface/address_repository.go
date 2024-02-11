package masterrepositoryinterface

import (
	"context"
	"healthcare-capt-america/entities/models"
)

type AddrressRepository interface {
	Save(ctx context.Context, address *models.Address) (*models.Address, error)
	FindByUserId(ctx context.Context, uid uint64) (*models.Address, error)
	FindById(ctx context.Context, id uint64) (*models.Address, error)
	FindAllByUserId(ctx context.Context, uid uint64) ([]*models.Address, error)
	SetDefault(ctx context.Context, address_id uint64) error
	Delete(ctx context.Context, address_id uint64) error
	Update(ctx context.Context, address *models.Address) (*models.Address, error)
}
