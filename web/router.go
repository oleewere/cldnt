package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gobuffalo/packr"
	"github.com/gorilla/mux"
	"github.com/oleewere/cldnt/client"
	"github.com/oleewere/cldnt/model"
)

// StartServer starts web server
func StartServer(port int) {
	box := packr.NewBox("../static/")
	router := mux.NewRouter()
	router.HandleFunc("/ping", PingHandler)
	router.HandleFunc("/airports", AirportsHandler)
	router.PathPrefix("/").Handler(http.FileServer(box))
	log.Println(fmt.Sprintf("Start server on port %v", port))
	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}

// PingHandler send a simple ping back by the web server
func PingHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

// AirportsHandler get closest airports
func AirportsHandler(w http.ResponseWriter, r *http.Request) {
	lat := r.FormValue("lat")
	lon := r.FormValue("lon")
	rowStr := r.FormValue("rows")
	var rows int = 10
	var rowError error
	if len(rowStr) > 0 {
		rows, rowError = strconv.Atoi(rowStr)
		handleError(rowError, w)
	}
	if rows > 200 {
		rows = 200
	}
	var location model.Location
	if len(lat) == 0 || len(lon) == 0 {
		lotLatPair, err := client.CalculateLocationFromIP()
		handleError(err, w)
		location = *lotLatPair
	} else {
		latutide, err := strconv.ParseFloat(lat, 64)
		handleError(err, w)
		longitude, err := strconv.ParseFloat(lon, 64)
		handleError(err, w)
		location = model.Location{Latitude: latutide, Longitude: longitude}
	}
	log.Println(fmt.Sprintf("Inputs - rows: %d - latitude: %f - longitue: %f", rows, location.Latitude, location.Longitude))
	airports, errAirp := client.ListAirportsByDistance(location, rows, "")
	log.Println(fmt.Sprintf("Airports response: %v", airports))
	handleError(errAirp, w)
	json.NewEncoder(w).Encode(airports)
}

func handleError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
