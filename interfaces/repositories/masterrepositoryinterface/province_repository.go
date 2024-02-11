package masterrepositoryinterface

import (
	"context"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/models"
)

type ProvinceRepository interface {
	Save(ctx context.Context, province *models.Province) (*models.Province, error)
	FindById(ctx context.Context, id uint64) (*models.Province, error)
	FindAll(ctx context.Context, qry *requests.GlobalQuery) ([]models.Province, *responses.Pagination, error)
}
