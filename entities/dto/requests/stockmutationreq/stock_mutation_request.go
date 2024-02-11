package stockmutationreq

import "healthcare-capt-america/entities/models"

type CreateStockMutationRequest struct {
	DrugId         uint64 `json:"drug_id" binding:"required"`
	FromPharmacyId uint64 `json:"from_pharmacy_id" binding:"required"`
	ToPharmacyId   uint64 `json:"to_pharmacy_id" binding:"required"`
	Quantity       int    `json:"quantity" binding:"required"`
	StatusPharmacy string `json:"status_pharmacy,omitempty"`
	Action         string `json:"action,omitempty"`
}

func (csmr *CreateStockMutationRequest) ToStockMutation() (mutation models.StockMutation) {
	mutation.DrugId = csmr.DrugId
	mutation.FromPharmacyId = csmr.FromPharmacyId
	mutation.ToPharmacyId = csmr.ToPharmacyId
	mutation.Quantity = csmr.Quantity
	return
}
