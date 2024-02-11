package masterrepository

import (
	"context"
	"errors"
	"fmt"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/repositories/masterrepositoryinterface"
	"healthcare-capt-america/utils"

	"gorm.io/gorm"
)

type provinceRepository struct {
	db *gorm.DB
}

func (repo *provinceRepository) FindAll(ctx context.Context, qry *requests.GlobalQuery) ([]models.Province, *responses.Pagination, error) {
	provinces := make([]models.Province, 0)
	limit, offset := qry.GetPagination()
	preloads := []string{"Cities"}
	preload := repo.db.WithContext(ctx).Model(&models.Province{})
	for _, load := range preloads {
		preload = preload.Preload(load)
	}
	for _, condition := range qry.Conditions {
		sql := fmt.Sprintf("%s %s ?", condition.Field, condition.Operation)
		preload.Where(sql, condition.Value)
	}
	searchQuery := fmt.Sprintf("name ILIKE %s", qry.Search)
	preload.Where(searchQuery).Order(qry.OrderedBy)
	err := preload.Limit(limit).Offset(offset).Find(&provinces).Error
	if err != nil {
		return nil, nil, apperror.NewServerError(err, "Database Error")
	}
	var totalItems, totalPages int64
	preload.Offset(-1).Count(&totalItems)
	totalPages = totalItems / int64(qry.PerPage)
	if totalItems%int64(qry.PerPage) != 0 {
		totalPages += 1
	}
	return provinces, &responses.Pagination{
		TotalItems:  totalItems,
		TotalPages:  totalPages,
		CurrentPage: int64(qry.Page),
	}, nil
}

func (repo *provinceRepository) FindById(ctx context.Context, id uint64) (*models.Province, error) {
	var province *models.Province
	err := repo.db.WithContext(ctx).
		Preload("Cities").Where("id = ?", id).Find(&province).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.NewClientError(err)
		}
		return nil, apperror.NewServerError(err)
	}
	return province, nil
}

func (repo *provinceRepository) Save(ctx context.Context, province *models.Province) (*models.Province, error) {
	return utils.SaveQuery[models.Province](ctx, repo.db, province, enums.Create)
}

func NewProvinceRepository(db *gorm.DB) *provinceRepository {
	return &provinceRepository{db: db}
}

var _ masterrepositoryinterface.ProvinceRepository = &provinceRepository{}
