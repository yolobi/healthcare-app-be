package models

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	ID        uint64         `gorm:"column:id;primaryKey;autoIncrement"`
	DrugId    uint64         `gorm:"column:drug_id;not null"`
	UserId    uint64         `gorm:"column:user_id;not null"`
	Quantity  int            `gorm:"column:quantity;not null;default: 0"`
	Drug      Drug           `gorm:"foreignKey:DrugId"`
	User      User           `gorm:"foreignKey:UserId"`
	CreatedAt time.Time      `gorm:"column:created_at;not null"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
