package cartrepository

import (
	"context"
	"healthcare-capt-america/entities/dto/responses/checkoutres"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/repositories/cartrepositoryinterface"
	"healthcare-capt-america/utils"

	"gorm.io/gorm"
)

type cartRepository struct {
	db *gorm.DB
}

func (repo *cartRepository) FindCartListUserForCheckout(ctx context.Context, userId uint64, pharmacyId uint64) ([]checkoutres.CartResponse, error) {
	var responses []checkoutres.CartResponse
	err := repo.db.WithContext(ctx).Raw(enums.GetCartForCheckout, pharmacyId, userId).Scan(&responses).Error
	if err != nil {
		return nil, err
	}
	return responses, nil
}

func (repo *cartRepository) DeleteCartByUserID(ctx context.Context, userId uint64) error {
	return repo.db.WithContext(ctx).Unscoped().Model(&models.Cart{}).Where("user_id = ?", userId).Delete(userId).Error
}

func (repo *cartRepository) FindByID(ctx context.Context, cartId uint64) (*models.Cart, error) {
	return utils.GetById[models.Cart](ctx, repo.db, cartId)
}

func (repo *cartRepository) Delete(ctx context.Context, cart *models.Cart) error {
	return repo.db.WithContext(ctx).Unscoped().Model(&models.Cart{}).Where("id = ?", cart.ID).Delete(&cart).Error
}

func (repo *cartRepository) UpdateDrugCartQuantity(ctx context.Context, cart *models.Cart) (*models.Cart, error) {
	err := repo.db.WithContext(ctx).Model(&models.Cart{}).
		Where("user_id = ? AND drug_id = ?", cart.UserId, cart.DrugId).
		Update("quantity", gorm.Expr("quantity + ?", cart.Quantity)).
		Scan(&cart).Error
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func (repo *cartRepository) CheckDrugInCartExists(ctx context.Context, drugId uint64, userId uint64) (bool, error) {
	var exists bool
	err := repo.db.WithContext(ctx).Raw(enums.CheckProductExistInCartUser, drugId, userId).Scan(&exists).Error
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (repo *cartRepository) FindCartListByUserId(ctx context.Context, userId uint64) ([]*models.Cart, error) {
	carts := make([]*models.Cart, 0)
	err := repo.db.WithContext(ctx).Model(&models.Cart{}).Preload("User").Preload("Drug.Form").Preload("Drug.Manufacture").
		Preload("Drug.Category").Where("user_id", userId).Find(&carts).Error
	if err != nil {
		return nil, err
	}
	return carts, nil
}

func (repo *cartRepository) Save(ctx context.Context, cart *models.Cart) (*models.Cart, error) {
	return utils.SaveQuery[models.Cart](ctx, repo.db, cart, enums.Create)
}

func NewCartRepository(db *gorm.DB) *cartRepository {
	return &cartRepository{db: db}
}

var _ cartrepositoryinterface.CartRepository = &cartRepository{}
