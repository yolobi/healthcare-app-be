package manufactureres

import "healthcare-capt-america/entities/models"

type ManufactureResponse struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

func NewManufacture(manufacture models.Manufacture) *ManufactureResponse {
	return &ManufactureResponse{ID: manufacture.ID, Name: manufacture.Name}
}

func NewManufactures(manufactures []models.Manufacture) []ManufactureResponse {
	var manufactuRess []ManufactureResponse
	for _, manufacture := range manufactures {
		manufactuRess = append(manufactuRess, *NewManufacture(manufacture))
	}
	return manufactuRess
}
