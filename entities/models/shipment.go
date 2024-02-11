package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Shipment struct {
	ID        uint64          `gorm:"column:id;primaryKey;autoIncrement"`
	Name      string          `gorm:"column:name;not null;unique"`
	CostPerKM decimal.Decimal `gorm:"column:cost_per_km;not null"`
	CreatedAt time.Time       `gorm:"column:created_at;not null"`
	UpdatedAt time.Time       `gorm:"column:updated_at;not null"`
	DeletedAt gorm.DeletedAt  `gorm:"index"`
}
