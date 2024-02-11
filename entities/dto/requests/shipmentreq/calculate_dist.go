package shipmentreq

type CalculateDistanceReq struct {
	AddressId  uint64 `json:"address_id" binding:"required"`
	PharmacyId uint64 `json:"pharmacy_id" binding:"required"`
	Weight     int64  `json:"weight" binding:"required"`
}
