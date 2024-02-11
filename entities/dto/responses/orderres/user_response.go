package orderres

import "healthcare-capt-america/entities/models"

type UserOrder struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Photo       string `json:"photo"`
}

func newUserOrder(user *models.User) *UserOrder {
	phone := ""
	if user.PhoneNumber != nil {
		phone = *user.PhoneNumber
	}
	return &UserOrder{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: phone,
		Photo:       user.Photo,
	}
}
