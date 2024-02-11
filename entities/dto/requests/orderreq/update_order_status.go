package orderreq

type UpdateOrderStatusReq struct {
	Status string `json:"order_status" binding:"required"`
}
