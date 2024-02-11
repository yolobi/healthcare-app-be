package masterusecase

import (
	"context"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/interfaces/repositories/masterrepositoryinterface"
	"healthcare-capt-america/interfaces/usecases/masterusecaseinterface"
)

type specializationUsecase struct {
	repo masterrepositoryinterface.SpecializationRepository
}

func (s *specializationUsecase) GetAllSpecializations(ctx context.Context) ([]models.Specialization, error) {
	return s.repo.Find(ctx)
}

func NewSpecializationUsecase(repo masterrepositoryinterface.SpecializationRepository) *specializationUsecase {
	return &specializationUsecase{repo: repo}
}

var _ masterusecaseinterface.SpecializationUsecase = &specializationUsecase{}
