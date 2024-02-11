package provinceres

import "healthcare-capt-america/entities/models"

type (
	ProvinceResponse struct {
		ID     uint64         `json:"id"`
		Name   string         `json:"name"`
		Cities []CityResponse `json:"cities"`
	}

	CityResponse struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
	}
)

func NewProvinceRes(province *models.Province) *ProvinceResponse {
	var cities []CityResponse
	for _, city := range province.Cities {
		cityRes := &CityResponse{
			ID:   city.ID,
			Name: city.Name,
		}
		cities = append(cities, *cityRes)
	}
	return &ProvinceResponse{
		ID:     province.ID,
		Name:   province.Name,
		Cities: cities,
	}
}
