package journalres

import (
	"healthcare-capt-america/entities/dto/responses/drugsres"
	"healthcare-capt-america/entities/dto/responses/pharmacies"
	"healthcare-capt-america/entities/models"
	"time"
)

type JournalResponse struct {
	ID           uint64                       `json:"id"`
	Status       string                       `json:"status"`
	Quantity     int                          `json:"quantity"`
	Drug         drugsres.DrugResponse        `json:"drug"`
	FromPharmacy *pharmacies.PharmacyResponse `json:"from_pharmacy,omitempty"`
	ToPharmacy   pharmacies.PharmacyResponse  `json:"to_pharmacy"`
	CreatedAt    time.Time                    `json:"created_at"`
}

func (jr *JournalResponse) Set(journal *models.Journal) {
	jr.ID = journal.ID
	jr.Status = journal.Status
	jr.Drug = drugsres.DrugResponse{}
	jr.Drug.Set(&journal.Drug)
	if journal.FromPharmacy != nil {
		jr.FromPharmacy = pharmacies.NewPharmacyResponse(journal.FromPharmacy)
	}
	jr.ToPharmacy = *pharmacies.NewPharmacyResponse(&journal.ToPharmacy)
	jr.Quantity = journal.Quantity
	jr.CreatedAt = journal.CreatedAt
}
