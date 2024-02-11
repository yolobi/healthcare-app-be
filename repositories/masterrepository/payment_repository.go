package masterrepository

import (
	"context"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/repositories/masterrepositoryinterface"
	"healthcare-capt-america/utils"

	"gorm.io/gorm"
)

type paymentRepository struct {
	db *gorm.DB
}

// Update implements masterrepositoryinterface.PaymentRepository.
func (repo *paymentRepository) Update(ctx context.Context, payment *models.Payment) (*models.Payment, error) {
	return utils.SaveQuery[models.Payment](ctx, repo.db, payment, enums.Update)
}

func (repo *paymentRepository) FindByID(ctx context.Context, id uint64) (*models.Payment, error) {
	return utils.GetById[models.Payment](ctx, repo.db, id)
}

func (repo *paymentRepository) ValidateUserPayment(ctx context.Context, paymentId uint64, userId uint64) (bool, error) {
	var isExist bool
	err := repo.db.WithContext(ctx).Raw(enums.CheckPaymentFromUser, userId, paymentId).Scan(&isExist).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

// Save implements repositories.PaymentRepository.
func (repo *paymentRepository) Save(ctx context.Context, payment *models.Payment) (*models.Payment, error) {
	return utils.SaveQuery[models.Payment](ctx, repo.db, payment, enums.Create)
}

func NewPaymentRepository(db *gorm.DB) *paymentRepository {
	return &paymentRepository{db: db}
}

var _ masterrepositoryinterface.PaymentRepository = &paymentRepository{}
