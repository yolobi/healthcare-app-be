package drugsreq

import (
	"healthcare-capt-america/entities/models"
	"mime/multipart"

	"github.com/shopspring/decimal"
)

type CreateDrugRequest struct {
	Content       string               `form:"content" json:"content" binding:"required"`
	Name          string               `form:"name" json:"name" binding:"required"`
	GenericName   string               `form:"generic_name" json:"generic_name" binding:"required"`
	Description   string               `form:"description" json:"description" binding:"required"`
	ManufactureID uint64               `form:"manufacture_id" json:"manufacture_id" binding:"required"`
	FormID        uint64               `form:"drug_form_id" json:"drug_form_id" binding:"required"`
	CategoryID    uint64               `form:"category_id" json:"category_id" binding:"required"`
	UnitInPack    string               `form:"unit_in_pack" json:"unit_in_pack" binding:"required"`
	Weight        *float64             `form:"weight" json:"weight" binding:"required,gt=0"`
	Height        *float64             `form:"height" json:"height" binding:"required,gt=0"`
	Length        *float64             `form:"length" json:"length" binding:"required,gt=0"`
	Width         *float64             `form:"width" json:"width" binding:"required,gt=0"`
	Image         multipart.FileHeader `form:"image" json:"image" binding:"required,gt=0"`
}

func (cdr *CreateDrugRequest) ToDrug() (drug models.Drug) {
	drug.Content = cdr.Content
	drug.Name = cdr.Name
	drug.GenericName = cdr.GenericName
	drug.Description = cdr.Description
	drug.ManufactureID = cdr.ManufactureID
	drug.FormID = cdr.FormID
	drug.CategoryID = cdr.CategoryID
	drug.UnitInPack = cdr.UnitInPack
	drug.Weight = decimal.NewFromFloat(*cdr.Weight)
	drug.Height = decimal.NewFromFloat(*cdr.Height)
	drug.Length = decimal.NewFromFloat(*cdr.Length)
	drug.Width = decimal.NewFromFloat(*cdr.Width)
	// drug.Image = cdr.Image
	return
}
