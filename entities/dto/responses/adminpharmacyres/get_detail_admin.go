package adminpharmacyres

import "healthcare-capt-america/entities/models"

type DetailAdmin struct {
	ID          uint64        `json:"id"`
	Name        string        `json:"name"`
	Email       string        `json:"email"`
	PhoneNumber string        `json:"phone_number"`
	Pharmacies  []PharmacyRes `json:"pharmacies"`
	Photo       string        `json:"photo"`
}

func (da *DetailAdmin) NewDetailAdmin(adm *models.AdminPharmacy) DetailAdmin {
	var phoneNumber string
	if adm.User.PhoneNumber != nil {
		phoneNumber = *adm.User.PhoneNumber
	}
	pharmacies := make([]PharmacyRes, 0)
	for _, ph := range adm.Pharmacies {
		var pharmacyRes PharmacyRes
		pharmacies = append(pharmacies, pharmacyRes.NewPharmacyRes(ph))
	}

	return DetailAdmin{
		ID:          adm.ID,
		Name:        adm.User.Name,
		Email:       adm.User.Email,
		PhoneNumber: phoneNumber,
		Pharmacies:  pharmacies,
		Photo:       adm.User.Photo,
	}
}
