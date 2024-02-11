package cartreq

import "healthcare-capt-america/entities/models"

type CartRequest struct {
	DrugID   uint64 `json:"drug_id" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
}

func (request *CartRequest) NewCart() *models.Cart {
	return &models.Cart{
		DrugId:   request.DrugID,
		Quantity: request.Quantity,
	}
}
