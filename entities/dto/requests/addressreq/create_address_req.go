package addressreq

import (
	"healthcare-capt-america/entities/models"

	"github.com/shopspring/decimal"
)

type CreateAddressReq struct {
	Detail     string  `json:"detail" binding:"required"`
	ProvinceID uint64  `json:"province_id" binding:"required"`
	CityID     uint64  `json:"city_id" binding:"required"`
	UserId     uint64  ``
	Longtitude float64 `json:"longitude" binding:"required"`
	Latitude   float64 `json:"latitude" binding:"required"`
}

func (car *CreateAddressReq) NewAddress() *models.Address {
	long := decimal.NewFromFloat(car.Longtitude)
	lat := decimal.NewFromFloat(car.Latitude)
	return &models.Address{
		Detail:     car.Detail,
		ProvinceID: car.ProvinceID,
		CityID:     car.CityID,
		Longtitude: long,
		Latitude:   lat,
		UserId:     &car.UserId,
	}
}
