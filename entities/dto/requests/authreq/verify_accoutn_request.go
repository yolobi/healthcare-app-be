package authreq

import "mime/multipart"

type VerifyAccountRequest struct {
	Token       string               `form:"token" binding:"required"`
	Password    string               `form:"password" binding:"required"`
	Certificate multipart.FileHeader `form:"certificate"`
}
