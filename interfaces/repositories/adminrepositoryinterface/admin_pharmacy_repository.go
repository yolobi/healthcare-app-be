package adminrepositoryinterface

import (
	"context"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/models"
)

type AdminPharmacyRepository interface {
	Create(ctx context.Context, user *models.User) (*models.AdminPharmacy, error)
	FindAll(ctx context.Context, qry *requests.GlobalQuery) ([]models.AdminPharmacy, *responses.Pagination, error)
	FindByUserId(ctx context.Context, user_id uint64) (*models.AdminPharmacy, error)
	FindBydId(ctx context.Context, id uint64) (*models.AdminPharmacy, error)
	Delete(ctx context.Context, admin *models.AdminPharmacy) error
}
