package seeding

import "healthcare-capt-america/entities/models"

type Province struct {
	Name string `csv:"name"`
}

func ModelProvince(inputs []*Province) (result []*models.Province) {
	for _, input := range inputs {
		var province = models.Province{}
		province.Name = input.Name
		result = append(result, &province)
	}
	return
}
