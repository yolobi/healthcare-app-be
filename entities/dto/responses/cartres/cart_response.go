package cartres

import "healthcare-capt-america/entities/models"

type CartResponse struct {
	UserID   uint64              `json:"user_id"`
	UserName string              `json:"user_name"`
	Drugs    []*DrugCartResponse `json:"carts"`
}

type DrugCartResponse struct {
	ID   uint64       `json:"id"`
	Drug DrugResponse `json:"drug"`
}

type DrugResponse struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Image    string `json:"image"`
	Quantity int    `json:"quantity"`
}

func NewCartResponse(carts []*models.Cart) *CartResponse {
	if len(carts) == 0 {
		return &CartResponse{}
	}
	drugs := make([]*DrugCartResponse, 0)
	userId := carts[0].User.ID
	userName := carts[0].User.Name
	for _, cart := range carts {
		res := &DrugCartResponse{
			ID: cart.ID,
			Drug: DrugResponse{
				ID:       cart.Drug.ID,
				Name:     cart.Drug.Name,
				Image:    cart.Drug.Image,
				Quantity: cart.Quantity,
			},
		}
		drugs = append(drugs, res)
	}
	return &CartResponse{
		UserID:   userId,
		UserName: userName,
		Drugs:    drugs,
	}
}
