package adminpharmacyres

import (
	"healthcare-capt-america/entities/models"
)

type PharmacyRes struct {
	ID             uint64     `json:"id"`
	Name           string     `json:"name"`
	PharmacistName string     `json:"pharmacist_name"`
	LicenseNumber  string     `json:"license_number"`
	PhoneNumber    string     `json:"phone_number"`
	Address        AddressRes `json:"address"`
}

func (p *PharmacyRes) NewPharmacyRes(ph models.Pharmacy) PharmacyRes {
	var addressRes AddressRes
	addr := addressRes.NewAddressRes(ph.Address)
	return PharmacyRes{
		ID:             ph.ID,
		Name:           ph.Name,
		PharmacistName: ph.PharmaciestName,
		LicenseNumber:  ph.LicenseNumber,
		PhoneNumber:    ph.PhoneNumber,
		Address:        addr,
	}
}
