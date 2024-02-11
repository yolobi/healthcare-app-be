package pharmacyusecase

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
	"healthcare-capt-america/interfaces/repositories/masterrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/operationalrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/pharmacyrepositoryinterface"
	"healthcare-capt-america/interfaces/usecases/pharmacyusecaseinterface"
)

type pharmacyUsecase struct {
	pharmacyRepo      pharmacyrepositoryinterface.PharmacyRepository
	operationalRepo   operationalrepositoryinterface.OperationalRepository
	adminPharmacyRepo adminrepositoryinterface.AdminPharmacyRepository
	addressRepo       masterrepositoryinterface.AddrressRepository
}

func (usecase *pharmacyUsecase) GetPharmacyByLoginAdminPharmacy(ctx context.Context, userId uint64, qry *requests.GlobalQuery) ([]models.Pharmacy, *responses.Pagination, error) {
	adminPhar, err := usecase.adminPharmacyRepo.FindByUserId(ctx, userId)
	if err != nil {
		return nil, nil, err
	}
	if adminPhar == nil {
		return nil, nil, apperror.NewClientError(fmt.Errorf("you not have pharmacy to handle with your user id is %d", userId))
	}
	pharmacies, pagination, err := usecase.pharmacyRepo.FindByAdminPharmacyId(ctx, adminPhar.ID, qry)
	if err != nil {
		return nil, nil, apperror.NewServerError(err)
	}
	return pharmacies, pagination, nil
}

func (usecase *pharmacyUsecase) GetPharmacyById(ctx context.Context, id uint64) (*models.Pharmacy, error) {
	pharmacy, err := usecase.pharmacyRepo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	if pharmacy == nil {
		return nil, apperror.NewClientError(fmt.Errorf("can't find pharmacy with id %d", id))
	}
	return pharmacy, nil
}

func (usecase *pharmacyUsecase) DeletePharmacyById(ctx context.Context, id uint64) error {
	adminPhar, err := usecase.adminPharmacyRepo.FindByUserId(ctx, ctx.Value(enums.UserIdKey.Key).(uint64))
	if err != nil {
		return err
	}
	if adminPhar == nil {
		return apperror.NewClientError(errors.New("you are not admin"))
	}
	pharmacy, err := usecase.pharmacyRepo.FindById(ctx, id)
	if err != nil {
		return err
	}
	if pharmacy == nil {
		return apperror.NewClientError(fmt.Errorf("can't find pharmacy with id %d", id))
	}
	if pharmacy.AdminPharmacyId != adminPhar.ID {
		return errors.New("you can't delete pharmacy")
	}
	return usecase.pharmacyRepo.Delete(ctx, pharmacy)
}

func (usecase *pharmacyUsecase) UpdatePharmacyById(ctx context.Context, id uint64, editedPharmacyReq *models.Pharmacy, editedAddressReq *models.Address) (*models.Pharmacy, error) {
	adminPhar, err := usecase.adminPharmacyRepo.FindByUserId(ctx, ctx.Value(enums.UserIdKey.Key).(uint64))
	if err != nil {
		return nil, err
	}
	if adminPhar == nil {
		return nil, apperror.NewClientError(errors.New("you are not admin"))
	}
	pharmacy, err := usecase.pharmacyRepo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	if pharmacy == nil {
		return nil, apperror.NewClientError(fmt.Errorf("can't find pharmacy with id %d", id))
	}
	if adminPhar.ID != pharmacy.AdminPharmacyId {
		return nil, apperror.NewClientError(errors.New("you not have access to update this pharmacy"))
	}
	pharmacy = pharmacy.EditRequest(editedPharmacyReq)
	if editedAddressReq != nil {
		address, err := usecase.addressRepo.FindById(ctx, pharmacy.AddressID)
		if err != nil {
			return nil, apperror.NewServerError(err)
		}
		address = address.EditAddressRequest(editedAddressReq)
		_, err = usecase.addressRepo.Update(ctx, address)
		if err != nil {
			return nil, err
		}
	}
	pharmacy, err = usecase.pharmacyRepo.Update(ctx, pharmacy)
	if err != nil {
		return nil, err
	}
	return usecase.pharmacyRepo.FindById(ctx, pharmacy.ID)
}

func (usecase *pharmacyUsecase) GetAllPharmacies(ctx context.Context, qry *requests.GlobalQuery) ([]models.Pharmacy, *responses.Pagination, error) {
	return usecase.pharmacyRepo.FindAll(ctx, qry)
}

func (usecase *pharmacyUsecase) AddPharmacy(ctx context.Context, address *models.Address, pharmacy *models.Pharmacy) (*models.Pharmacy, error) {
	adminPharmacy, err := usecase.adminPharmacyRepo.FindByUserId(ctx, ctx.Value(enums.UserIdKey.Key).(uint64))
	if err != nil {
		return nil, err
	}
	if adminPharmacy == nil {
		return nil, apperror.NewClientError(errors.New("you can't use this action"))
	}
	pharmacy.AdminPharmacyId = adminPharmacy.ID
	pharmacy, err = usecase.pharmacyRepo.AddPharmacy(ctx, address, pharmacy)
	if err != nil {
		return nil, err
	}
	return usecase.pharmacyRepo.FindById(ctx, pharmacy.ID)
}

func NewPharmacyUsecase(pharmacyRepo pharmacyrepositoryinterface.PharmacyRepository, adminPharmacyRepo adminrepositoryinterface.AdminPharmacyRepository, operationalRepo operationalrepositoryinterface.OperationalRepository, addressRepo masterrepositoryinterface.AddrressRepository) *pharmacyUsecase {
	return &pharmacyUsecase{pharmacyRepo: pharmacyRepo, adminPharmacyRepo: adminPharmacyRepo, operationalRepo: operationalRepo, addressRepo: addressRepo}
}

var _ pharmacyusecaseinterface.PharmacyUsecase = &pharmacyUsecase{}
