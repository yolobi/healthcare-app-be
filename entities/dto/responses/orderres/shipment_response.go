package orderres

import "github.com/shopspring/decimal"

type ShipmentOrder struct {
	ShippingName string  `json:"name"`
	ShippingCost float64 `json:"cost_per_km"`
}

func newShipmentOrder(name string, cost decimal.Decimal) *ShipmentOrder {
	shipCost, _ := cost.Float64()
	return &ShipmentOrder{
		ShippingName: name,
		ShippingCost: shipCost,
	}
}
