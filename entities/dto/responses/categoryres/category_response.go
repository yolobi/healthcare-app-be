package categoryres

import (
	"healthcare-capt-america/entities/models"
	"time"
)

type CategoryResponse struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Icon      string    `json:"icon"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (cr *CategoryResponse) Set(pharmacydrug *models.Category) {
	cr.ID = pharmacydrug.ID
	cr.Name = pharmacydrug.Name
	cr.Icon = pharmacydrug.Icon
	cr.CreatedAt = pharmacydrug.CreatedAt
	cr.UpdatedAt = pharmacydrug.UpdatedAt
}

func getIconDir(icon string) string {
	if icon != "" {
		icon = icon[len("/varmasea"):]
		icon = icon + "?authuser=2"
		return icon
	}
	return ""
}

type CategoryRes struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

func NewCategory(category *models.Category) *CategoryRes {
	return &CategoryRes{ID: category.ID, Name: category.Name, Icon: category.Icon}
}

func NewCategories(categories []models.Category) []CategoryRes {
	var categoriesRes []CategoryRes
	for _, category := range categories {
		categoriesRes = append(categoriesRes, *NewCategory(&category))
	}
	return categoriesRes
}
