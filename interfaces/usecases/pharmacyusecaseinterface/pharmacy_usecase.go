package pharmacyusecaseinterface

import (
	"context"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/models"
)

type PharmacyUsecase interface {
	AddPharmacy(ctx context.Context, address *models.Address, pharmacy *models.Pharmacy) (*models.Pharmacy, error)
	GetAllPharmacies(ctx context.Context, qry *requests.GlobalQuery) ([]models.Pharmacy, *responses.Pagination, error)
	GetPharmacyById(ctx context.Context, id uint64) (*models.Pharmacy, error)
	UpdatePharmacyById(ctx context.Context, id uint64, pharmacy *models.Pharmacy, address *models.Address) (*models.Pharmacy, error)
	DeletePharmacyById(ctx context.Context, id uint64) error
	GetPharmacyByLoginAdminPharmacy(ctx context.Context, userId uint64, qry *requests.GlobalQuery) ([]models.Pharmacy, *responses.Pagination, error)
}
