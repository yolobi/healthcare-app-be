package doctorreq

import (
	"healthcare-capt-america/entities/models/transaction"
	"mime/multipart"
)

type UpdateDoctorReq struct {
	Name              string               `form:"name" binding:"required"`
	PhoneNumber       string               `form:"phone_number" binding:"min=12,max=15,required"`
	Photo             multipart.FileHeader `form:"photo"`
	NewPassword       *string              `form:"new_password"`
	OldPassword       *string              `form:"old_password"`
	Certificate       string               `form:"certificate" binding:"required"`
	YearsOfExperience int                  `form:"years_of_experience" binding:"required"`
	SpecializationId  uint64               `form:"specialization_id" binding:"required"`
}

func (udr *UpdateDoctorReq) NewDoctor() *transaction.UpdateDoctor {
	return &transaction.UpdateDoctor{
		Name:              udr.Name,
		OldPassword:       udr.OldPassword,
		NewPassword:       udr.NewPassword,
		PhoneNumber:       udr.PhoneNumber,
		Certificate:       udr.Certificate,
		YearsOfExperience: udr.YearsOfExperience,
		SpecializationId:  udr.SpecializationId,
	}
}
