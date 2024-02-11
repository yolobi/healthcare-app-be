package models

import (
	"time"

	"gorm.io/gorm"
)

type ForgotPasswordToken struct {
	ID        uint64         `gorm:"column:id;primaryKey;autoIncrement"`
	UserId    uint64         `gorm:"column:user_id;not null"`
	User      User           `gorm:"foreignKey:UserId"`
	Token     string         `gorm:"column:token;unique;not null"`
	IsValid   bool           `gorm:"column:is_valid;default:true"`
	CreatedAt time.Time      `gorm:"column:created_at;not null"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
