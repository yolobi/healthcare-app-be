package pharmacyres

import (
	"healthcare-capt-america/entities/models"
)

type PharmacyResponse struct {
	Name            string          `json:"name"`
	PharmaciestName string          `json:"pharmaciest_name"`
	LicenseNumber   string          `json:"pharmaciest_license_number"`
	PhoneNumber     string          `json:"pharmaciest_phone_number"`
	Address         AddressResponse `json:"address"`
}

type AddressResponse struct {
	ID           uint64  `json:"id"`
	Detail       string  `json:"detail"`
	ProvinceID   uint64  `json:"province_id"`
	ProvinceName string  `json:"province"`
	CityID       uint64  `json:"city_id"`
	CityName     string  `json:"city"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
}

type PharmacyDetailResponse struct {
	ID              uint64           `json:"id"`
	Name            string           `json:"name"`
	PharmaciestName string           `json:"pharmaciest_name"`
	LicenseNumber   string           `json:"pharmaciest_license_number"`
	PhoneNumber     string           `json:"pharmaciest_phone_number"`
	OperationalDays []OperationalDay `json:"operational_days"`
}

type OperationalDay struct {
	ID        uint64  `json:"id"`
	Day       string  `json:"day"`
	OpenTime  float64 `json:"open_time"`
	CloseTime float64 `json:"close_time"`
}

func NewPharmacyResponse(pharmacy *models.Pharmacy) *PharmacyResponse {

	long, _ := pharmacy.Address.Longtitude.Float64()
	lat, _ := pharmacy.Address.Latitude.Float64()
	return &PharmacyResponse{
		Name:            pharmacy.Name,
		PharmaciestName: pharmacy.PharmaciestName,
		LicenseNumber:   pharmacy.LicenseNumber,
		PhoneNumber:     pharmacy.PhoneNumber,
		Address: AddressResponse{
			ID:           pharmacy.AddressID,
			Detail:       pharmacy.Address.Detail,
			CityID:       pharmacy.Address.CityID,
			CityName:     pharmacy.Address.City.Name,
			ProvinceID:   pharmacy.Address.ProvinceID,
			ProvinceName: pharmacy.Address.Province.Name,
			Longitude:    long,
			Latitude:     lat,
		},
	}
}

func NewPharmacyDetailResponse(operationals []models.Operational) *PharmacyDetailResponse {
	var operationalDays []OperationalDay
	for _, operational := range operationals {
		operationalDay := &OperationalDay{
			ID:        operational.ID,
			Day:       operational.Day,
			OpenTime:  operational.OpenTime.Hours(),
			CloseTime: operational.CloseTime.Hours(),
		}
		operationalDays = append(operationalDays, *operationalDay)
	}
	operational := operationals[0]
	return &PharmacyDetailResponse{
		ID:              operational.Pharmacy.ID,
		Name:            operational.Pharmacy.Name,
		PharmaciestName: operational.Pharmacy.PharmaciestName,
		LicenseNumber:   operational.Pharmacy.LicenseNumber,
		PhoneNumber:     operational.Pharmacy.PhoneNumber,
		OperationalDays: operationalDays,
	}
}

func NewPharmacyResponses(pharmacies []models.Pharmacy) []PharmacyResponse {
	var responses []PharmacyResponse
	for _, pharmacy := range pharmacies {
		responses = append(responses, *NewPharmacyResponse(&pharmacy))
	}
	return responses
}
