package client

import (
	"encoding/json"

	"github.com/oleewere/cldnt/model"
)

// ConvertToAirports converts bytes response to airport objects
func ConvertToAirports(bytesResponse []byte) ([]model.Airport, error) {
	airports := make([]model.Airport, 0)
	var airportResponse map[string]interface{}
	jsonErr := json.Unmarshal(bytesResponse, &airportResponse)
	if jsonErr != nil {
		return nil, jsonErr
	}

	if rowsVal, ok := airportResponse["rows"]; ok {
		rowsList := rowsVal.([]interface{})
		airport := model.Airport{}
		for _, rowVal := range rowsList {
			rowMap := rowVal.(map[string]interface{})
			if fields, ok := rowMap["fields"]; ok {
				fieldsMap := fields.(map[string]interface{})
				location := model.Location{}
				if name, ok := fieldsMap["name"]; ok {
					airport.Name = name.(string)
				}
				if lat, ok := fieldsMap["lat"]; ok {
					location.Latitude = lat.(float64)
				}
				if lon, ok := fieldsMap["lon"]; ok {
					location.Longitude = lon.(float64)
				}
				airport.Location = &location
				airports = append(airports, airport)
			}
		}
	}

	return airports, nil
}
