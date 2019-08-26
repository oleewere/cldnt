package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/oleewere/cldnt/model"
)

// CalculateLocationFromIP get geo location data by using public ip
func CalculateLocationFromIP() (*model.Location, error) {
	res, _ := http.Get("https://api.ipify.org")
	ip, _ := ioutil.ReadAll(res.Body)

	geoRes, _ := http.Get(fmt.Sprintf("http://ip-api.com/json/%v", string(ip)))
	locationResponse, _ := ioutil.ReadAll(geoRes.Body)

	geoMap := make(map[string]interface{})
	jsonErr := json.Unmarshal(locationResponse, &geoMap)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return &model.Location{Latitude: geoMap["lat"].(float64), Longitude: geoMap["lon"].(float64)}, nil
}

// ListAirportsByDistance list airports based geo location distance
func ListAirportsByDistance(client Client, location model.Location, rows int, airportSearchUrl string) ([]model.Airport, error) {
	rangeValue := model.RangeValue{RangeStart: "0", RangeEnd: "90"}
	fieldQuery := model.FieldQuery{FieldName: "lat", RangeValue: &rangeValue}

	query := model.CreateQuery(fieldQuery.ToQueryString())
	query.AddDistanceSort(&location)
	query.AddLimit(rows)

	request, err := CreateGetRequest(*query.Params, airportSearchUrl)
	if err != nil {
		return nil, err
	}
	byteResponse, processErr := client.ProcessRequest(request)
	if processErr != nil {
		return nil, processErr
	}
	var airportResponse model.AirportResponse
	jsonErr := json.Unmarshal(byteResponse, &airportResponse)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return airportResponse.Airports, nil
}
