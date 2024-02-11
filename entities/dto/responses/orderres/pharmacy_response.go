package orderres

import "healthcare-capt-america/entities/models"

type PharmacyOrder struct {
	ID             uint64       `json:"id"`
	Name           string       `json:"name"`
	PharmacistName string       `json:"pharmacist_name"`
	LicenseNumber  string       `json:"license_number"`
	PhoneNumber    string       `json:"phone_number"`
	Address        AddressOrder `json:"address"`
}

func newPharmacyOrderResponse(pharmacy *models.Pharmacy) *PharmacyOrder {
	return &PharmacyOrder{
		ID:             pharmacy.ID,
		Name:           pharmacy.Name,
		PharmacistName: pharmacy.PharmaciestName,
		LicenseNumber:  pharmacy.LicenseNumber,
		PhoneNumber:    pharmacy.PhoneNumber,
		Address:        *newAddressOrder(&pharmacy.Address),
	}
}
