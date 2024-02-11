package productcategoryusecase

import (
	"context"
	"fmt"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/repositories/masterrepositoryinterface"
	"healthcare-capt-america/interfaces/usecases/productcategoryusecaseinterface"
	"healthcare-capt-america/services"
	"healthcare-capt-america/utils"
	"log"
	"mime/multipart"
)

type categoryUsecase struct {
	categoryRepo masterrepositoryinterface.CategoryRepository
}

func (usecase *categoryUsecase) FindAllCategories(ctx context.Context) ([]models.Category, error) {
	categories, err := usecase.categoryRepo.Find(ctx)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	return categories, nil
}

func NewCategoryUsecase(categoryRepo masterrepositoryinterface.CategoryRepository) *categoryUsecase {
	return &categoryUsecase{categoryRepo: categoryRepo}
}

func (usecase *categoryUsecase) EditCategory(ctx context.Context, category *models.Category) (*models.Category, error) {
	_, err := usecase.categoryRepo.FindByID(ctx, category.ID)
	if err != nil {
		return nil, err
	}
	if ctx.Value(enums.CategoryIconKey.Key) != nil {
		a := ctx.Value(enums.CategoryIconKey.Key).(*multipart.FileHeader)
		fileExt, _ := utils.GetFileType(a)
		fileName := fmt.Sprintf("%d%s", category.ID, fileExt)
		log.Println(fileName)
		path, err := services.UploadIcon(ctx, a, fileName)
		if err != nil {
			return nil, apperror.NewServerError(err)
		}
		category.Icon = *path
	}
	return usecase.categoryRepo.Update(ctx, category)
}

func (usecase *categoryUsecase) FindByID(ctx context.Context, category_id uint64) (*models.Category, error) {
	category, err := usecase.categoryRepo.FindByID(ctx, category_id)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, apperror.NewClientError(fmt.Errorf("can't find category with id %d", category_id))
	}
	return category, nil
}

func (usecase *categoryUsecase) FindAllCategory(ctx context.Context, qry *requests.GlobalQuery) ([]models.Category, *responses.Pagination, error) {
	return usecase.categoryRepo.FindAll(ctx, qry)
}

func (usecase *categoryUsecase) CreateCategory(ctx context.Context, category *models.Category) (*models.Category, error) {
	category, err := usecase.categoryRepo.Save(ctx, category)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	a := ctx.Value(enums.CategoryIconKey.Key).(*multipart.FileHeader)
	fileExt, _ := utils.GetFileType(a)
	fileName := fmt.Sprintf("%d%s", category.ID, fileExt)
	path, err := services.UploadIcon(ctx, a, fileName)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	category.Icon = *path
	return usecase.categoryRepo.Update(ctx, category)
}

func (usecase *categoryUsecase) DeleteCategory(ctx context.Context, category_id uint64) error {
	pharmacyDrug, err := usecase.categoryRepo.FindByID(ctx, category_id)
	if err != nil {
		return err
	}
	return usecase.categoryRepo.Delete(ctx, pharmacyDrug)
}

var _ productcategoryusecaseinterface.CategoryUsecase = &categoryUsecase{}
