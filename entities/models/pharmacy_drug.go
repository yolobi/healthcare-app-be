package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type PharmacyDrug struct {
	ID          uint64          `gorm:"column:id;primaryKey;autoIncrement"`
	DrugId      uint64          `gorm:"column:drug_id;not null;index:pharmacy_drug_index,unique"`
	PharmacyId  uint64          `gorm:"pharmacy_id;not null;index:pharmacy_drug_index,unique"`
	Stock       int             `gorm:"column:stock;not null"`
	SellingUnit decimal.Decimal `gorm:"column:selling_unit;not null;type:numeric"`
	Status      string          `gorm:"column:status;not null"`
	Drug        Drug            `gorm:"foreignKey:DrugId"`
	Pharmacy    Pharmacy        `gorm:"foreignKey:PharmacyId"`
	CreatedAt   time.Time       `gorm:"column:created_at;not null"`
	UpdatedAt   time.Time       `gorm:"column:updated_at;not null"`
	DeletedAt   gorm.DeletedAt  `gorm:"index"`
}
