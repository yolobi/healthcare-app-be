package telemedicinerepository

import (
	"context"
	"errors"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/repositories/telemedicinerepositoryinterface"
	"healthcare-capt-america/utils"

	"gorm.io/gorm"
)

type chatRepository struct {
	db *gorm.DB
}

func (repo *chatRepository) Delete(ctx context.Context, chat *models.Chat) error {
	return utils.Delete[models.Chat](ctx, repo.db, chat)
}

func (repo *chatRepository) FindAllByRoomId(ctx context.Context, room_id uint64) ([]*models.Chat, error) {
	var chats = make([]*models.Chat, 0)
	err := repo.db.Model(&chats).Preload("Sender").Where("room_id = ?", room_id).Find(&chats).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, apperror.NewServerError(err)
	}
	return chats, nil
}

func (repo *chatRepository) FindByMessage(ctx context.Context, message string) (chat []*models.Chat, err error) {
	err = repo.db.WithContext(ctx).Model(&chat).Where("chat = ?", message).Find(&chat).Error
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	return chat, nil
}

func (repo *chatRepository) Save(ctx context.Context, chat *models.Chat) (*models.Chat, error) {
	return utils.SaveQuery[models.Chat](ctx, repo.db, chat, enums.Create)
}

func (repo *chatRepository) Update(ctx context.Context, chat *models.Chat) (*models.Chat, error) {
	return utils.SaveQuery[models.Chat](ctx, repo.db, chat, enums.Update)
}

func NewChatRepository(db *gorm.DB) *chatRepository {
	return &chatRepository{db: db}
}

var _ telemedicinerepositoryinterface.ChatRepository = &chatRepository{}
