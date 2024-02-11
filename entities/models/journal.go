package models

import (
	"time"

	"gorm.io/gorm"
)

type Journal struct {
	ID             uint64         `gorm:"column:id;primaryKey;autoIncrement"`
	DrugId         uint64         `gorm:"column:drug_id;not null"`
	FromPharmacyId *uint64        `gorm:"column:from_pharmacy_id"`
	ToPharmacyId   uint64         `gorm:"column:to_pharmacy_id;not null"`
	Status         string         `gorm:"column:status;not null"`
	Quantity       int            `gorm:"column:quantity;not null"`
	Drug           Drug           `gorm:"foreignKey:DrugId"`
	FromPharmacy   *Pharmacy      `gorm:"foreignKey:FromPharmacyId"`
	ToPharmacy     Pharmacy       `gorm:"foreignKey:ToPharmacyId"`
	CreatedAt      time.Time      `gorm:"column:created_at;not null"`
	UpdatedAt      time.Time      `gorm:"column:updated_at;not null"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
