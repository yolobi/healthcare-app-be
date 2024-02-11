package masterrepository

import (
	"context"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/repositories/masterrepositoryinterface"
	"healthcare-capt-america/utils"

	"gorm.io/gorm"
)

type cityRepository struct {
	db *gorm.DB
}

// Save implements repositories.CityRepository.
func (repo *cityRepository) Save(ctx context.Context, city *models.City) (*models.City, error) {
	return utils.SaveQuery[models.City](ctx, repo.db, city, enums.Create)
}

func NewCityRepository(db *gorm.DB) *cityRepository {
	return &cityRepository{db: db}
}

var _ masterrepositoryinterface.CityRepository = &cityRepository{}
