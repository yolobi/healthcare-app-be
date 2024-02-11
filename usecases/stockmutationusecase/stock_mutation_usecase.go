package stockmutationusecase

import (
	"context"
	"errors"
	"fmt"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/repositories/adminrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/pharmacyproductrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/pharmacyrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/stockmutationrepositoryinterface"
	"healthcare-capt-america/interfaces/usecases/stockmutationusecaseinterface"
)

type stockMutationUsecase struct {
	stockMutationRepo stockmutationrepositoryinterface.StockMutationRepository
	pharmacyDrugRepo  pharmacyproductrepositoryinterface.PharmacyDrugRepository
	pharmacyRepo      pharmacyrepositoryinterface.PharmacyRepository
	AdminPharmacyRepo adminrepositoryinterface.AdminPharmacyRepository
}

func (usecase *stockMutationUsecase) FindByID(ctx context.Context, stockMutation_id uint64, uId uint64) (*models.StockMutation, error) {
	stockMutation, err := usecase.stockMutationRepo.FindByID(ctx, stockMutation_id)
	if err != nil {
		return nil, err
	}
	if stockMutation == nil {
		return nil, apperror.NewClientError(fmt.Errorf("stock mutation with id %d is not found", stockMutation_id))
	}
	if err = usecase.validateFromToAdminId(ctx, stockMutation.FromPharmacyId, stockMutation.ToPharmacyId, uId); err != nil {
		return nil, err
	}
	return stockMutation, nil
}

func NewStockMutationUsecase(stockMutationRepo stockmutationrepositoryinterface.StockMutationRepository,
	pharmacyDrugRepo pharmacyproductrepositoryinterface.PharmacyDrugRepository,
	pharmacyRepo pharmacyrepositoryinterface.PharmacyRepository,
	AdminPharmacyRepo adminrepositoryinterface.AdminPharmacyRepository) *stockMutationUsecase {
	return &stockMutationUsecase{stockMutationRepo: stockMutationRepo,
		pharmacyDrugRepo:  pharmacyDrugRepo,
		pharmacyRepo:      pharmacyRepo,
		AdminPharmacyRepo: AdminPharmacyRepo}
}

func (usecase *stockMutationUsecase) CreateRequestStockMutation(ctx context.Context, stockMutation *models.StockMutation, uId uint64) (*models.StockMutation, error) {
	if stockMutation.FromPharmacyId == stockMutation.ToPharmacyId {
		return nil, apperror.NewClientError(fmt.Errorf("can't mutation to same pharmacy"))
	}
	err := usecase.validateStock(ctx, stockMutation)
	if err != nil {
		return nil, err
	}
	if err = usecase.validateAdminId(ctx, stockMutation.ToPharmacyId, uId); err != nil {
		return nil, err
	}
	stockMutation.StatusMutationId = enums.Pending
	return usecase.stockMutationRepo.CreateRequestStockMutation(ctx, stockMutation)
}

func (usecase *stockMutationUsecase) Delete(ctx context.Context, stockMutation_id uint64, uId uint64) error {
	stockMutation, err := usecase.stockMutationRepo.FindByID(ctx, stockMutation_id)
	if err != nil {
		return err
	}
	if stockMutation == nil {
		return apperror.NewClientError(fmt.Errorf("stock mutation with id %d is not found", stockMutation_id))
	}
	if err = usecase.validateAdminId(ctx, stockMutation.ToPharmacyId, uId); err != nil {
		return err
	}
	return usecase.stockMutationRepo.Delete(ctx, stockMutation)
}

func (usecase *stockMutationUsecase) FindAll(ctx context.Context, qry *requests.GlobalQuery, uId uint64, action string) ([]*models.StockMutation, *responses.Pagination, error) {
	admin, err := usecase.AdminPharmacyRepo.FindByUserId(ctx, uId)
	if err != nil {
		return nil, nil, err
	}
	if admin == nil {
		return nil, nil, apperror.NewClientError(errors.New("you can't use this action"))
	}
	return usecase.stockMutationRepo.FindAll(ctx, qry, admin.ID, action)
}

