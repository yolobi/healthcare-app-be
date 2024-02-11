package pharmacyproductusecase

import (
	"context"
	"errors"
	"fmt"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/interfaces/repositories/adminrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/pharmacyproductrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/pharmacyrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/productrepositoryinterface"
	"healthcare-capt-america/interfaces/usecases/pharmacyproductusecaseinterface"
)

//need connection to journal for updating stock
type pharmacyDrugUsecase struct {
	pharmacyDrugRepo pharmacyproductrepositoryinterface.PharmacyDrugRepository
	pharmacyRepo     pharmacyrepositoryinterface.PharmacyRepository
	drugRepo         productrepositoryinterface.DrugRepository
	adminRepo        adminrepositoryinterface.AdminPharmacyRepository
}

func (usecase *pharmacyDrugUsecase) FindAllByLoginAdminId(ctx context.Context, qry *requests.GlobalQuery, userId uint64) ([]models.PharmacyDrug, *responses.Pagination, error) {
	admin, err := usecase.adminRepo.FindByUserId(ctx, userId)
	if err != nil {
		return nil, nil, err
	}
	results, paginations, err := usecase.pharmacyDrugRepo.FindByAdminPharmacyID(ctx, qry, admin.ID)
	if err != nil {
		return nil, nil, err
	}
	return results, paginations, nil
}

func NewPharmacyDrugUsecase(pharmacyDrugRepo pharmacyproductrepositoryinterface.PharmacyDrugRepository, pharmacyRepo pharmacyrepositoryinterface.PharmacyRepository, drugRepo productrepositoryinterface.DrugRepository, adminRepo adminrepositoryinterface.AdminPharmacyRepository) *pharmacyDrugUsecase {
	return &pharmacyDrugUsecase{pharmacyDrugRepo: pharmacyDrugRepo, pharmacyRepo: pharmacyRepo, drugRepo: drugRepo, adminRepo: adminRepo}
}

func (usecase *pharmacyDrugUsecase) EditPharmacyDrug(ctx context.Context, pharmacyDrug *models.PharmacyDrug) (*models.PharmacyDrug, error) {
	get, err := usecase.pharmacyDrugRepo.FindByID(ctx, pharmacyDrug.PharmacyId, pharmacyDrug.DrugId)
	if err != nil {
		return nil, err
	}
	if get == nil {
		return nil, apperror.NewClientError(fmt.Errorf("unknown pharmacy id %d with drug id %d", pharmacyDrug.PharmacyId, pharmacyDrug.DrugId))
	}
	return usecase.pharmacyDrugRepo.Update(ctx, pharmacyDrug)
}

func (usecase *pharmacyDrugUsecase) FindByID(ctx context.Context, pharmacy_id uint64, drug_id uint64) (*models.PharmacyDrug, error) {
	pharmacyDrug, err := usecase.pharmacyDrugRepo.FindByID(ctx, pharmacy_id, drug_id)
	if err != nil {
		return nil, err
	}
	if pharmacyDrug == nil {
		return nil, apperror.NewClientError(fmt.Errorf("can't find pharmacy id %d with drug id %d", pharmacyDrug.PharmacyId, pharmacyDrug.DrugId))
	}
	return pharmacyDrug, nil
}

func (usecase *pharmacyDrugUsecase) FindAllByPharmacyID(ctx context.Context, qry *requests.GlobalQuery, pharmacy_id uint64) ([]models.PharmacyDrug, *responses.Pagination, error) {
	return usecase.pharmacyDrugRepo.FindByPharmacyID(ctx, qry, pharmacy_id)
}

func (usecase *pharmacyDrugUsecase) CreatePharmacyDrug(ctx context.Context, pharmacyDrug *models.PharmacyDrug) (*models.PharmacyDrug, error) {
	err := usecase.checkForeignKey(ctx, pharmacyDrug)
	if err != nil {
		return nil, err
	}
	return usecase.pharmacyDrugRepo.Save(ctx, pharmacyDrug)
}

func (usecase *pharmacyDrugUsecase) DeletePharmacyDrug(ctx context.Context, pharmacy_id uint64, drug_id uint64) error {
	pharmacyDrug, err := usecase.pharmacyDrugRepo.FindByID(ctx, pharmacy_id, drug_id)
	if err != nil {
		return err
	}
	if pharmacyDrug == nil {
		return apperror.NewClientError(fmt.Errorf("can't find pharmacy id %d with drug id %d", pharmacyDrug.PharmacyId, pharmacyDrug.DrugId))
	}
	return usecase.pharmacyDrugRepo.Delete(ctx, pharmacyDrug)
}

func (usecase *pharmacyDrugUsecase) checkForeignKey(ctx context.Context, pharmacyDrug *models.PharmacyDrug) error {
	noForeignKey := errors.New("error foreign key")
	_, err := usecase.pharmacyRepo.FindById(ctx, pharmacyDrug.PharmacyId)
	if err != nil {
		return noForeignKey
	}
	_, err = usecase.drugRepo.FindByID(ctx, pharmacyDrug.DrugId)
	if err != nil {
		return noForeignKey
	}
	return nil
}

var _ pharmacyproductusecaseinterface.PharmacyDrugUsecase = &pharmacyDrugUsecase{}
