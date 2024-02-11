package operationalreq

import (
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/utils"
)

type OperationRequest struct {
	PharmacyID      uint64           `json:"pharmacy_id"`
	OperationalDays []OperationalDay `json:"operational_days"`
}

type OperationalDay struct {
	Day       string `json:"day"`
	OpenTime  string `json:"open_time"`
	CloseTime string `json:"close_time"`
}

func (request *OperationRequest) OperationalModels() ([]models.Operational, error) {
	var operationals []models.Operational
	for _, day := range request.OperationalDays {
		openTimeDur, err := utils.ParseTime(day.OpenTime)
		if err != nil {
			return nil, err
		}
		closeTimeDur, err := utils.ParseTime(day.CloseTime)
		if err != nil {
			return nil, err
		}
		operational := &models.Operational{
			PharmacyID: request.PharmacyID,
			Day:        day.Day,
			OpenTime:   openTimeDur,
			CloseTime:  closeTimeDur,
		}
		operationals = append(operationals, *operational)
	}
	return operationals, nil
}
