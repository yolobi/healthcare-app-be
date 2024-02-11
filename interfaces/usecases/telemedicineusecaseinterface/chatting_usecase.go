package telemedicineusecaseinterface

import (
	"context"
	"healthcare-capt-america/entities/models"

	"github.com/gorilla/websocket"
)

type ChattingUsecase interface {
	HandleWebSocket(ctx context.Context, uid uint64, room_id uint64, conn *websocket.Conn) error
	GetAllChatHistories(ctx context.Context, roomId uint64) ([]*models.Chat, error)
}
