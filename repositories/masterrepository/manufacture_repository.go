package masterrepository

import (
	"context"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/repositories/masterrepositoryinterface"
	"healthcare-capt-america/utils"

	"gorm.io/gorm"
)

type manufactureRepository struct {
	db *gorm.DB
}

func (repo *manufactureRepository) FindAll(ctx context.Context) ([]models.Manufacture, error) {
	return utils.SelectQuery[models.Manufacture](ctx, repo.db)
}

// Save implements repositories.ManufactureRepository.
func (repo *manufactureRepository) Save(ctx context.Context, manufacture *models.Manufacture) (*models.Manufacture, error) {
	return utils.SaveQuery[models.Manufacture](ctx, repo.db, manufacture, enums.Create)
}

func (repo *manufactureRepository) FindByID(ctx context.Context, manufacture_id uint64) (*models.Manufacture, error) {
	return utils.GetById[models.Manufacture](ctx, repo.db, manufacture_id)
}

func NewManufactureRepository(db *gorm.DB) *manufactureRepository {
	return &manufactureRepository{db: db}
}

var _ masterrepositoryinterface.ManufactureRepository = &manufactureRepository{}
