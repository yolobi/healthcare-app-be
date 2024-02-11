package doctorusecaseinterface

import (
	"context"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/entities/models/transaction"
)

type DoctorUsecase interface {
	GetAllDoctor(ctx context.Context, qry *requests.GlobalQuery) ([]*models.Doctor, *responses.Pagination, error)
	GetDetailDoctor(ctx context.Context, id uint64) (*models.Doctor, error)
	GetCurrentDetailDoctor(ctx context.Context, id uint64) (*models.Doctor, error)
	UpdateStatusDoctor(ctx context.Context, status bool, id uint64) error
	UpdateProfile(ctx context.Context, doctor *transaction.UpdateDoctor) (*models.Doctor, error)
}
