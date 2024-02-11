package seeding

import (
	"healthcare-capt-america/entities/models"

	"github.com/shopspring/decimal"
)

type Drug struct {
	Content       string          `csv:"content"`
	Name          string          `csv:"name"`
	GenericName   string          `csv:"generic_name"`
	Description   string          `csv:"description"`
	ManufactureID uint64          `csv:"manufacture_id"`
	FormID        uint64          `csv:"form_id"`
	Image         string          `csv:"image"`
	CategoryID    uint64          `csv:"category_id"`
	UnitInPack    string          `csv:"unit_in_pack"`
	Weight        decimal.Decimal `csv:"weight"`
	Height        decimal.Decimal `csv:"height"`
	Length        decimal.Decimal `csv:"length"`
	Width         decimal.Decimal `csv:"width"`
}

func ModelProduct(inputs []*Drug) (result []*models.Drug) {
	for _, input := range inputs {
		var product = models.Drug{}
		product.Content = input.Content
		product.Name = input.Name
		product.GenericName = input.GenericName
		product.Description = input.Description
		product.ManufactureID = input.ManufactureID
		product.FormID = input.FormID
		product.CategoryID = input.CategoryID
		product.UnitInPack = input.UnitInPack
		product.Image = input.Image
		product.Weight = input.Weight
		product.Height = input.Height
		product.Length = input.Length
		product.Width = input.Width
		result = append(result, &product)
	}
	return
}
