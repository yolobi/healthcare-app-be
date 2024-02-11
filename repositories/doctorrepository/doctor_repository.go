package doctorrepository

import (
	"context"
	"errors"
	"fmt"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/repositories/doctorrepositoryinterface"
	"healthcare-capt-america/utils"

	"gorm.io/gorm"
)

type doctorRepository struct {
	db *gorm.DB
}

func (dr *doctorRepository) UpdateProfile(ctx context.Context, user *models.User, doctor *models.Doctor) (*models.Doctor, error) {
	var resp *models.Doctor
	err := dr.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		_, err := utils.SaveQuery[models.User](ctx, tx, user, enums.Update)
		if err != nil {
			return apperror.NewServerError(err)
		}
		// _, err = utils.SaveQuery[models.Doctor](ctx, tx, doctor, enums.Update)
		err = tx.Updates(models.Doctor{
			ID:     doctor.ID,
			UserId: doctor.UserId,
			// User:              doctor.User,
			Certificate:       doctor.Certificate,
			YearsOfExperience: doctor.YearsOfExperience,
			SpecializationId:  doctor.SpecializationId,
			// Specialization:    doctor.Specialization,
		}).Error
		if err != nil {
			return apperror.NewServerError(err)
		}
		err = tx.Model(&models.Doctor{}).
			Where("id", doctor.ID).
			Preload("User").
			Preload("Specialization").
			First(&resp).Error
		if err != nil {
			return apperror.NewServerError(err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (dr *doctorRepository) FindAll(ctx context.Context, qry *requests.GlobalQuery) ([]*models.Doctor, *responses.Pagination, error) {
	doctors := make([]*models.Doctor, 0)
	limit, offset := qry.GetPagination()
	preload := dr.db.WithContext(ctx).
		Model(&models.Doctor{}).
		Joins("User").
		Joins("Specialization").
		Where(`"User".password IS NOT NULL`)
	for _, condition := range qry.Conditions {
		sql := fmt.Sprintf("%s %s ?", condition.Field, condition.Operation)
		preload.Where(sql, condition.Value)
	}
	searchQuery := fmt.Sprintf(`"User".name ILIKE %s OR "Specialization".name ILIKE %s`, qry.Search, qry.Search)
	preload.Where(searchQuery).Order(fmt.Sprintf(`"User".%s`, qry.OrderedBy))
	err := preload.Limit(limit).Offset(offset).Find(&doctors).Error
	if err != nil {
		return nil, nil, apperror.NewServerError(err)
	}
	var totalItems, totalPages int64
	err = dr.db.WithContext(ctx).Model(&models.Doctor{}).Offset(-1).Count(&totalItems).Error
	if err != nil {
		return nil, nil, apperror.NewServerError(err, "Database error")
	}
	totalPages = totalItems / int64(qry.PerPage)
	if totalItems%int64(qry.PerPage) != 0 {
		totalPages += 1
	}
	return doctors, &responses.Pagination{
		TotalItems:  totalItems,
		TotalPages:  totalPages,
		CurrentPage: int64(qry.Page),
	}, nil
}

func (dr *doctorRepository) FindById(ctx context.Context, id uint64) (*models.Doctor, error) {
	var doctor *models.Doctor
	err := dr.db.WithContext(ctx).Model(&models.Doctor{}).
		Where("id", id).
		Preload("User").
		Preload("Specialization").
		First(&doctor).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, apperror.NewServerError(err)
	}
	return doctor, nil
}

func (dr *doctorRepository) FindByUserId(ctx context.Context, user_id uint64) (*models.Doctor, error) {
	var doctor *models.Doctor
	err := dr.db.WithContext(ctx).
		Where("user_id", user_id).
		Preload("User").
		Preload("Specialization").
		First(&doctor).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, apperror.NewServerError(err)
	}
	return doctor, nil
}

func (dr *doctorRepository) Update(ctx context.Context, doctor *models.Doctor) (*models.Doctor, error) {
	err := dr.db.WithContext(ctx).Save(&doctor).Error
	if err != nil {
		return nil, err
	}
	return doctor, nil
}

func NewDoctorRepository(db *gorm.DB) *doctorRepository {
	return &doctorRepository{
		db: db,
	}
}

var _ doctorrepositoryinterface.DoctorRepository = &doctorRepository{}
