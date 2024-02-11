package adminpharmacyreq

import "healthcare-capt-america/entities/models"

type CreateAdminPharmacyRequest struct {
	Name        string `json:"name" binding:"required,min=3"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phone_number" binding:"required,min=12,max=15"`
}

func (capr *CreateAdminPharmacyRequest) ToUser() *models.User {
	return &models.User{
		Name:        capr.Name,
		Email:       capr.Email,
		PhoneNumber: &capr.PhoneNumber,
		Photo:       "public/user_photo/default.jpg",
	}
}
