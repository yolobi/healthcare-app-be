package seeding

import "healthcare-capt-america/entities/models"

type Pharmacy struct {
	AddressID       uint64 `csv:"address_id"`
	Name            string `csv:"name"`
	PharmaciestName string `csv:"pharmaciest_name"`
	LicenseNumber   string `csv:"license_number"`
	PhoneNumber     string `csv:"phone_number"`
	AdminPharmacyId uint64 `csv:"admin_pharmacy_id"`
}

func ModelPharmacy(inputs []*Pharmacy) (result []*models.Pharmacy) {
	for _, input := range inputs {
		var pharmacy = models.Pharmacy{}
		pharmacy.AddressID = input.AddressID
		pharmacy.Name = input.Name
		pharmacy.PharmaciestName = input.PharmaciestName
		pharmacy.LicenseNumber = input.LicenseNumber
		pharmacy.PhoneNumber = input.PhoneNumber
		pharmacy.AdminPharmacyId = input.AdminPharmacyId
		result = append(result, &pharmacy)
	}
	return
}
