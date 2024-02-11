package operationalusecase

import (
	"context"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/interfaces/repositories/operationalrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/pharmacyrepositoryinterface"
	"healthcare-capt-america/interfaces/usecases/operationalusecaseinterface"
)

type operationalUsecase struct {
	pharmacyRepo    pharmacyrepositoryinterface.PharmacyRepository
	operationalRepo operationalrepositoryinterface.OperationalRepository
}

// AddOperationalDays implements operationalusecaseinterface.OperationalUsecase.
func (usecase *operationalUsecase) AddOperationalDays(ctx context.Context, operationals []models.Operational) ([]models.Operational, error) {
	return usecase.operationalRepo.SaveMultiple(ctx, operationals)
}

// DeleteOperationalDay implements operationalusecaseinterface.OperationalUsecase.
func (usecase *operationalUsecase) DeleteOperationalDay(ctx context.Context, operationalId uint64) error {
	return usecase.operationalRepo.Delete(ctx, operationalId)
}

// GetPharmacyOperationalsDay implements operationalusecaseinterface.OperationalUsecase.
func (*operationalUsecase) GetPharmacyOperationalsDay(ctx context.Context, pharmacyId uint64) ([]models.Pharmacy, error) {
	panic("unimplemented")
}

func NewOperationalUsecase(pharmacyRepo pharmacyrepositoryinterface.PharmacyRepository, operationalRepo operationalrepositoryinterface.OperationalRepository) *operationalUsecase {
	return &operationalUsecase{
		pharmacyRepo:    pharmacyRepo,
		operationalRepo: operationalRepo,
	}
}

var _ operationalusecaseinterface.OperationalUsecase = &operationalUsecase{}
