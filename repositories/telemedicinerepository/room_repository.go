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
	"gorm.io/gorm/clause"
)

type roomRepository struct {
	db *gorm.DB
}

func (repo *roomRepository) FindByID2(ctx context.Context, id uint64) (*models.Room, error) {
	var room *models.Room
	err := repo.db.WithContext(ctx).Model(&models.Room{}).Preload("Doctor.User").Preload("User").
		Where("id", id).First(&room).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, apperror.NewServerError(err, "Database error")
	}
	return room, nil
}

func (repo *roomRepository) FindAll(ctx context.Context, uId *uint64, dId *uint64) ([]*models.Room, error) {
	rooms := make([]*models.Room, 0)
	preload := repo.db.WithContext(ctx).Model(&rooms).Preload("RoomStatus").Preload("User").Preload("Doctor.User")

	if uId != nil {
		preload.Where("user_id", uId)
	} else if dId != nil {
		preload.Where("doctor_id", dId)
	}
	err := preload.Preload(clause.Associations).Find(&rooms).Error
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (repo *roomRepository) Delete(ctx context.Context, room *models.Room) error {
	return utils.Delete[models.Room](ctx, repo.db, room)
}

func (repo *roomRepository) FindByID(ctx context.Context, id uint64) (*models.Room, error) {
	var room *models.Room
	err := repo.db.Model(&models.Room{}).
		Joins("Doctor").
		Where("rooms.id", id).
		First(&room).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, apperror.NewServerError(err)
	}
	return room, nil
}

func (repo *roomRepository) Save(ctx context.Context, room *models.Room) (*models.Room, error) {
	return utils.SaveQuery[models.Room](ctx, repo.db, room, enums.Create)
}

func (repo *roomRepository) Update(ctx context.Context, room *models.Room) (*models.Room, error) {
	err := repo.db.WithContext(ctx).
		Preload(clause.Associations).
		Updates(&room).First(&room).Error
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	return room, nil
}

func NewRoomRepository(db *gorm.DB) *roomRepository {
	return &roomRepository{db: db}
}

var _ telemedicinerepositoryinterface.RoomRepository = &roomRepository{}
