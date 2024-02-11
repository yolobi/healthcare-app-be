package adminpharmacyres

import (
	"healthcare-capt-america/entities/models"
)

type AdminResponse struct {
	Id          uint64        `json:"id"`
	Name        string        `json:"name"`
	Email       string        `json:"email"`
	PhoneNumber *string       `json:"phone_number"`
	Jobs        []PharmacyRes `json:"pharmacies"`
	Photo       string        `json:"photo"`
}

func NewAdminResponse(admin *models.AdminPharmacy) AdminResponse {
	pharmacies := make([]PharmacyRes, 0)
	for _, ph := range admin.Pharmacies {
		var pharmacyRes PharmacyRes
		pharmacies = append(pharmacies, pharmacyRes.NewPharmacyRes(ph))
	}

	return AdminResponse{
		Id:          admin.ID,
		Name:        admin.User.Name,
		Email:       admin.User.Email,
		PhoneNumber: admin.User.PhoneNumber,
		Jobs:        pharmacies,
		Photo:       admin.User.Photo,
	}
}
