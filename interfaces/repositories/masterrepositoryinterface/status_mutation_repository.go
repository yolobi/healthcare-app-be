package masterrepositoryinterface

import (
	"context"
	"healthcare-capt-america/entities/models"
)

type StatusMutationRepository interface {
	Save(ctx context.Context, form *models.StatusMutation) (*models.StatusMutation, error)
	FindByID(ctx context.Context, form_id uint64) (*models.StatusMutation, error)
}
