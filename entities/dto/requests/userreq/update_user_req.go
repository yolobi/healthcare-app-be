package userreq

import (
	"healthcare-capt-america/entities/models/transaction"
	"mime/multipart"
)

type UpdateUserReq struct {
	Name        string               `form:"name" binding:"required"`
	PhoneNumber string               `form:"phone_number" binding:"min=12,max=15,required"`
	Photo       multipart.FileHeader `form:"photo"`
	NewPassword *string              `form:"new_password"`
	OldPassword *string              `form:"old_password"`
}

func (uur *UpdateUserReq) NewUser() *transaction.UpdateUser {
	return &transaction.UpdateUser{
		Name:        uur.Name,
		OldPassword: uur.OldPassword,
		NewPassword: uur.NewPassword,
		PhoneNumber: uur.PhoneNumber,
	}
}
