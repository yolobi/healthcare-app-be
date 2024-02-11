package pharmacyproductusecaseinterface

import (
	"context"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/models"
)

type PharmacyDrugUsecase interface {
	CreatePharmacyDrug(context.Context, *models.PharmacyDrug) (*models.PharmacyDrug, error)
	EditPharmacyDrug(context.Context, *models.PharmacyDrug) (*models.PharmacyDrug, error)
	FindAllByPharmacyID(context.Context, *requests.GlobalQuery, uint64) ([]models.PharmacyDrug, *responses.Pagination, error)
	FindByID(context.Context, uint64, uint64) (*models.PharmacyDrug, error)
	DeletePharmacyDrug(context.Context, uint64, uint64) error
	FindAllByLoginAdminId(ctx context.Context, qry *requests.GlobalQuery, userId uint64) ([]models.PharmacyDrug, *responses.Pagination, error)
}
