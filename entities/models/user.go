package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint64         `gorm:"column:id;primaryKey;autoIncrement"`
	Name         string         `gorm:"column:name;not null"`
	Email        string         `gorm:"column:email;not null;unique"`
	Password     *string        `gorm:"column:password"`
	IsVerify     bool           `gorm:"column:is_verify;not null"`
	PhoneNumber  *string        `gorm:"column:phone_number;unique"`
	OnlineStatus string         `gorm:"column:online_status;not null"`
	Photo        string         `gorm:"column:photo;not null"`
	Addresses    []*Address     `gorm:"foreignKey:UserId"`
	CreatedAt    time.Time      `gorm:"column:created_at;not null"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;not null"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (u *User) UpdateUserData(newUser *User) *User {
	if newUser.Name != "" {
		u.Name = newUser.Name
	}
	if newUser.PhoneNumber != nil {
		u.PhoneNumber = newUser.PhoneNumber
	}
	if newUser.Photo != "" {
		u.Photo = newUser.Photo
	}
	return u
}
