package telemedicineusecaseinterface

import (
	"context"
	"healthcare-capt-america/entities/models"
)

type TelemedicineUsecase interface {
	FindAll(context.Context, uint64) ([]*models.Room, error)
	FindAllByRoomId(ctx context.Context, room_id uint64, uId uint64) ([]*models.Chat, error)
	Update(ctx context.Context, room *models.Room, uId uint64) (*models.Room, error)
	CreateRoom(ctx context.Context, room *models.Room) (*models.Room, error)
}
