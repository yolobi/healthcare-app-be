package models

import (
	"time"

	"gorm.io/gorm"
)

type City struct {
	ID         uint64         `gorm:"column:id;primaryKey;autoIncrement" csv:"city_id"`
	Name       string         `gorm:"column:name;not null;unique" csv:"name"`
	ProvinceId uint64         `gorm:"column:province_id;not null" csv:"province_id"`
	Province   Province       `gorm:"foreignKey:ProvinceId"`
	CreatedAt  time.Time      `gorm:"column:created_at;not null"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;not null"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
