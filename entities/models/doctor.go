package models

import (
	"time"

	"gorm.io/gorm"
)

type Doctor struct {
	ID                uint64          `gorm:"column:id;primaryKey;autoIncrement"`
	UserId            uint64          `gorm:"column:user_id"`
	User              User            `gorm:"foreignKey:UserId"`
	Certificate       string          `gorm:"column:certificate"`
	YearsOfExperience int             `gorm:"column:years_of_experience"`
	SpecializationId  *uint64         `gorm:"column:specialization_id"`
	Specialization    *Specialization `gorm:"foreignKey:SpecializationId"`
	CreatedAt         time.Time       `gorm:"column:created_at;not null"`
	UpdatedAt         time.Time       `gorm:"column:updated_at;not null"`
	DeletedAt         gorm.DeletedAt  `gorm:"index"`
}
