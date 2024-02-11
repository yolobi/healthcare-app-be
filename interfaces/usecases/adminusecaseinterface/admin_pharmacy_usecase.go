package adminusecaseinterface

import (
	"context"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/models"
)

type AdminPharmacyUsecase interface {
	CreateAdminPharmacy(ctx context.Context, adminPharmacy *models.User) (*models.AdminPharmacy, error)
	GetAllAdmin(ctx context.Context, qry *requests.GlobalQuery) ([]models.AdminPharmacy, *responses.Pagination, error)
	Delete(ctx context.Context, admin_id uint64) error
	GetDetailAdmin(ctx context.Context, id uint64) (*models.AdminPharmacy, error)
	UpdateAdmin(ctx context.Context, admin_id uint64, user *models.User) error
}
