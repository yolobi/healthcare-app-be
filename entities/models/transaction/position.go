package transaction

import "github.com/shopspring/decimal"

type Position struct {
	Longitude decimal.Decimal
	Latitude  decimal.Decimal
}

func NewPosition(long, lat decimal.Decimal) Position {
	return Position{
		Longitude: long,
		Latitude:  lat,
	}
}
