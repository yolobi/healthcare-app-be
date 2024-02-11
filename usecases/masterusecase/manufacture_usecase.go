package masterusecase

import (
	"context"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/interfaces/repositories/masterrepositoryinterface"
	"healthcare-capt-america/interfaces/usecases/masterusecaseinterface"
)

type manufactureUsecase struct {
	repo masterrepositoryinterface.ManufactureRepository
}

func (m manufactureUsecase) GetAllManufactures(ctx context.Context) ([]models.Manufacture, error) {
	manufactures, err := m.repo.FindAll(ctx)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	return manufactures, nil
}

func NewManufactureUsecase(repo masterrepositoryinterface.ManufactureRepository) *manufactureUsecase {
	return &manufactureUsecase{repo: repo}
}

var _ masterusecaseinterface.ManufactureUsecase = &manufactureUsecase{}
