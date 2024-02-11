package models

import (
	"time"

	"gorm.io/gorm"
)

type Chat struct {
	ID        uint64         `gorm:"column:id;primaryKey;autoIncrement"`
	RoomID    uint64         `gorm:"column:room_id;not null"`
	SenderID  uint64         `gorm:"column:sender_id;not null"`
	Message   string         `gorm:"column:message;not null"`
	CreatedAt time.Time      `gorm:"column:created_at;not null;default:current_timestamp"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:current_timestamp"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Room      Room           `gorm:"foreignKey:RoomID;references:ID"`
	Sender    User           `gorm:"foreignKey:SenderID;references:ID"`
}
