package pharmacyreq

import (
	"healthcare-capt-america/entities/models"

	"github.com/shopspring/decimal"
)

type PharmacyRequest struct {
	Name            string         `json:"name"`
	PharmaciestName string         `json:"pharmacist_name"`
	LicenseNumber   string         `json:"pharmacist_license_number"`
	PhoneNumber     string         `json:"pharmacist_phone_number"`
	Address         AddressRequest `json:"address"`
}

type AddressRequest struct {
	ProvinceID uint64          `json:"province_id"`
	CityID     uint64          `json:"city_id"`
	Longitude  decimal.Decimal `json:"longitude"`
	Latitude   decimal.Decimal `json:"latitude"`
	Detail     string          `json:"detail"`
}

func (request *PharmacyRequest) NewAddress() *models.Address {
	return &models.Address{
		Detail:     request.Address.Detail,
		ProvinceID: request.Address.ProvinceID,
		CityID:     request.Address.CityID,
		Longtitude: request.Address.Longitude,
		Latitude:   request.Address.Latitude,
		UserId:     nil,
	}
}

func (request *PharmacyRequest) NewPharmacy() *models.Pharmacy {
	return &models.Pharmacy{
		Name:            request.Name,
		PharmaciestName: request.PharmaciestName,
		LicenseNumber:   request.LicenseNumber,
		PhoneNumber:     request.PhoneNumber,
	}
}
