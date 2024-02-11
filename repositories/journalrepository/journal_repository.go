package journalrepository

import (
	"context"
	"fmt"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/interfaces/repositories/journalrepositoryinterface"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type journalRepository struct {
	db *gorm.DB
}

func NewJournalRepository(db *gorm.DB) *journalRepository {
	return &journalRepository{db: db}
}

func (repo *journalRepository) FindAll(ctx context.Context, qry *requests.GlobalQuery, uId uint64) ([]models.Journal, *responses.Pagination, error) {
	journals := make([]models.Journal, 0)
	limit, offset := qry.GetPagination()
	preload := repo.db.WithContext(ctx).Model(&journals).
		Joins("JOIN pharmacies ON pharmacies.id=journals.from_pharmacy_id OR pharmacies.id=journals.to_pharmacy_id AND pharmacies.admin_pharmacy_id = ?", uId)
	for _, condition := range qry.Conditions {
		sql := fmt.Sprintf("%s %s ?", condition.Field, condition.Operation)
		preload.Where(sql, condition.Value)
	}
	preload.Offset(offset).Limit(limit).Order(qry.OrderedBy)
	err := preload.Preload(clause.Associations).Find(&journals).Error
	if err != nil {
		return nil, nil, err
	}
	var totalItems, totalPages int64
	preload.Offset(-1).Count(&totalItems)
	totalPages = totalItems / int64(qry.PerPage)
	if totalItems%int64(qry.PerPage) != 0 {
		totalPages += 1
	}
	return journals, &responses.Pagination{
		TotalItems:  totalItems,
		TotalPages:  totalPages,
		CurrentPage: int64(qry.Page),
	}, nil
}

var _ journalrepositoryinterface.JournalRepository = &journalRepository{}
