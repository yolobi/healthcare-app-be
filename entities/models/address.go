package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Address struct {
	ID         uint64          `gorm:"column:id;primaryKey;autoIncrement"`
	Detail     string          `gorm:"column:detail" csv:"detail"`
	ProvinceID uint64          `gorm:"column:province_id;not null" csv:"province_id"`
	CityID     uint64          `gorm:"column:city_id;not null" csv:"city_id"`
	UserId     *uint64         `gorm:"column:user_id;nullable"`
	Longtitude decimal.Decimal `gorm:"column:longtitude;not null;type:numeric"`
	Latitude   decimal.Decimal `gorm:"column:latitude;not null;type:numeric"`
	IsDefault  bool            `gorm:"column:is_default;nullable"`
	CreatedAt  time.Time       `gorm:"column:created_at;not null"`
	UpdatedAt  time.Time       `gorm:"column:updated_at;not null"`
	DeletedAt  gorm.DeletedAt  `gorm:"index"`
	Province   Province        `gorm:"foreignKey:ProvinceID"`
	City       City            `gorm:"foreignKey:CityID"`
	User       *User           `gorm:"foreignKey:UserId"`
}

func (address *Address) EditAddressRequest(editedAddressReq *Address) *Address {
	address.Detail = editedAddressReq.Detail
	address.Longtitude = editedAddressReq.Longtitude
	address.Latitude = editedAddressReq.Latitude
	address.CityID = editedAddressReq.CityID
	address.ProvinceID = editedAddressReq.ProvinceID
	return address
}
