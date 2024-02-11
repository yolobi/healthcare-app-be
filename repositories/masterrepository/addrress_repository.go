package masterrepository

import (
	"context"
	"errors"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/repositories/masterrepositoryinterface"
	"healthcare-capt-america/utils"

	"gorm.io/gorm"
)

type AddrressRepository struct {
	db *gorm.DB
}

func NewADdressRepository(db *gorm.DB) *AddrressRepository {
	return &AddrressRepository{
		db: db,
	}
}

func (ar *AddrressRepository) FindByUserId(ctx context.Context, uid uint64) (*models.Address, error) {
	var address *models.Address

	err := ar.db.WithContext(ctx).
		Where("user_id", uid).
		Where("is_default", true).
		First(&address).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, apperror.NewServerError(err)
	}
	return address, nil
}

func (ar *AddrressRepository) Save(ctx context.Context, address *models.Address) (*models.Address, error) {
	address, err := utils.SaveQuery[models.Address](ctx, ar.db, address, enums.Create)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	return address, nil
}

func (ar *AddrressRepository) FindById(ctx context.Context, id uint64) (*models.Address, error) {
	var address *models.Address
	err := ar.db.WithContext(ctx).
		Where("id", id).
		First(&address).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, apperror.NewServerError(err)
	}
	return address, nil
}

func (ar *AddrressRepository) FindAllByUserId(ctx context.Context, uid uint64) ([]*models.Address, error) {
	addresses := make([]*models.Address, 0)
	err := ar.db.WithContext(ctx).Model(&models.Address{}).
		Preload("Province").
		Preload("City").
		Where("user_id", uid).
		Find(&addresses).Error
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	return addresses, nil
}

func (ar *AddrressRepository) SetDefault(ctx context.Context, address_id uint64) error {
	uid := ctx.Value(enums.UserIdKey).(uint64)
	err := ar.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var oldDefault *models.Address
		err := tx.Model(&models.Address{}).
			Where("user_id", uid).
			Where("is_default", true).
			First(&oldDefault).Error
		if err != nil {
			return apperror.NewServerError(err)
		}
		oldDefault.IsDefault = false
		err = tx.Save(&oldDefault).Error
		if err != nil {
			return apperror.NewServerError(err)
		}

		var newDefault *models.Address
		err = tx.Model(&models.Address{}).
			Where("id", address_id).
			First(&newDefault).Error
		if err != nil {
			return apperror.NewServerError(err)
		}
		newDefault.IsDefault = true
		err = tx.Save(&newDefault).Error
		if err != nil {
			return apperror.NewServerError(err)
		}
		return nil
	})
	return err
}

func (ar *AddrressRepository) Delete(ctx context.Context, address_id uint64) error {
	err := ar.db.WithContext(ctx).Delete(&models.Address{}, address_id).Error
	if err != nil {
		return apperror.NewServerError(err)
	}
	return nil
}

func (ar *AddrressRepository) Update(ctx context.Context, address *models.Address) (*models.Address, error) {
	a, err := utils.SaveQuery[models.Address](ctx, ar.db, address, enums.Update)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	return a, nil
}

var _ masterrepositoryinterface.AddrressRepository = &AddrressRepository{}
