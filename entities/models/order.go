package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Order struct {
	ID               uint64          `gorm:"column:id;primaryKey;autoIncrement"`
	UserId           uint64          `gorm:"column:user_id;not null"`
	PharmacyId       uint64          `gorm:"column:pharmacy_id;not null"`
	OderDate         time.Time       `gorm:"column:order_date"`
	AddressId        uint64          `gorm:"column:address_id;not null"`
	ShipmentName     string          `gorm:"column:shipment_name;not null"`
	PaymentId        uint64          `gorm:"column:payment_id;not null"`
	OrderStatus      string          `gorm:"column:order_status;not null"`
	ShippingCost     decimal.Decimal `gorm:"column:shipping_cost;not null"`
	TotalDrugsAmount decimal.Decimal `gorm:"column:total_drugs_amount;not null"`
	TotalAmount      decimal.Decimal `gorm:"column:total_amount;not null"`
	OrderDetails     []OrderDetail   `gorm:"foreignKey:OrderId"`
	User             User            `gorm:"foreignKey:UserId"`
	Pharmacy         Pharmacy        `gorm:"foreignKey:PharmacyId"`
	Address          Address         `gorm:"foreignKey:AddressId"`
	Payment          Payment         `gorm:"foreignKey:PaymentId"`
	CreatedAt        time.Time       `gorm:"column:created_at;not null"`
	UpdatedAt        time.Time       `gorm:"column:updated_at;not null"`
	DeletedAt        gorm.DeletedAt  `gorm:"index"`
}
