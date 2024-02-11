package productrepository

import (
	"context"
	"errors"
	"fmt"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/dto/responses/drugsres"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/entities/models/transaction"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/repositories/productrepositoryinterface"
	"healthcare-capt-america/services"
	"healthcare-capt-america/utils"
	"mime/multipart"

	"gorm.io/gorm"
)

type drugRepository struct {
	db *gorm.DB
}

// FindByIDUpdate implements productrepositoryinterface.DrugRepository.
func (repo *drugRepository) FindByIDUpdate(ctx context.Context, drug_id uint64) (*models.Drug, error) {
	var drug *models.Drug
	err := repo.db.WithContext(ctx).Model(&models.Drug{}).Preload("Manufacture").Preload("Form").
		Preload("Category").Where("id = ?", drug_id).Find(&drug).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, apperror.NewServerError(err)
	}
	return drug, nil
}

func NewDrugRepository(db *gorm.DB) *drugRepository {
	return &drugRepository{db: db}
}

func (repo *drugRepository) FindAll(ctx context.Context, qry *requests.GlobalQuery) ([]drugsres.DrugMaster, *responses.Pagination, error) {
	drugs := make([]drugsres.DrugMaster, 0)
	limit, offset := qry.GetPagination()
	preload := repo.db.WithContext(ctx).
		Model(&models.Drug{}).
		Select(`"drugs".ID`,
			`"drugs".name`,
			`"drugs".generic_name`,
			`"drugs".content`,
			`m.id as manufacture_id`,
			`c.id as category_id`,
			`f.id as form_id`,
			`f.name as form`,
			`"drugs".unit_in_pack`,
			`"drugs".weight`,
			`"drugs".height`,
			`"drugs".length`,
			`"drugs".width`,
			`"drugs".image`, `"drugs".unit_in_pack`, `c.name as category`,
			`m.name as manufacture`, `drugs.updated_at`, `max(pd.selling_unit) as max_selling_unit`, `min(pd.selling_unit) as min_selling_unit`, `sum(pd.stock) as total_stock`).
		Joins(`FULL OUTER JOIN forms f ON "drugs".form_id = f.id`).
		Joins(`FULL OUTER JOIN categories c ON "drugs".category_id = c.id`).
		Joins(`FULL OUTER JOIN manufactures m ON "drugs".manufacture_id = m.id`).
		Joins(`FULL OUTER JOIN pharmacy_drugs pd ON "drugs"."id" = pd.drug_id`).
		Joins(`FULL OUTER JOIN pharmacies p ON pd.pharmacy_id = p.id`).
		Joins(`FULL OUTER JOIN addresses a ON p.address_id = a.id`)
	for _, condition := range qry.Conditions {
		sql := fmt.Sprintf("%s %s ?", condition.Field, condition.Operation)
		preload.Where(sql, condition.Value)
	}
	searchQuery := fmt.Sprintf("drugs.name ILIKE %s OR drugs.content ILIKE %s OR drugs.generic_name ILIKE %s OR drugs.description ILIKE %s", qry.Search, qry.Search, qry.Search, qry.Search)
	preload.Where(searchQuery).Order(fmt.Sprintf(`"drugs".%s`, qry.OrderedBy))
	preload.Group(`"drugs".ID`).Group(`"drugs".name`).Group(`"drugs".generic_name`).Group(`"drugs".content`).
		Group(`m.id`).Group(`c.id`).Group(`f.id`).Group(`f.name`).Group(`"drugs".image`).Group(`c.name`).Group(`"drugs".unit_in_pack`).Group("m.name").Group(`"drugs".updated_at`)
	err := preload.Limit(limit).Offset(offset).Find(&drugs).Error
	if err != nil {
		return nil, nil, apperror.NewServerError(err, "Database Error")
	}
	var totalItems, totalPages int64
	err = repo.db.WithContext(ctx).Model(&models.Drug{}).Offset(-1).Count(&totalItems).Error
	if err != nil {
		return nil, nil, apperror.NewServerError(err, "Database Error")
	}
	totalPages = totalItems / int64(qry.PerPage)
	if totalItems%int64(qry.PerPage) != 0 {
		totalPages += 1
	}
	return drugs, &responses.Pagination{
		TotalItems:  totalItems,
		TotalPages:  totalPages,
		CurrentPage: int64(qry.Page),
	}, nil
}