func (usecase *stockMutationUsecase) Update(ctx context.Context, stockMutation *models.StockMutation, uId uint64) (*models.StockMutation, error) {
	data, err := usecase.stockMutationRepo.FindByID(ctx, stockMutation.ID)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, apperror.NewClientError(fmt.Errorf("stock mutation with id %d is not found", stockMutation.ID))
	}
	if err = usecase.validateFromToAdminId(ctx, stockMutation.FromPharmacyId, stockMutation.ToPharmacyId, uId); err != nil {
		return nil, err
	}
	if stockMutation.StatusMutationId == enums.Accepted {
		journal := &models.Journal{
			DrugId:         stockMutation.DrugId,
			FromPharmacyId: &stockMutation.FromPharmacyId,
			ToPharmacyId:   stockMutation.ToPharmacyId,
			Quantity:       stockMutation.Quantity,
			Status:         "Transfer",
		}
		stockMutation.CreatedAt = data.CreatedAt
		return usecase.stockMutationRepo.AcceptRequestStockMutation(ctx, stockMutation, journal)
	} else if stockMutation.StatusMutationId != enums.Pending {
		stockMutation.DrugId = data.DrugId
		stockMutation.ToPharmacyId = data.ToPharmacyId
		stockMutation.FromPharmacyId = data.FromPharmacyId
		stockMutation.Quantity = data.Quantity
	}
	return usecase.stockMutationRepo.Update(ctx, stockMutation)
}

func (usecase *stockMutationUsecase) validateStock(ctx context.Context, stockMutation *models.StockMutation) error {
	noStock := errors.New("insufficient stock")
	fromPharmacy, err := usecase.pharmacyDrugRepo.FindByID(ctx, stockMutation.FromPharmacyId, stockMutation.DrugId)
	if err != nil {
		return err
	}
	if fromPharmacy == nil {
		return apperror.NewClientError(fmt.Errorf("can't find pharmacy with id %d and drug id %d", stockMutation.FromPharmacyId, stockMutation.DrugId))
	}
	if fromPharmacy.Stock-stockMutation.Quantity < 0 {
		return apperror.NewClientError(noStock)
	}
	return nil
}

func (usecase *stockMutationUsecase) validateFromToAdminId(ctx context.Context, from_pharmacy_id uint64, to_pharmacy_id uint64, uId uint64) error {
	admin, err := usecase.AdminPharmacyRepo.FindByUserId(ctx, uId)
	if err != nil {
		return err
	}
	if admin == nil {
		return apperror.NewClientError(errors.New("you can't use this action"))
	}
	from_pharmacy, err := usecase.pharmacyRepo.FindById(ctx, from_pharmacy_id)
	if err != nil {
		return err
	}
	to_pharmacy, err := usecase.pharmacyRepo.FindById(ctx, to_pharmacy_id)
	if err != nil {
		return err
	}
	if from_pharmacy == nil || to_pharmacy == nil {
		return apperror.NewClientError(fmt.Errorf("can't find pharmacy with id %d", from_pharmacy_id))
	}
	if from_pharmacy.AdminPharmacyId != admin.ID && to_pharmacy.AdminPharmacyId != admin.ID {
		return apperror.NewClientError(fmt.Errorf("you can't handle mutation for pharmacy id %d", from_pharmacy_id))
	}
	return nil
}

func (usecase *stockMutationUsecase) validateAdminId(ctx context.Context, pharmacy_id uint64, uId uint64) error {
	admin, err := usecase.AdminPharmacyRepo.FindByUserId(ctx, uId)
	if err != nil {
		return err
	}
	if admin == nil {
		return apperror.NewClientError(errors.New("you can't use this action"))
	}
	pharmacy, err := usecase.pharmacyRepo.FindById(ctx, pharmacy_id)
	if err != nil {
		return err
	}
	if pharmacy == nil {
		return apperror.NewClientError(fmt.Errorf("can't find pharmacy with id %d", pharmacy_id))
	}
	if pharmacy.AdminPharmacyId != admin.ID {
		return apperror.NewClientError(fmt.Errorf("you can't handle mutation for pharmacy id %d", pharmacy_id))
	}
	return nil
}

var _ stockmutationusecaseinterface.StockMutationUsecase = &stockMutationUsecase{}
