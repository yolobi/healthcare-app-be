package models

import (
	"time"

	"gorm.io/gorm"
)

type Room struct {
	ID           uint64         `gorm:"column:id;primaryKey;autoIncrement"`
	DoctorID     uint64         `gorm:"column:doctor_id;not null"`
	UserID       uint64         `gorm:"column:user_id;not null"`
	RoomStatusID uint64         `gorm:"column:room_status_id;not null"`
	CreatedAt    time.Time      `gorm:"column:created_at;not null;default:current_timestamp"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;not null;default:current_timestamp"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Doctor       *Doctor        `gorm:"foreignKey:DoctorID;references:ID"`
	User         User           `gorm:"foreignKey:UserID;references:ID"`
	RoomStatus   RoomStatus     `gorm:"foreignKey:RoomStatusID;references:ID"`
}