func (repo *drugRepository) FindByID(ctx context.Context, drug_id uint64) (*drugsres.DrugMaster, error) {
	var drug *drugsres.DrugMaster
	err := repo.db.WithContext(ctx).Model(&models.Drug{}).
		Select(`"drugs".ID`,
			`"drugs".name`,
			`"drugs".generic_name`,
			`"drugs".content`,
			`"drugs".description`,
			`m.id as manufacture_id`,
			`c.id as category_id`,
			`f.id as form_id`,
			`f.name as form`,
			`"drugs".unit_in_pack`,
			`"drugs".weight`,
			`"drugs".height`,
			`"drugs".length`,
			`"drugs".width`,
			`"drugs".image`, `"drugs".unit_in_pack`, `c.name as category`,
			`m.name as manufacture`, `drugs.updated_at`, `max(pd.selling_unit) as max_selling_unit`, `min(pd.selling_unit) as min_selling_unit`, `sum(pd.stock) as total_stock`).
		Joins(`FULL OUTER JOIN categories c ON "drugs".category_id = c.id`).
		Joins(`FULL OUTER JOIN manufactures m ON "drugs".manufacture_id = m.id`).
		Joins(`FULL OUTER JOIN pharmacy_drugs pd ON "drugs"."id" = pd.drug_id`).
		Joins(`FULL OUTER JOIN pharmacies p ON pd.pharmacy_id = p.id`).
		Joins(`FULL OUTER JOIN addresses a ON p.address_id = a.id`).
		Joins(`FULL OUTER JOIN forms f ON f.id = "drugs".form_id`).
		Group(`"drugs".ID`).Group(`"drugs".name`).Group(`"drugs".generic_name`).Group(`"drugs".content`).
		Group(`m.id`).Group(`c.id`).Group(`f.id`).Group(`f.name`).Group(`"drugs".image`).Group(`c.name`).
		Group(`"drugs".unit_in_pack`).Group("m.name").Group(`"drugs".updated_at`).Where("drugs.id = ?", drug_id).Find(&drug).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, apperror.NewServerError(err)
	}
	return drug, nil
}

func (repo *drugRepository) Update(ctx context.Context, drug *models.Drug) (*models.Drug, error) {
	err := repo.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		fh := ctx.Value(enums.ProductImageKey).(*multipart.FileHeader)
		if fh != nil {
			s, err := services.UploadUserImage(ctx, fh, drug.Name+"_"+drug.GenericName)
			if err != nil {
				return apperror.NewServerError(err)
			}
			drug.Image = *s
		}
		_, err := utils.SaveQuery[models.Drug](ctx, repo.db, drug, enums.Update)
		if err != nil {
			return apperror.NewServerError(err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return drug, nil
}

func (repo *drugRepository) Save(ctx context.Context, drug *models.Drug) (*models.Drug, error) {
	err := repo.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		f := ctx.Value(enums.ProductImageKey.Key).(*multipart.FileHeader)
		s, err := services.UploadProductImage(ctx, f, drug.Name+"_"+drug.GenericName)
		if err != nil {
			return err
		}
		drug.Image = *s
		_, err = utils.SaveQuery[models.Drug](ctx, tx, drug, enums.Create)
		if err != nil {
			return apperror.NewServerError(err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return drug, nil
}

func (dr *drugRepository) FindAllWithDist(ctx context.Context, pos transaction.Position) ([]*drugsres.DrugUserRes, error) {
	drugs := make([]*drugsres.DrugUserRes, 0)
	s := fmt.Sprintf(enums.GetRadius25km, pos.Longitude.String(), pos.Latitude.String())
	err := dr.db.WithContext(ctx).
		Model(&models.Drug{}).
		Select(`"drugs".id, "drugs".name`, `"drugs".image`, `"drugs".unit_in_pack`, `c.name as category`, `m.name as manufacture`, `max(pd.selling_unit) as max_price`, `min(pd.selling_unit) as min_price`, `sum(pd.stock) as stock`).
		Joins(`JOIN categories c ON "drugs".category_id = c.id`).
		Joins(`JOIN manufactures m ON "drugs".manufacture_id = m.id`).
		Joins(`JOIN pharmacy_drugs pd ON "drugs"."id" = pd.drug_id`).
		Joins(`JOIN pharmacies p ON pd.pharmacy_id = p.id`).
		Joins(`JOIN addresses a ON p.address_id = a.id`).
		Where("pd.status = 'active'").
		Where(s).
		Group(`"drugs".id`).Group(`"drugs".name`).Group(`"drugs".image`).Group(`c.name`).Group(`"drugs".unit_in_pack`).Group("m.name").
		Scan(&drugs).Error
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	return drugs, nil
}

var _ productrepositoryinterface.DrugRepository = &drugRepository{}
