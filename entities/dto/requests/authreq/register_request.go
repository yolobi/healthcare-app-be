package authreq

import "healthcare-capt-america/entities/models"

type RegisterRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Role  string `json:"role" binding:"required"`
}

func (rr *RegisterRequest) NewUser() *models.User {
	return &models.User{
		Name:  rr.Name,
		Email: rr.Email,
		Photo: "public/user_photo/default.jpg",
	}
}
