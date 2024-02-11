package masterrepositoryinterface

import (
	"context"
	"healthcare-capt-america/entities/models"
)

type SpecializationRepository interface {
	Save(ctx context.Context, specialization *models.Specialization) (*models.Specialization, error)
	FindByID(ctx context.Context, id uint64) (*models.Specialization, error)
	Find(ctx context.Context) ([]models.Specialization, error)
}
