package provinceusecaseinterface

import (
	"context"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/models"
)

type ProvinceUsecase interface {
	GetProvinceById(ctx context.Context, id uint64) (*models.Province, error)
	GetAllProvinces(ctx context.Context, qry *requests.GlobalQuery) ([]models.Province, *responses.Pagination, error)
}
