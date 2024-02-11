package orderres

import (
	"healthcare-capt-america/entities/models"
	"time"
)

type OrderResponse struct {
	ID               uint64        `json:"id"`
	User             UserOrder     `json:"user"`
	Pharmacy         PharmacyOrder `json:"pharmacy"`
	OrderDate        time.Time     `json:"order_date"`
	Address          AddressOrder  `json:"address"`
	ShipmentName     string        `json:"shipping_name"`
	Payment          PaymentOrder  `json:"payment"`
	OrderStatus      string        `json:"order_status"`
	ShippingCost     float64       `json:"shipping_cost"`
	TotalDrugsAmount int64         `json:"total_drugs_amount"`
	TotalAmount      float64       `json:"total_amount"`
	Products         []OrderDetail `json:"products"`
	UpdatedAt        time.Time     `json:"updated_at"`
}

func NewOrderResponse(order *models.Order) *OrderResponse {
	products := NewOrdersDetail(order.OrderDetails)
	totalAmount, _ := order.TotalAmount.Float64()
	shippingCost, _ := order.ShippingCost.Float64()
	return &OrderResponse{
		ID:               order.ID,
		User:             *newUserOrder(&order.User),
		Pharmacy:         *newPharmacyOrderResponse(&order.Pharmacy),
		OrderDate:        order.OderDate,
		Address:          *newAddressOrder(&order.Address),
		Payment:          *newPaymentOrder(&order.Payment),
		ShipmentName:     order.ShipmentName,
		OrderStatus:      order.OrderStatus,
		ShippingCost:     shippingCost,
		TotalDrugsAmount: order.TotalDrugsAmount.IntPart(),
		TotalAmount:      totalAmount,
		Products:         products,
		UpdatedAt:        order.UpdatedAt,
	}
}

type OrderDetail struct {
	Id          uint64  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Amount      int64   `json:"amount"`
	Image       string  `json:"image"`
	SellingUnit float64 `json:"selling_unit"`
}

func NewOrderDetail(detail *models.OrderDetail) OrderDetail {
	sellingUnit, _ := detail.PharmacyDrug.SellingUnit.Float64()
	return OrderDetail{
		Id:          detail.PharmacyDrugId,
		Name:        detail.PharmacyDrug.Drug.Name,
		Description: detail.PharmacyDrug.Drug.Description,
		Amount:      int64(detail.Quantity),
		Image:       detail.PharmacyDrug.Drug.Image,
		SellingUnit: sellingUnit,
	}
}

func NewOrdersDetail(details []models.OrderDetail) []OrderDetail {
	var responses []OrderDetail
	for _, detail := range details {
		responses = append(responses, NewOrderDetail(&detail))
	}
	return responses
}

func NewOrderResponses(orders []models.Order) []OrderResponse {
	var responses []OrderResponse
	for _, order := range orders {
		responses = append(responses, *NewOrderResponse(&order))
	}
	return responses
}
