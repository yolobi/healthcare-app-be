package operationalres

import "healthcare-capt-america/entities/models"

type OperationalResponse struct {
	PharmacyID   uint64   `json:"pharmacy_id"`
	PharmacyName string   `json:"pharmacy_name"`
	Days         []string `json:"days"`
}

func OperationalResponseConvert(operationals []models.Operational) *OperationalResponse {
	var days []string
	for _, operational := range operationals {
		days = append(days, operational.Day)
	}
	return &OperationalResponse{
		PharmacyID:   operationals[0].PharmacyID,
		PharmacyName: operationals[0].Pharmacy.Name,
		Days:         days,
	}
}
