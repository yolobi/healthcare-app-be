package masterrepository

import (
	"context"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/repositories/masterrepositoryinterface"
	"healthcare-capt-america/utils"

	"gorm.io/gorm"
)

type specializationRepository struct {
	db *gorm.DB
}

// Find implements masterrepositoryinterface.SpecializationRepository.
func (repo *specializationRepository) Find(ctx context.Context) ([]models.Specialization, error) {
	res, err := utils.SelectQuery[models.Specialization](ctx, repo.db)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	return res, nil
}

// FindByID implements masterrepositoryinterface.SpecializationRepository.
func (repo *specializationRepository) FindByID(ctx context.Context, id uint64) (*models.Specialization, error) {
	return utils.GetById[models.Specialization](ctx, repo.db, id)
}

// Save implements masterrepositoryinterface.SpecializationRepository.
func (repo *specializationRepository) Save(ctx context.Context, specialization *models.Specialization) (*models.Specialization, error) {
	return utils.SaveQuery[models.Specialization](ctx, repo.db, specialization, enums.Create)
}

func NewSpecializationRepository(db *gorm.DB) *specializationRepository {
	return &specializationRepository{
		db: db,
	}
}

var _ masterrepositoryinterface.SpecializationRepository = &specializationRepository{}
