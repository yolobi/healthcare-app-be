package drugsres

import (
	"healthcare-capt-america/entities/models"
	"time"
)

type DrugResponse struct {
	ID            uint64    `json:"id"`
	Content       string    `json:"content"`
	Name          string    `json:"name"`
	GenericName   string    `json:"generic_name"`
	Description   string    `json:"description"`
	ManufactureID uint64    `json:"manufacture_id"`
	Manufacture   string    `json:"manufacture"`
	FormID        uint64    `json:"drug_form_id"`
	Form          string    `json:"drug_form"`
	CategoryID    uint64    `json:"category_id"`
	Category      string    `json:"category"`
	UnitInPack    string    `json:"unit_in_pack"`
	Weight        float64   `json:"weight"`
	Height        float64   `json:"height"`
	Length        float64   `json:"length"`
	Width         float64   `json:"width"`
	Image         string    `json:"image"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (dr *DrugResponse) Set(drug *models.Drug) {
	dr.ID = drug.ID
	dr.Content = drug.Content
	dr.Name = drug.Name
	dr.GenericName = drug.GenericName
	dr.Description = drug.Description
	dr.Manufacture = drug.Manufacture.Name
	dr.Form = drug.Form.Name
	dr.Category = drug.Category.Name
	dr.UnitInPack = drug.UnitInPack
	weight, _ := drug.Weight.Float64()
	height, _ := drug.Height.Float64()
	length, _ := drug.Length.Float64()
	width, _ := drug.Width.Float64()
	dr.Weight = weight
	dr.Height = height
	dr.Length = length
	dr.Width = width
	dr.Image = drug.Image
	dr.ManufactureID = drug.Manufacture.ID
	dr.CategoryID = drug.Category.ID
	dr.FormID = drug.Form.ID
	dr.CreatedAt = drug.CreatedAt
	dr.UpdatedAt = drug.UpdatedAt
}

func getIconDir(icon string) string {
	if icon != "" {
		return icon[len("/varmasea"):]
	}
	return ""
}
