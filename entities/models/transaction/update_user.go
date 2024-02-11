package transaction

import "mime/multipart"

type UpdateUser struct {
	ID          uint64
	Name        string
	PhoneNumber string
	Photo       multipart.FileHeader
	NewPassword *string
	OldPassword *string
}
