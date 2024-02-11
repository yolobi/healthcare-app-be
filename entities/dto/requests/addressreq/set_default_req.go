package addressreq

type SetDefaultReq struct {
	AddressId uint64 `json:"address_id" binding:"required"`
}
