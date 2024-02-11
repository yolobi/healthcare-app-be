package doctorrepositoryinterface

import (
	"context"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/models"
)

type DoctorRepository interface {
	FindAll(ctx context.Context, qry *requests.GlobalQuery) ([]*models.Doctor, *responses.Pagination, error)
	FindByUserId(ctx context.Context, user_id uint64) (*models.Doctor, error)
	FindById(ctx context.Context, id uint64) (*models.Doctor, error)
	Update(ctx context.Context, doctor *models.Doctor) (*models.Doctor, error)
	UpdateProfile(ctx context.Context, user *models.User, doctor *models.Doctor) (*models.Doctor, error)
}
