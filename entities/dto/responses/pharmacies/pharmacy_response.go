package pharmacies

import (
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/utils"
)

type PharmacyResponse struct {
	ID              uint64                `json:"id"`
	Name            string                `json:"name"`
	PharmacistName  string                `json:"pharmacist_name"`
	LicenseNumber   string                `json:"license_number"`
	PhoneNumber     string                `json:"phone_number"`
	AdminPharmacyId uint64                `json:"admin_pharmacy_id"`
	AdminName       string                `json:"admin_name"`
	Address         AddressResponse       `json:"address"`
	Operationals    []OperationalResponse `json:"operationals"`
}

type OperationalResponse struct {
	Day       string `json:"day"`
	OpenTime  string `json:"open_time"`
	CloseTime string `json:"close_time"`
}

type AddressResponse struct {
	ID         uint64  `json:"id"`
	Detail     string  `json:"detail"`
	ProvinceID uint64  `json:"province_id"`
	Province   string  `json:"province"`
	CityID     uint64  `json:"city_id"`
	City       string  `json:"city"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}

func NewPharmacyResponse(pharmacy *models.Pharmacy) *PharmacyResponse {
	var operationals []OperationalResponse
	for _, operational := range pharmacy.Operationals {
		operationals = append(operationals, OperationalResponse{
			Day:       operational.Day,
			OpenTime:  utils.TimeParseString(operational.OpenTime),
			CloseTime: utils.TimeParseString(operational.CloseTime),
		})
	}
	if operationals == nil {
		operationals = []OperationalResponse{}
	}
	longtitude, _ := pharmacy.Address.Longtitude.Float64()
	latitude, _ := pharmacy.Address.Latitude.Float64()
	return &PharmacyResponse{
		ID:             pharmacy.ID,
		Name:           pharmacy.Name,
		PharmacistName: pharmacy.PharmaciestName,
		LicenseNumber:  pharmacy.LicenseNumber,
		PhoneNumber:    pharmacy.PhoneNumber,
		AdminName:      pharmacy.AdminPharmacies.User.Name,
		Address: AddressResponse{
			ID:         pharmacy.Address.ID,
			Detail:     pharmacy.Address.Detail,
			ProvinceID: pharmacy.Address.ProvinceID,
			Province:   pharmacy.Address.Province.Name,
			CityID:     pharmacy.Address.CityID,
			City:       pharmacy.Address.City.Name,
			Latitude:   latitude,
			Longitude:  longtitude,
		},
		Operationals: operationals,
	}
}

func NewPharmacyResponses(pharmacies []models.Pharmacy) []PharmacyResponse {
	var responses []PharmacyResponse
	for _, pharmacy := range pharmacies {
		responses = append(responses, *NewPharmacyResponse(&pharmacy))
	}
	return responses
}
