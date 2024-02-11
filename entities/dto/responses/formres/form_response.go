package formres

import "healthcare-capt-america/entities/models"

type FormResponse struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

func NewForm(form models.Form) *FormResponse {
	return &FormResponse{ID: form.ID, Name: form.Name}
}

func NewForms(forms []models.Form) []FormResponse {
	var formRess []FormResponse
	for _, form := range forms {
		formRess = append(formRess, *NewForm(form))
	}
	return formRess
}
