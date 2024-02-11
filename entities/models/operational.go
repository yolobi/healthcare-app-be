package models

import (
	"time"

	"gorm.io/gorm"
)

type Operational struct {
	ID         uint64         `gorm:"column:id;primaryKey;autoIncrement"`
	PharmacyID uint64         `gorm:"column:pharmacy_id;not null"`
	Day        string         `gorm:"column:day;not null"`
	OpenTime   time.Duration  `gorm:"column:open_time;not null"`
	CloseTime  time.Duration  `gorm:"column:close_time;not null"`
	CreatedAt  time.Time      `gorm:"column:created_at;not null"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;not null"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Pharmacy   Pharmacy       `gorm:"foreignKey:PharmacyID"`
}
