package masterusecase

import (
	"context"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/interfaces/repositories/masterrepositoryinterface"
	"healthcare-capt-america/interfaces/usecases/masterusecaseinterface"
)

type formUsecase struct {
	repo masterrepositoryinterface.FormRepository
}

func (f formUsecase) GetAllForms(ctx context.Context) ([]models.Form, error) {
	forms, err := f.repo.FindAll(ctx)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	return forms, nil
}

func NewFormUsecase(repo masterrepositoryinterface.FormRepository) *formUsecase {
	return &formUsecase{repo: repo}
}

var _ masterusecaseinterface.FormUsecase = &formUsecase{}
