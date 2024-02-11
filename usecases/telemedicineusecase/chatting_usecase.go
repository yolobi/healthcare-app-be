package telemedicineusecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/interfaces/repositories/doctorrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/telemedicinerepositoryinterface"
	"healthcare-capt-america/interfaces/usecases/telemedicineusecaseinterface"
	"sync"

	"gorm.io/gorm"

	"github.com/gorilla/websocket"
)

type ChattingUsecase struct {
	roomRepo   telemedicinerepositoryinterface.RoomRepository
	chatRepo   telemedicinerepositoryinterface.ChatRepository
	doctorRepo doctorrepositoryinterface.DoctorRepository
	manager    *ChatManager
	mutex      sync.Mutex
}

func (cu *ChattingUsecase) GetAllChatHistories(ctx context.Context, roomId uint64) ([]*models.Chat, error) {
	return cu.chatRepo.FindAllByRoomId(ctx, roomId)
}

type ChatManager struct {
	Rooms map[uint64]*ChatRoom
}

type ChatRoom struct {
	Clients []*Client
}

type Client struct {
	ID   uint64
	Name string
	Conn *websocket.Conn
}

type Message struct {
	IsTyping bool   `json:"is_typing"`
	SendTime string `json:"send_time"`
	SenderId uint64 `json:"sender_id"`
	Message  string `json:"message"`
}

func NewChattingUsecase(rr telemedicinerepositoryinterface.RoomRepository, cr telemedicinerepositoryinterface.ChatRepository, dr doctorrepositoryinterface.DoctorRepository) *ChattingUsecase {
	return &ChattingUsecase{
		roomRepo:   rr,
		chatRepo:   cr,
		doctorRepo: dr,
		manager: &ChatManager{
			Rooms: make(map[uint64]*ChatRoom),
		},
		mutex: sync.Mutex{},
	}
}

func (cu *ChattingUsecase) HandleWebSocket(ctx context.Context, uid uint64, room_id uint64, conn *websocket.Conn) error {
	r, err := cu.roomRepo.FindByID(ctx, room_id)
	if err != nil {
		return err
	}
	if r == nil {
		cu.sendError(ctx, conn, apperror.NewClientError(fmt.Errorf("room not found")))
		return apperror.NewClientError(fmt.Errorf("room not found"))
	}
	doctor, err := cu.doctorRepo.FindByUserId(ctx, uid)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return apperror.NewServerError(err, "Database error")
	}
	if r.UserID != uid {
		if doctor == nil || r.DoctorID != doctor.ID {
			cu.sendError(ctx, conn, apperror.NewClientError(fmt.Errorf("can't join this room chat")))
			return apperror.NewClientError(fmt.Errorf("can't join this room chat"))
		}
	}
	_, ok := cu.manager.Rooms[room_id]
	if !ok {
		cu.manager.Rooms[room_id] = &ChatRoom{
			Clients: make([]*Client, 0),
		}
	}
	room := cu.manager.Rooms[room_id]
	client := &Client{
		ID:   uid,
		Name: "contoh",
		Conn: conn,
	}
	room.Clients = append(room.Clients, client)
	go cu.readIncomingMessage(ctx, room_id, conn)

	for {
		var message Message
		if err := conn.ReadJSON(&message); err != nil {
			return apperror.NewServerError(err)
		}
		chat := &models.Chat{
			RoomID:   room_id,
			SenderID: client.ID,
			Message:  message.Message,
		}
		if !message.IsTyping {
			_, err := cu.chatRepo.Save(ctx, chat)
			if err != nil {
				cu.sendError(ctx, conn, err)
				return err
			}
		}
		message.SenderId = client.ID
		cu.broadcastMessage(ctx, cu.manager.Rooms[room_id], message, conn)
	}
}

func (cu *ChattingUsecase) readIncomingMessage(ctx context.Context, room_id uint64, conn *websocket.Conn) {
	for {
		_, rawMessage, err := conn.ReadMessage()
		if err != nil {
			room := cu.manager.Rooms[room_id]
			cu.removeConnection(ctx, room, conn)
			break
		}

		var message Message
		if err := json.Unmarshal(rawMessage, &message); err != nil {
			continue
		}

		cu.broadcastMessage(ctx, cu.manager.Rooms[room_id], message, conn)
	}
}

func (cu *ChattingUsecase) removeConnection(ctx context.Context, room *ChatRoom, conn *websocket.Conn) {
	for idx, client := range room.Clients {
		if client.Conn == conn {
			room.Clients[idx] = room.Clients[len(room.Clients)-1]
			room.Clients = room.Clients[:len(room.Clients)-1]
			break
		}
	}
}

func (cu *ChattingUsecase) broadcastMessage(ctx context.Context, room *ChatRoom, message Message, sender *websocket.Conn) error {
	cu.mutex.Lock()
	defer cu.mutex.Unlock()
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		return apperror.NewServerError(err)
	}

	for _, client := range room.Clients {
		if client.Conn == sender {
			continue
		}
		err := client.Conn.WriteMessage(websocket.TextMessage, jsonMessage)
		if err != nil {
			return apperror.NewServerError(err)
		}
	}
	return nil
}

func (cu *ChattingUsecase) sendError(ctx context.Context, conn *websocket.Conn, err error) {
	message, _ := json.Marshal(Message{
		SenderId: 0,
		Message:  err.Error(),
	})
	conn.WriteMessage(websocket.TextMessage, message)
}

var _ telemedicineusecaseinterface.ChattingUsecase = &ChattingUsecase{}
