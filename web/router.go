package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gobuffalo/packr"
	"github.com/gorilla/mux"
)

// StartServer starts web server
func StartServer(port int) {
	box := packr.NewBox("../static/")
	router := mux.NewRouter()
	router.HandleFunc("/ping", PingHandler)
	router.PathPrefix("/").Handler(http.FileServer(box))
	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}

// PingHandler send a simple ping back by the web server
func PingHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
