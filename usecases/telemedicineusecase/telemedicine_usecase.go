package telemedicineusecase

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/interfaces/repositories/doctorrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/telemedicinerepositoryinterface"
	"healthcare-capt-america/interfaces/usecases/telemedicineusecaseinterface"
)

type telemedicineUsecase struct {
	roomRepo   telemedicinerepositoryinterface.RoomRepository
	chatRepo   telemedicinerepositoryinterface.ChatRepository
	doctorRepo doctorrepositoryinterface.DoctorRepository
}

func (usecase *telemedicineUsecase) CreateRoom(ctx context.Context, room *models.Room) (*models.Room, error) {
	room, err := usecase.roomRepo.Save(ctx, room)
	if err != nil {
		return nil, apperror.NewServerError(err, "Database error")
	}
	room, err = usecase.roomRepo.FindByID2(ctx, room.ID)
	if err != nil {
		return nil, apperror.NewServerError(err, "Database error")
	}
	return room, nil
}

func NewTelemedicineUsecase(roomRepo telemedicinerepositoryinterface.RoomRepository,
	chatRepo telemedicinerepositoryinterface.ChatRepository, doctoRepo doctorrepositoryinterface.DoctorRepository) *telemedicineUsecase {
	return &telemedicineUsecase{roomRepo: roomRepo, chatRepo: chatRepo, doctorRepo: doctoRepo}
}

func (usecase *telemedicineUsecase) FindAll(ctx context.Context, uId uint64) ([]*models.Room, error) {
	doctor, err := usecase.doctorRepo.FindByUserId(ctx, uId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, apperror.NewServerError(err)
	}
	if doctor != nil {
		return usecase.roomRepo.FindAll(ctx, nil, &doctor.ID)
	}
	return usecase.roomRepo.FindAll(ctx, &uId, nil)
}

func (usecase *telemedicineUsecase) FindAllByRoomId(ctx context.Context, room_id uint64, uId uint64) ([]*models.Chat, error) {
	room, err := usecase.roomRepo.FindByID(ctx, room_id)
	if err != nil {
		return nil, err
	}
	if room == nil {
		return nil, apperror.NewClientError(errors.New("room not found"))
	}
	if room.UserID != uId { //doctor.User.Id
		return nil, apperror.NewClientError(errors.New("you can't access this room"))
	}
	return usecase.chatRepo.FindAllByRoomId(ctx, room_id)
}

func (usecase *telemedicineUsecase) Update(ctx context.Context, room *models.Room, uId uint64) (*models.Room, error) {
	data, err := usecase.roomRepo.FindByID(ctx, room.ID)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, apperror.NewClientError(fmt.Errorf("room with id %d is not found", room.ID))
	}
	if room.UserID != uId { //doctor.User.Id
		return nil, apperror.NewClientError(errors.New("you not have access to update this room"))
	}
	data.RoomStatusID = room.RoomStatusID
	return usecase.roomRepo.Update(ctx, room)
}

var _ telemedicineusecaseinterface.TelemedicineUsecase = &telemedicineUsecase{}
