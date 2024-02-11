package seeding

import (
	"healthcare-capt-america/entities/models"

	"github.com/shopspring/decimal"
)

type PharmacyDrug struct {
	DrugId      uint64          `csv:"drug_id"`
	PharmacyId  uint64          `csv:"pharmacy_id"`
	Stock       int             `csv:"stock"`
	SellingUnit decimal.Decimal `csv:"selling_unit"`
	Status      string          `csv:"status"`
}

func ModelPharmacyDrug(inputs []*PharmacyDrug) (result []*models.PharmacyDrug) {
	for _, input := range inputs {
		var pharmacyDrug = models.PharmacyDrug{}
		pharmacyDrug.DrugId = input.DrugId
		pharmacyDrug.PharmacyId = input.PharmacyId
		pharmacyDrug.Stock = input.Stock
		pharmacyDrug.SellingUnit = input.SellingUnit
		pharmacyDrug.Status = input.Status
		result = append(result, &pharmacyDrug)
	}
	return
}
