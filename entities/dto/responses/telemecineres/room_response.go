package telemecineres

import (
	"healthcare-capt-america/entities/models"
	"time"
)

type RoomResponse struct {
	ID           uint64    `json:"id"`
	DoctorID     uint64    `json:"doctor_id"`
	DoctorUserID uint64    `json:"doctor_user_id"`
	DoctorName   string    `json:"doctor_name"`
	DoctorPhoto  string    `json:"doctor_photo"`
	UserID       uint64    `json:"user_id"`
	UserName     string    `json:"user_name"`
	UserPhoto    string    `json:"user_photo"`
	RoomStatusID uint64    `json:"room_status_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (jr *RoomResponse) Set(room *models.Room) {
	jr.ID = room.ID
	jr.DoctorID = room.Doctor.ID
	jr.DoctorName = room.Doctor.User.Name
	jr.DoctorPhoto = room.Doctor.User.Photo
	jr.DoctorUserID = room.Doctor.UserId
	jr.UserID = room.UserID
	jr.UserName = room.User.Name
	jr.UserPhoto = room.User.Photo
	jr.RoomStatusID = room.RoomStatusID
	jr.CreatedAt = room.CreatedAt
	jr.UpdatedAt = room.UpdatedAt
}
