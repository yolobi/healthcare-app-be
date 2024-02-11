package shipmentreq

import "healthcare-capt-america/entities/models"

type GetAllShipmentReq struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

func NewGetAllReq(ship *models.Shipment) GetAllShipmentReq {
	return GetAllShipmentReq{
		ID:   ship.ID,
		Name: ship.Name,
	}
}
