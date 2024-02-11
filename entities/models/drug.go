package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Drug struct {
	ID            uint64          `gorm:"column:id;primaryKey;autoIncrement"`
	Content       string          `gorm:"column:content;not null" csv:"content"`
	Name          string          `gorm:"column:name;not null;unique" csv:"name"`
	GenericName   string          `gorm:"column:generic_name;not null" csv:"generic_name"`
	Description   string          `gorm:"column:description;not null" csv:"description"`
	ManufactureID uint64          `gorm:"column:manufacture_id;not null" csv:"manufacture_id"`
	FormID        uint64          `gorm:"column:form_id;not null" csv:"form_id"`
	CategoryID    uint64          `gorm:"column:category_id;not null" csv:"category_id"`
	UnitInPack    string          `gorm:"column:unit_in_pack;not null" csv:"unit_in_pack"`
	Weight        decimal.Decimal `gorm:"column:weight;not null;type:numeric"`
	Height        decimal.Decimal `gorm:"column:height;not null;type:numeric"`
	Length        decimal.Decimal `gorm:"column:length;not null;type:numeric"`
	Width         decimal.Decimal `gorm:"column:width;not null;type:numeric"`
	Image         string          `gorm:"column:image;not null" csv:"image"`
	CreatedAt     time.Time       `gorm:"column:created_at;not null"`
	UpdatedAt     time.Time       `gorm:"column:updated_at;not null"`
	DeletedAt     gorm.DeletedAt  `gorm:"index"`
	Manufacture   Manufacture     `gorm:"foreignKey:ManufactureID"`
	Form          Form            `gorm:"foreignKey:FormID"`
	Category      Category        `gorm:"foreignKey:CategoryID"`
}
