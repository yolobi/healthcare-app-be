package masterrepository

import (
	"context"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/repositories/masterrepositoryinterface"
	"healthcare-capt-america/utils"

	"gorm.io/gorm"
)

type subdistrictRepository struct {
	db *gorm.DB
}

func (repo *subdistrictRepository) Save(ctx context.Context, subdistrict *models.SubDistrict) (*models.SubDistrict, error) {
	return utils.SaveQuery[models.SubDistrict](ctx, repo.db, subdistrict, enums.Create)
}

func NewSubdistrictRepository(db *gorm.DB) *subdistrictRepository {
	return &subdistrictRepository{db: db}
}

var _ masterrepositoryinterface.SubDistrictRepository = &subdistrictRepository{}
