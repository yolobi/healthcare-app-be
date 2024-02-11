package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type OrderDetail struct {
	ID             uint64          `gorm:"column:id;primaryKey;autoIncrement"`
	OrderId        uint64          `gorm:"column:order_id;not null"`
	PharmacyDrugId uint64          `gorm:"column:pharmacy_drug_id;not null"`
	Quantity       int             `gorm:"column:quantity;not null;default: 0"`
	Amount         decimal.Decimal `gorm:"column:amount;not null"`
	Order          Order           `gorm:"foreignKey:OrderId"`
	PharmacyDrug   PharmacyDrug    `gorm:"foreignKey:PharmacyDrugId"`
	CreatedAt      time.Time       `gorm:"column:created_at;not null"`
	UpdatedAt      time.Time       `gorm:"column:updated_at;not null"`
	DeletedAt      gorm.DeletedAt  `gorm:"index"`
}
