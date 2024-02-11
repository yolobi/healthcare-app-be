package stockmutationres

import (
	"healthcare-capt-america/entities/dto/responses/drugsres"
	"healthcare-capt-america/entities/dto/responses/pharmacies"
	"healthcare-capt-america/entities/models"
	"time"
)

type StockMutationResponse struct {
	ID             uint64                      `json:"id"`
	Drug           drugsres.DrugResponse       `json:"drug"`
	FromPharmacy   pharmacies.PharmacyResponse `json:"from_pharmacy"`
	ToPharmacy     pharmacies.PharmacyResponse `json:"to_pharmacy"`
	Quantity       int                         `json:"quantity"`
	StatusPharmacy string                      `json:"status_pharmacy,omitempty"`
	StatusMutation string                      `json:"status_mutation"`
	CreatedAt      time.Time                   `json:"created_at"`
	UpdatedAt      time.Time                   `json:"updated_at"`
}

func (smr *StockMutationResponse) Set(stockMutation *models.StockMutation) {
	smr.ID = stockMutation.ID
	smr.Drug.Set(&stockMutation.Drug)
	smr.FromPharmacy = *pharmacies.NewPharmacyResponse(&stockMutation.FromPharmacy)
	smr.ToPharmacy = *pharmacies.NewPharmacyResponse(&stockMutation.ToPharmacy)
	smr.Quantity = stockMutation.Quantity
	smr.StatusMutation = stockMutation.StatusMutation.Name
	smr.CreatedAt = stockMutation.CreatedAt
	smr.UpdatedAt = stockMutation.UpdatedAt
}
