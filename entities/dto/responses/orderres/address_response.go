package orderres

import (
	"healthcare-capt-america/entities/models"
)

type AddressOrder struct {
	Detail    string  `json:"detail"`
	Province  string  `json:"province"`
	City      string  `json:"city"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func newAddressOrder(address *models.Address) *AddressOrder {
	lat, _ := address.Latitude.Float64()
	long, _ := address.Longtitude.Float64()

	return &AddressOrder{
		Detail:    address.Detail,
		Province:  address.Province.Name,
		City:      address.City.Name,
		Latitude:  lat,
		Longitude: long,
	}
}
