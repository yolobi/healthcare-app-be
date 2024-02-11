package pharmacyrepositoryinterface

import (
	"context"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/models"
)

type PharmacyRepository interface {
	AddPharmacy(ctx context.Context, address *models.Address, pharmacy *models.Pharmacy) (*models.Pharmacy, error)
	FindAll(ctx context.Context, qrt *requests.GlobalQuery) ([]models.Pharmacy, *responses.Pagination, error)
	Update(ctx context.Context, pharmacy *models.Pharmacy) (*models.Pharmacy, error)
	FindById(ctx context.Context, id uint64) (*models.Pharmacy, error)
	Delete(ctx context.Context, pharmacy *models.Pharmacy) error
	FindByAdminPharmacyId(ctx context.Context, adminPharmacyId uint64, qry *requests.GlobalQuery) ([]models.Pharmacy, *responses.Pagination, error)
}
