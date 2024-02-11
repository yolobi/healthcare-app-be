package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        uint64         `gorm:"column:id;primaryKey;autoIncrement"`
	Name      string         `gorm:"column:name;not null;unique"`
	Icon      string         `gorm:"column:icon"`
	CreatedAt time.Time      `gorm:"column:created_at;not null"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
