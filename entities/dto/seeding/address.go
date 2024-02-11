package seeding

import (
	"healthcare-capt-america/entities/models"

	"github.com/shopspring/decimal"
)

type Address struct {
	Detail     string          `csv:"detail"`
	ProvinceID uint64          `csv:"province_id"`
	CityID     uint64          `csv:"city_id"`
	Longtitude decimal.Decimal `csv:"longitude"`
	Latitude   decimal.Decimal `csv:"latitude"`
	UserId     uint64          `csv:"user_id"`
}

func ModelAddress(inputs []*Address) (result []*models.Address) {
	for _, input := range inputs {
		var address = models.Address{}
		address.Detail = input.Detail
		address.ProvinceID = input.ProvinceID
		address.CityID = input.CityID
		address.Longtitude = input.Longtitude
		address.Latitude = input.Latitude
		address.UserId = &input.UserId
		if input.UserId == 0 {
			address.UserId = nil
		}
		result = append(result, &address)
	}
	return
}
