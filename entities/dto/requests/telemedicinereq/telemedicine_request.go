package telemedicinereq

import "healthcare-capt-america/entities/models"

type UpdateRoomStatus struct {
	Status string `json:"room_status,omitempty"`
}

func (urs *UpdateRoomStatus) ToRoom() (room models.Room) {
	switch urs.Status {
	case "ongoing":
		room.RoomStatusID = 1
	case "end":
		room.RoomStatusID = 2
	}
	return
}

type CreateRoom struct {
	DoctorID uint64 `json:"doctor_id"`
}
