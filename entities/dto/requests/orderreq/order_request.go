package orderreq

import (
	"github.com/shopspring/decimal"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/enums"
	"time"
)

type OrderRequest struct {
	PharmacyID        uint64  `json:"pharmacy_id" binding:"required"`
	AddressID         uint64  `json:"address_id" binding:"required"`
	ShippingName      string  `json:"shipping_name" binding:"required"`
	ShippingCost      float64 `json:"shipping_cost" binding:"required"`
	TotalProductPrice float64 `json:"total_product_price" binding:"required"`
}

func (request *OrderRequest) NewOrderModel() *models.Order {
	return &models.Order{
		PharmacyId:       request.PharmacyID,
		AddressId:        request.AddressID,
		ShipmentName:     request.ShippingName,
		ShippingCost:     decimal.NewFromFloat(request.ShippingCost),
		TotalDrugsAmount: decimal.NewFromFloat(request.TotalProductPrice),
		TotalAmount:      decimal.NewFromFloat(request.ShippingCost + request.TotalProductPrice),
		OderDate:         time.Now(),
		OrderStatus:      enums.WaitingForPayment,
	}
}
