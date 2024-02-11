package authreq

import "healthcare-capt-america/entities/models/transaction"

type ResetPasswordRequest struct {
	Token    string `json:"token" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (rpr *ResetPasswordRequest) NewResetPassword() *transaction.ResetPassword {
	return &transaction.ResetPassword{
		Token:    rpr.Token,
		Password: rpr.Password,
	}
}
