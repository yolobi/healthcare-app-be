package services

import (
	"encoding/json"
	"fmt"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/models/transaction"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func RajaongkirCost(origin uint64, destination uint64, weight int64) ([]*transaction.ShipmentFee, error) {
	url := "https://api.rajaongkir.com/starter/cost"
	payload := strings.NewReader(fmt.Sprintf("origin=%d&destination=%d&weight=%d&courier=jne", origin, destination, weight))
	req, err := http.NewRequest(http.MethodPost, url, payload)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	req.Header.Add("key", os.Getenv("RAJAONGKIR_API_KEY"))
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	var jsonResp RajaOngkir
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	response := make([]*transaction.ShipmentFee, 0)
	for _, cost := range jsonResp.Resp.Results[0].Costs {
		fee := strconv.Itoa(cost.CostDetails[0].Value)
		response = append(response, &transaction.ShipmentFee{
			Name: "JNE-" + cost.Service,
			Fee:  fee,
		})
	}
	return response, nil
}

type Cost struct {
	Service     string `json:"service"`
	CostDetails []struct {
		Value int `json:"value"`
	} `json:"cost"`
}

type Results struct {
	Costs []Cost `json:"costs"`
}

type RajaOngkirResp struct {
	Results []Results `json:"results"`
}

type RajaOngkir struct {
	Resp RajaOngkirResp `json:"rajaongkir"`
}
