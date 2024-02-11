package telemedicinerepositoryinterface

import (
	"context"
	"healthcare-capt-america/entities/models"
)

type ChatRepository interface {
	Save(ctx context.Context, chat *models.Chat) (*models.Chat, error)
	Update(ctx context.Context, chat *models.Chat) (*models.Chat, error)
	FindByMessage(ctx context.Context, message string) ([]*models.Chat, error)
	FindAllByRoomId(ctx context.Context, room_id uint64) ([]*models.Chat, error)
	Delete(ctx context.Context, chat *models.Chat) error
}
