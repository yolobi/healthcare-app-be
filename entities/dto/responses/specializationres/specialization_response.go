package specializationres

import "healthcare-capt-america/entities/models"

type SpecializationResponse struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

func NewSpecialization(specialization *models.Specialization) *SpecializationResponse {
	return &SpecializationResponse{
		ID:   specialization.ID,
		Name: specialization.Name,
	}
}

func NewSpecializations(specializations []models.Specialization) (result []*SpecializationResponse) {
	for _, specialization := range specializations {
		res := NewSpecialization(&specialization)
		result = append(result, res)
	}
	if result == nil {
		return []*SpecializationResponse{}
	}
	return result
}
