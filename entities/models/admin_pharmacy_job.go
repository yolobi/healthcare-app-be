package models

import (
	"time"

	"gorm.io/gorm"
)

type AdminPharmacyJob struct {
	ID              uint64         `gorm:"column:id;primaryKey;autoIncrement"`
	AdminPharmacyId uint64         `gorm:"column:admin_pharmacy_id;not null"`
	PharmacyId      uint64         `gorm:"column:pharmacy_id;not null"`
	AdminPharmacy   AdminPharmacy  `gorm:"foreignKey:AdminPharmacyId"`
	Pharmacy        Pharmacy       `gorm:"foreignKey:PharmacyId"`
	CreatedAt       time.Time      `gorm:"column:created_at;not null"`
	UpdatedAt       time.Time      `gorm:"column:updated_at;not null"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
