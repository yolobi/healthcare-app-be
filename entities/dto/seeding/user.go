package seeding

import (
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/utils"
)

type User struct {
	Name        string  `csv:"name"`
	Email       string  `csv:"email"`
	IsVerify    bool    `csv:"is_verify"`
	PhoneNumber *string `csv:"phone_number"`
	Photo       string  `csv:"photo"`
}

func ModelUser(inputs []*User) (result []*models.User) {
	password := hotHash("Test4321.")
	for _, input := range inputs {
		var user = models.User{}
		user.Name = input.Name
		user.Email = input.Email
		user.Password = password
		user.IsVerify = input.IsVerify
		user.PhoneNumber = input.PhoneNumber
		user.Photo = input.Photo
		result = append(result, &user)
	}
	return
}

func hotHash(pass string) *string {
	res, _ := utils.HashPassword(pass)
	return &res
}
