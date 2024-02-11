package doctorres

import "healthcare-capt-america/entities/models"

type DoctorResponse struct {
	ID                uint64  `json:"id"`
	Name              string  `json:"name"`
	Email             string  `json:"email"`
	PhoneNumber       *string `json:"phone_number"`
	SpecializationId  *uint64 `json:"specialization_id"`
	Specialization    *string `json:"specialization"`
	YearsOfExperience int     `json:"years_of_experience"`
	Certificate       string  `json:"certificate"`
	Photo             string  `json:"photo"`
	IsVerify          bool    `json:"is_verify"`
}

func NewDoctorResponse(doctor *models.Doctor) *DoctorResponse {
	var special *string
	if doctor.Specialization != nil {
		special = &doctor.Specialization.Name
	}
	var phone_number *string
	if doctor.User.PhoneNumber != nil {
		phone_number = doctor.User.PhoneNumber
	}
	var special_id *uint64
	if doctor.SpecializationId != nil {
		special_id = doctor.SpecializationId
	}

	return &DoctorResponse{
		ID:                doctor.ID,
		Name:              doctor.User.Name,
		Email:             doctor.User.Email,
		PhoneNumber:       phone_number,
		SpecializationId:  special_id,
		Specialization:    special,
		YearsOfExperience: doctor.YearsOfExperience,
		Certificate:       doctor.Certificate,
		Photo:             doctor.User.Photo,
		IsVerify:          doctor.User.IsVerify,
	}
}
