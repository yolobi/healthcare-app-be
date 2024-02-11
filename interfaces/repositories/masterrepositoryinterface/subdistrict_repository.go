package masterrepositoryinterface

import (
	"context"
	"healthcare-capt-america/entities/models"
)

type SubDistrictRepository interface {
	Save(ctx context.Context, subdistrict *models.SubDistrict) (*models.SubDistrict, error)
}
