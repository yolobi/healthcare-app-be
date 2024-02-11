package masterrepositoryinterface

import (
	"context"
	"healthcare-capt-america/entities/models"
)

type CityRepository interface {
	Save(ctx context.Context, city *models.City) (*models.City, error)
}
