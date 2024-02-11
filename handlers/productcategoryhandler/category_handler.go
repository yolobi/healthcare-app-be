package productcategoryhandler

import (
	"context"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/requests/categoryreq"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/dto/responses/categoryres"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/handlers/productcategoryhandlerinterface"
	"healthcare-capt-america/interfaces/usecases/productcategoryusecaseinterface"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	usecase productcategoryusecaseinterface.CategoryUsecase
}

func (handler *CategoryHandler) FindAllCategories(c *gin.Context) {
	categories, err := handler.usecase.FindAllCategories(c.Request.Context())
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, responses.DefaultResponse{Data: gin.H{"categories": categoryres.NewCategories(categories)}})
}

func NewCategoryHandler(usecase productcategoryusecaseinterface.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{usecase: usecase}
}

func (handler *CategoryHandler) CreateCategory(ctx *gin.Context) {
	var request categoryreq.CategoryRequest
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.Error(apperror.NewValidationError(err))
		return
	}
	fh, err := ctx.FormFile(enums.FileIcon)
	if err != nil {
		ctx.Error(apperror.NewValidationError(err))
		return
	}
	ctx2 := ctx.Request.Context()
	ctx3 := context.WithValue(ctx2, enums.CategoryIconKey.Key, fh)
	ctx.Request = ctx.Request.WithContext(ctx3)
	category := request.ToCategory()
	createdCategory, err := handler.usecase.CreateCategory(ctx.Request.Context(), &category)
	if err != nil {
		ctx.Error(err)
		return
	}
	responseCategory := categoryres.CategoryResponse{}
	responseCategory.Set(createdCategory)
	response := responses.DefaultResponse{Data: gin.H{"category": responseCategory}}
	ctx.JSON(http.StatusCreated, response)
}

func (handler *CategoryHandler) DeleteCategory(ctx *gin.Context) {
	category_id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.Error(apperror.NewClientError(err))
		return
	}
	err = handler.usecase.DeleteCategory(ctx.Request.Context(), category_id)
	if err != nil {
		ctx.Error(err)
		return
	}
	response := responses.DefaultResponse{Message: "success delete category"}
	ctx.JSON(http.StatusOK, response)
}

func (handler *CategoryHandler) EditCategory(ctx *gin.Context) {
	category_id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.Error(apperror.NewClientError(err))
		return
	}
	var request categoryreq.CategoryRequest
	err = ctx.ShouldBind(&request)
	if err != nil {
		ctx.Error(apperror.NewValidationError(err))
		return
	}
	fh, _ := ctx.FormFile(enums.FileIcon)
	if fh != nil {
		ctx2 := ctx.Request.Context()
		ctx3 := context.WithValue(ctx2, enums.CategoryIconKey.Key, fh)
		ctx.Request = ctx.Request.WithContext(ctx3)
	}
	category := request.ToCategory()
	category.ID = category_id
	editedCategory, err := handler.usecase.EditCategory(ctx.Request.Context(), &category)
	if err != nil {
		ctx.Error(err)
		return
	}
	responseCategory := categoryres.CategoryResponse{}
	responseCategory.Set(editedCategory)
	response := responses.DefaultResponse{Data: gin.H{"category": responseCategory}}
	ctx.JSON(http.StatusOK, response)
}

func (handler *CategoryHandler) FindAllCategory(ctx *gin.Context) {
	var request requests.GlobalFilter
	err := ctx.ShouldBindQuery(&request)
	if err != nil {
		ctx.Error(apperror.NewServerError(err))
		return
	}
	qry := requests.NewQuery(&request)
	categories, pagination, err := handler.usecase.FindAllCategory(ctx, qry)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := []*categoryres.CategoryResponse{}
	for _, category := range categories {
		respCategory := categoryres.CategoryResponse{}
		respCategory.Set(&category)
		resp = append(resp, &respCategory)
	}
	pagination.Items = resp
	ctx.JSON(http.StatusOK, responses.DefaultResponse{
		Data: responses.NewPagination(pagination, "categories"),
	})
}

func (handler *CategoryHandler) FindByID(ctx *gin.Context) {
	category_id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.Error(apperror.NewClientError(err))
		return
	}
	category, err := handler.usecase.FindByID(ctx.Request.Context(), category_id)
	if err != nil {
		ctx.Error(err)
		return
	}
	responseCategory := categoryres.CategoryResponse{}
	responseCategory.Set(category)
	response := responses.DefaultResponse{Data: gin.H{"category": responseCategory}}
	ctx.JSON(http.StatusOK, response)
}

var _ productcategoryhandlerinterface.CategoryHandler = &CategoryHandler{}
