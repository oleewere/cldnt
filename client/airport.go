package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/oleewere/cldnt/model"
)

func CalculateLocationFromIP() (*model.Location, error) {
	res, _ := http.Get("https://api.ipify.org")
	ip, _ := ioutil.ReadAll(res.Body)

	geoRes, _ := http.Get(fmt.Sprintf("http://ip-api.com/json/%v", string(ip)))
	locationResponse, _ := ioutil.ReadAll(geoRes.Body)

	geoMap := make(map[string]interface{})
	jsonErr := json.Unmarshal(locationResponse, &geoMap)
	if jsonErr != nil {
		fmt.Println(jsonErr)
		return nil, jsonErr
	}
	return &model.Location{Latitude: geoMap["lat"].(float64), Longitude: geoMap["lon"].(float64)}, nil
}

func ListAirportsByDistance(*model.Location) []model.Airport {
	return nil
}
