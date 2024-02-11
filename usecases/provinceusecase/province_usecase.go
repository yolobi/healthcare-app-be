package provinceusecase

import (
	"context"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/interfaces/repositories/masterrepositoryinterface"
	"healthcare-capt-america/interfaces/usecases/provinceusecaseinterface"
)

type provinceUsecase struct {
	provincerepo masterrepositoryinterface.ProvinceRepository
}

func (usecase *provinceUsecase) GetAllProvinces(ctx context.Context, qry *requests.GlobalQuery) ([]models.Province, *responses.Pagination, error) {
	return usecase.provincerepo.FindAll(ctx, qry)
}

func (usecase *provinceUsecase) GetProvinceById(ctx context.Context, id uint64) (*models.Province, error) {
	return usecase.provincerepo.FindById(ctx, id)
}

func NewProvinceUsecase(provincerepo masterrepositoryinterface.ProvinceRepository) *provinceUsecase {
	return &provinceUsecase{
		provincerepo: provincerepo,
	}
}

var _ provinceusecaseinterface.ProvinceUsecase = &provinceUsecase{}
