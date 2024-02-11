package doctorreq

type UpdateStatusDoctor struct {
	IsVerify bool `json:"is_verify" binding:"required"`
}
