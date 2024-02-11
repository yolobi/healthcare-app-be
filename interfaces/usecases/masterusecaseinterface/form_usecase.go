package masterusecaseinterface

import (
	"context"
	"healthcare-capt-america/entities/models"
)

type FormUsecase interface {
	GetAllForms(ctx context.Context) ([]models.Form, error)
}
