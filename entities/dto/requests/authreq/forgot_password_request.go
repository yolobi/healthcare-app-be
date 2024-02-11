package authreq

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required"`
}
