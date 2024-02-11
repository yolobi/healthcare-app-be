package models

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	ID        uint64         `gorm:"column:id;primaryKey;autoIncrement"`
	File      string         `gorm:"column:file;not null"`
	Status    string         `gorm:"column:status;not null"`
	CreatedAt time.Time      `gorm:"column:created_at;not null"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
