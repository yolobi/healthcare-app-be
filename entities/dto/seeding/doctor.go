package seeding

import "healthcare-capt-america/entities/models"

type Doctor struct {
	UserId            uint64 `csv:"user_id"`
	Certificate       string `csv:"certificate"`
	YearsOfExperience int    `csv:"years_of_experience"`
	SpecializationId  uint64 `csv:"specialization_id"`
}

func ModelDoctor(inputs []*Doctor) (result []*models.Doctor) {
	for _, input := range inputs {
		var doctor = models.Doctor{}
		doctor.UserId = input.UserId
		doctor.Certificate = input.Certificate
		doctor.YearsOfExperience = input.YearsOfExperience
		doctor.SpecializationId = &input.SpecializationId
		result = append(result, &doctor)
	}
	return
}