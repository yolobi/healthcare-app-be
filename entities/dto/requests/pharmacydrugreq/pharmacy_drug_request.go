package pharmacydrugreq

import (
	"healthcare-capt-america/entities/models"

	"github.com/shopspring/decimal"
)

type CreatePharmacyDrugRequest struct {
	DrugId      uint64          `json:"drug_id" binding:"required"`
	Status      string          `json:"status" binding:"required"`
	Stock       *int            `json:"stock" binding:"required,gte=0"`
	SellingUnit decimal.Decimal `json:"selling_unit" binding:"required"`
}

func (pdr *CreatePharmacyDrugRequest) ToPharmacyDrug() (pharmacydrug models.PharmacyDrug) {
	pharmacydrug.DrugId = pdr.DrugId
	pharmacydrug.Stock = *pdr.Stock
	pharmacydrug.SellingUnit = pdr.SellingUnit
	pharmacydrug.Status = pdr.Status
	return
}

type EditPharmacyDrugRequest struct {
	Status      string          `json:"status" binding:"required"`
	Stock       *int            `json:"stock" binding:"required,gte=0"`
	SellingUnit decimal.Decimal `json:"selling_unit" binding:"required"`
}

func (pdr *EditPharmacyDrugRequest) ToPharmacyDrug() (pharmacydrug models.PharmacyDrug) {
	pharmacydrug.Stock = *pdr.Stock
	pharmacydrug.SellingUnit = pdr.SellingUnit
	pharmacydrug.Status = pdr.Status
	return
}

type EditPharmacysDrugRequest struct {
	PharmacyID  uint64          `json:"pharmacy_id" binding:"required"`
	Stock       int             `json:"stock" binding:"required"`
	SellingUnit decimal.Decimal `json:"selling_unit"`
	Status      string          `json:"status"`
}
