package productusecase

import (
	"context"
	"errors"
	"fmt"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/dto/responses/drugsres"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/entities/models/transaction"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/repositories/masterrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/productrepositoryinterface"
	"healthcare-capt-america/interfaces/usecases/productusecaseinterface"
	"healthcare-capt-america/services"
	"healthcare-capt-america/utils"
	"mime/multipart"
	"strings"

	"github.com/shopspring/decimal"
)

type drugUsecase struct {
	drugRepo        productrepositoryinterface.DrugRepository
	categoryRepo    masterrepositoryinterface.CategoryRepository
	formRepo        masterrepositoryinterface.FormRepository
	manufactureRepo masterrepositoryinterface.ManufactureRepository
	addressRepo     masterrepositoryinterface.AddrressRepository
}

func NewDrugUsecase(drugRepo productrepositoryinterface.DrugRepository, categoryRepo masterrepositoryinterface.CategoryRepository, formRepo masterrepositoryinterface.FormRepository, manufactureRepo masterrepositoryinterface.ManufactureRepository, ar masterrepositoryinterface.AddrressRepository) *drugUsecase {
	return &drugUsecase{
		drugRepo:        drugRepo,
		categoryRepo:    categoryRepo,
		formRepo:        formRepo,
		manufactureRepo: manufactureRepo,
		addressRepo:     ar,
	}
}

func (usecase *drugUsecase) EditDrug(ctx context.Context, drug *models.Drug) (*models.Drug, error) {
	res, err := usecase.drugRepo.FindByID(ctx, drug.ID)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, apperror.NewClientError(fmt.Errorf("can't find drug with id %d", drug.ID))
	}
	a := ctx.Value(enums.ProductImageKey.Key).(*multipart.FileHeader)
	path, err := services.UploadProductImage(ctx, a, drug.Name)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	if path != nil {
		drug.Image = *path
	}
	return usecase.drugRepo.Update(ctx, drug)
}

func (usecase *drugUsecase) FindByID(ctx context.Context, drug_id uint64) (*drugsres.DrugMaster, error) {
	drug, err := usecase.drugRepo.FindByID(ctx, drug_id)
	if err != nil {
		return nil, err
	}
	if drug == nil {
		return nil, apperror.NewClientError(fmt.Errorf("can't find drug with id %d", drug_id))
	}
	return drug, nil
}

func (usecase *drugUsecase) UpdateDrug(ctx context.Context, drug_id uint64, status string) error {
	drug, err := usecase.drugRepo.FindByIDUpdate(ctx, drug_id)
	if err != nil {
		return err
	}
	if drug == nil {
		return apperror.NewClientError(fmt.Errorf("can't find drug with id %d", drug_id))
	}
	_, err = usecase.drugRepo.Update(ctx, drug)
	if err != nil {
		return err
	}
	return nil
}

func (usecase *drugUsecase) FindAllDrug(ctx context.Context, qry *requests.GlobalQuery) ([]drugsres.DrugMaster, *responses.Pagination, error) {
	return usecase.drugRepo.FindAll(ctx, qry)
}

func (usecase *drugUsecase) CreateDrug(ctx context.Context, drug *models.Drug) (*models.Drug, error) {
	a := ctx.Value(enums.ProductImageKey.Key).(*multipart.FileHeader)
	extName, err := utils.GetFileType(a)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	fileName := strings.ReplaceAll(strings.ToLower(drug.Name), " ", "_") + extName
	path, err := services.UploadProductImage(ctx, a, fileName)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	drug.Image = *path
	drug, err = usecase.drugRepo.Save(ctx, drug)
	if err != nil {
		return nil, err
	}
	return drug, nil
}

func (usecase *drugUsecase) checkForeignKey(ctx context.Context, drug *models.Drug) error {
	noForeignKey := errors.New("error foreign key")
	_, err := usecase.categoryRepo.FindByID(ctx, drug.CategoryID)
	if err != nil {
		return noForeignKey
	}
	_, err = usecase.formRepo.FindByID(ctx, drug.FormID)
	if err != nil {
		return noForeignKey
	}
	_, err = usecase.manufactureRepo.FindByID(ctx, drug.ManufactureID)
	if err != nil {
		return noForeignKey
	}
	return nil
}

func greaterThanZero(drug *models.Drug) error {
	zero := decimal.NewFromInt(0)
	height := drug.Height.LessThanOrEqual(zero)
	weight := drug.Weight.LessThanOrEqual(zero)
	length := drug.Length.LessThanOrEqual(zero)
	width := drug.Width.LessThanOrEqual(zero)
	if height || weight || length || width {
		err := errors.New("must greater than zero")
		return err
	}
	return nil
}

func (du *drugUsecase) FindAllWithDist(ctx context.Context) ([]*drugsres.DrugUserRes, error) {
	uid := ctx.Value(enums.UserIdKey).(uint64)
	address, err := du.addressRepo.FindByUserId(ctx, uid)
	if err != nil {
		return nil, err
	}
	if address == nil {
		return nil, apperror.NewClientError(fmt.Errorf("user doesn't have default address"))
	}
	drugs, err := du.drugRepo.FindAllWithDist(ctx, transaction.NewPosition(address.Longtitude, address.Latitude))
	if err != nil {
		return nil, err
	}
	return drugs, nil
}

var _ productusecaseinterface.DrugUsecase = &drugUsecase{}
