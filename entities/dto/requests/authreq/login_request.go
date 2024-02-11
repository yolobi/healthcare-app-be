package authreq

import "healthcare-capt-america/entities/models/transaction"

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (lr *LoginRequest) NewAuth() *transaction.Authentication {
	return &transaction.Authentication{
		Email:    lr.Email,
		Password: lr.Password,
	}
}
