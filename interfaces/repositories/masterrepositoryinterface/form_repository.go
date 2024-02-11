package masterrepositoryinterface

import (
	"context"
	"healthcare-capt-america/entities/models"
)

type FormRepository interface {
	Save(ctx context.Context, form *models.Form) (*models.Form, error)
	FindByID(ctx context.Context, form_id uint64) (*models.Form, error)
	FindAll(ctx context.Context) ([]models.Form, error)
}
