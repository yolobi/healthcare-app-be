package checkoutres

import (
	"healthcare-capt-america/entities/models"
)

type CheckoutResponse struct {
	PharmacyID uint64         `json:"pharmacy_id"`
	TotalPrice float64        `json:"total_price"`
	Carts      []CartResponse `json:"carts"`
}

func GetTotalPrice(carts []CartResponse) float64 {
	var sum float64
	for _, cart := range carts {
		sum += cart.Price
	}
	return sum
}

type CartResponse struct {
	ID             uint64  `json:"id"`
	PharmacyDrugID uint64  `json:"pharmacy_drug_id"`
	DrugID         uint64  `json:"drug_id"`
	DrugName       string  `json:"drug_name"`
	DrugImage      string  `json:"drug_image"`
	SellingUnit    float64 `json:"selling_unit"`
	Quantity       int     `json:"quantity"`
	Weight         float64 `json:"weight"`
	Height         float64 `json:"height"`
	Length         float64 `json:"length"`
	Width          float64 `json:"width"`
	Price          float64 `json:"price"`
}

func NewCheckoutResponse(pharmacyID uint64, carts []models.Cart) *CheckoutResponse {
	var cartResponses []CartResponse
	for _, cart := range carts {
		cartResponses = append(cartResponses, CartResponse{ID: cart.ID, DrugID: cart.Drug.ID, DrugName: cart.Drug.Name, Quantity: cart.Quantity})
	}
	return &CheckoutResponse{
		PharmacyID: pharmacyID,
		Carts:      cartResponses,
	}
}
