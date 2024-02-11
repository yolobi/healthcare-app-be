package seeding

import (
	"healthcare-capt-america/entities/models"
)

type City struct {
	ID         uint64 `csv:"city_id"`
	Name       string `csv:"name"`
	ProvinceId uint64 `csv:"province_id"`
}

func ModelCity(inputs []*City) (result []*models.City) {
	for _, input := range inputs {
		var city = models.City{}
		city.ProvinceId = input.ProvinceId
		city.ID = input.ID
		city.Name = input.Name
		result = append(result, &city)
	}
	return
}
