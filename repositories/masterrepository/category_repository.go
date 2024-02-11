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

type categoryRepository struct {
	db *gorm.DB
}

func (repo *categoryRepository) Find(ctx context.Context) ([]models.Category, error) {
	return utils.SelectQuery[models.Category](ctx, repo.db)
}

func (repo *categoryRepository) FindAll(ctx context.Context, qry *requests.GlobalQuery) ([]models.Category, *responses.Pagination, error) {
	categories := make([]models.Category, 0)
	limit, offset := qry.GetPagination()
	preload := repo.db.WithContext(ctx).Model(&models.Category{})
	for _, condition := range qry.Conditions {
		sql := fmt.Sprintf("%s %s ?", condition.Field, condition.Operation)
		preload.Where(sql, condition.Value)
	}
	searchQuery := fmt.Sprintf("name ILIKE %s ", qry.Search)
	preload.Where(searchQuery).Order(qry.OrderedBy)
	err := preload.Limit(limit).Offset(offset).Find(&categories).Error
	if err != nil {
		return nil, nil, apperror.NewServerError(err, "Database Error")
	}
	var totalItems, totalPages int64
	preload.Offset(-1).Count(&totalItems)
	totalPages = totalItems / int64(qry.PerPage)
	if totalItems%int64(qry.PerPage) != 0 {
		totalPages += 1
	}
	return categories, &responses.Pagination{
		TotalItems:  totalItems,
		TotalPages:  totalPages,
		CurrentPage: int64(qry.Page),
	}, nil
}

func (repo *categoryRepository) FindByID(ctx context.Context, category_id uint64) (category *models.Category, err error) {
	category, err = utils.GetById[models.Category](ctx, repo.db, category_id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, apperror.NewServerError(err)
	}
	return
}

func (repo *categoryRepository) Update(ctx context.Context, category *models.Category) (*models.Category, error) {
	return utils.SaveQuery[models.Category](ctx, repo.db, category, enums.Update)
}

func (repo *categoryRepository) Save(ctx context.Context, category *models.Category) (*models.Category, error) {
	return utils.SaveQuery[models.Category](ctx, repo.db, category, enums.Create)
}

func (repo *categoryRepository) Delete(ctx context.Context, category *models.Category) error {
	return utils.Delete[models.Category](ctx, repo.db, category)
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db: db}
}

var _ masterrepositoryinterface.CategoryRepository = &categoryRepository{}
