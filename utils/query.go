package utils

import (
	"context"
	"fmt"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/pkg/databases/pagination"

	"gorm.io/gorm"
)

func SelectQuery[T any](ctx context.Context, db *gorm.DB) ([]T, error) {
	var entities []T
	err := db.WithContext(ctx).Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func GetAllByKeyId[T any](ctx context.Context, db *gorm.DB, id uint64, key string) ([]T, error) {
	var entites []T
	keyId := fmt.Sprintf("%s = ?", key)
	err := db.WithContext(ctx).Model(&entites).Where(keyId, id).Find(&entites).Error
	if err != nil {
		return nil, err
	}
	return entites, nil
}

func SaveQuery[T any](ctx context.Context, db *gorm.DB, entity *T, action string) (*T, error) {
	query := db.WithContext(ctx).Model(entity)
	switch action {
	case enums.Create:
		query.Create(&entity)
	case enums.Update:
		query.Updates(&entity)
	}
	err := query.Error
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func GetById[T any](ctx context.Context, db *gorm.DB, id uint64) (entity *T, err error) {
	err = db.WithContext(ctx).Model(&entity).Where("id = ?", id).First(&entity).Error
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func CountTotalItems[T any](ctx context.Context, db *gorm.DB, entity *T) (uint, error) {
	var res int64 = 0
	err := db.WithContext(ctx).Model(&entity).Count(&res).Error
	out := uint(res)
	if err != nil {
		return out, err
	}
	return out, nil
}

func CountTotalItemsCondition[T any](ctx context.Context, db *gorm.DB, entity *T, field string, value string) (uint, error) {
	var res int64 = 0
	condition := fmt.Sprintf("%s = ?", field)
	err := db.WithContext(ctx).Model(&entity).Count(&res).Where(condition, value).Error
	out := uint(res)
	if err != nil {
		return out, err
	}
	return out, nil
}

func Delete[T any](ctx context.Context, db *gorm.DB, entity *T) error {
	return db.WithContext(ctx).Delete(entity).Error
}

func SelectQueryPagination[T any](ctx context.Context, db *gorm.DB, pagination *pagination.Pagination, preloads ...string) ([]T, error) {
	var res []T
	offset := pagination.Page * pagination.Limit
	err := db.WithContext(ctx).Order("id").Limit(pagination.Limit).Offset(offset).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func QueryPagination(repoQuery *gorm.DB, pagination *pagination.Pagination) *gorm.DB {
	offset := pagination.Page * pagination.Limit
	return repoQuery.Order("id").Limit(pagination.Limit).Offset(offset)
}

func SelectQueryPreload[T any](ctx context.Context, db *gorm.DB, preloads []string) ([]T, error) {
	var entities []T
	gorm := db.WithContext(ctx)
	for _, preload := range preloads {
		gorm = gorm.Preload(preload)
	}
	err := gorm.Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return entities, nil
}
