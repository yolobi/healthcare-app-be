package masterusecaseinterface

import (
	"context"
	"healthcare-capt-america/entities/models"
)

type SpecializationUsecase interface {
	GetAllSpecializations(ctx context.Context) ([]models.Specialization, error)
}
