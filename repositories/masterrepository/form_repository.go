package masterrepository

import (
	"context"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/repositories/masterrepositoryinterface"
	"healthcare-capt-america/utils"

	"gorm.io/gorm"
)

type formRepository struct {
	db *gorm.DB
}

func (repo *formRepository) FindAll(ctx context.Context) ([]models.Form, error) {
	return utils.SelectQuery[models.Form](ctx, repo.db)
}

// Save implements repositories.FormRepository.
func (repo *formRepository) Save(ctx context.Context, form *models.Form) (*models.Form, error) {
	return utils.SaveQuery[models.Form](ctx, repo.db, form, enums.Create)
}

func (repo *formRepository) FindByID(ctx context.Context, form_id uint64) (*models.Form, error) {
	return utils.GetById[models.Form](ctx, repo.db, form_id)
}

func NewFormRepository(db *gorm.DB) *formRepository {
	return &formRepository{db: db}
}

var _ masterrepositoryinterface.FormRepository = &formRepository{}
