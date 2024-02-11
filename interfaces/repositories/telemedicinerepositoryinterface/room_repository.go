package telemedicinerepositoryinterface

import (
	"context"
	"healthcare-capt-america/entities/models"
)

type RoomRepository interface {
	Save(ctx context.Context, room *models.Room) (*models.Room, error)
	FindByID(ctx context.Context, id uint64) (*models.Room, error)
	FindByID2(ctx context.Context, id uint64) (*models.Room, error)
	FindAll(ctx context.Context, uId *uint64, dId *uint64) ([]*models.Room, error)
	Update(ctx context.Context, room *models.Room) (*models.Room, error)
	Delete(ctx context.Context, room *models.Room) error
}
