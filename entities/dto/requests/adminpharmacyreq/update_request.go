package adminpharmacyreq

import (
	"healthcare-capt-america/entities/models"
	"mime/multipart"
)

type UpdateRequest struct {
	Name        string               `form:"name" binding:"required"`
	PhoneNumber string               `form:"phone_number" binding:"required"`
	Photo       multipart.FileHeader `form:"photo"`
}

func (ur *UpdateRequest) NewUser() *models.User {
	return &models.User{
		Name:        ur.Name,
		PhoneNumber: &ur.PhoneNumber,
	}
}
