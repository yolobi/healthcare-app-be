package transaction

import "mime/multipart"

type UpdateDoctor struct {
	UserId            uint64
	Name              string
	PhoneNumber       string
	Photo             multipart.FileHeader
	NewPassword       *string
	OldPassword       *string
	Certificate       string
	YearsOfExperience int
	SpecializationId  uint64
}
