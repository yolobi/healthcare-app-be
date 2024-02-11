package services

import (
	"encoding/json"
	"fmt"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/models/transaction"
	"log"
	"net/http"
	"os"

	"github.com/shopspring/decimal"
)

func DistanceMatrix(from transaction.Position, to transaction.Position) (*decimal.Decimal, error) {
	url := "https://maps.googleapis.com/maps/api/distancematrix/"
	output := "json"
	destinations := fmt.Sprintf("?destinations=%s%%2C%s", from.Latitude.String(), from.Longitude.String())
	origins := fmt.Sprintf("&origins=%s%%2C%s", to.Latitude.String(), to.Longitude.String())
	key := fmt.Sprintf("&key=%s", os.Getenv("GMAPS_API_KEY"))
	log.Println("PAYLOAD", url+output+destinations+origins+key)

	resp, err := http.Get(url + output + destinations + origins + key)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	var response JSONResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	log.Println(response)
	d := decimal.NewFromInt(int64(response.Rows[0].Elements[0].Distance.Value))
	return &d, nil
}

type Distance struct {
	Value int `json:"value"`
}

type Element struct {
	Distance Distance `json:"distance"`
}

type Row struct {
	Elements []Element `json:"elements"`
}

type JSONResponse struct {
	Rows []Row `json:"rows"`
}
