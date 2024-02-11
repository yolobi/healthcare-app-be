package models

import (
	"time"

	"gorm.io/gorm"
)

type Province struct {
	ID        uint64         `gorm:"column:id;primaryKey;autoIncrement"`
	Name      string         `gorm:"column:name;not null;unique" csv:"name"`
	Cities    []City         `gorm:"foreignKey:ProvinceId"`
	CreatedAt time.Time      `gorm:"column:created_at;not null"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
