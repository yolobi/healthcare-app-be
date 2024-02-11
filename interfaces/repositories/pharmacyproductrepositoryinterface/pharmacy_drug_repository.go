package pharmacyproductrepositoryinterface

import (
	"context"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/models"
)

type PharmacyDrugRepository interface {
	Save(context.Context, *models.PharmacyDrug) (*models.PharmacyDrug, error)
	Update(context.Context, *models.PharmacyDrug) (*models.PharmacyDrug, error)
	FindByID(context.Context, uint64, uint64) (*models.PharmacyDrug, error)
	FindByPharmacyID(context.Context, *requests.GlobalQuery, uint64) ([]models.PharmacyDrug, *responses.Pagination, error)
	Delete(context.Context, *models.PharmacyDrug) error
	CountDrugStock(ctx context.Context, drugId uint64) (uint64, error)
	CheckDrugIfExists(ctx context.Context, drugId uint64) (bool, error)
	FindByAdminPharmacyID(ctx context.Context, qry *requests.GlobalQuery, adminId uint64) ([]models.PharmacyDrug, *responses.Pagination, error)
	FindPharmacyIDCheckout(ctx context.Context, address *models.Address, cart []*models.Cart) (uint64, error)
}
