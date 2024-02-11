package categoryreq

import (
	"healthcare-capt-america/entities/models"
	"mime/multipart"
)

type CategoryRequest struct {
	Name string               `form:"name" binding:"required"`
	Icon multipart.FileHeader `form:"icon"`
}

func (cr *CategoryRequest) ToCategory() (category models.Category) {
	category.Name = cr.Name
	return
}
