package masterrepository

import (
	"context"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/repositories/masterrepositoryinterface"
	"healthcare-capt-america/utils"

	"gorm.io/gorm"
)

type statusMutationRepository struct {
	db *gorm.DB
}

// Save implements repositories.StatusMutationRepository.
func (repo *statusMutationRepository) Save(ctx context.Context, statusMutation *models.StatusMutation) (*models.StatusMutation, error) {
	return utils.SaveQuery[models.StatusMutation](ctx, repo.db, statusMutation, enums.Create)
}

func (repo *statusMutationRepository) FindByID(ctx context.Context, statusMutation_id uint64) (*models.StatusMutation, error) {
	return utils.GetById[models.StatusMutation](ctx, repo.db, statusMutation_id)
}

func NewStatusMutationRepository(db *gorm.DB) *statusMutationRepository {
	return &statusMutationRepository{db: db}
}

var _ masterrepositoryinterface.StatusMutationRepository = &statusMutationRepository{}
