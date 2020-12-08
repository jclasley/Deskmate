package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Launch starts the webserver for Deskmate and waits for incoming requests
func Launch() {
	r := mux.NewRouter()
	r.HandleFunc("/slack", SlackHandler)

	// Web App Endpoints
	s := r.PathPrefix("/api").Subrouter()
	// "/api/"
	s.HandleFunc("/", APIHandler)
	// "/api/config"
	s.HandleFunc("/config", ConfigHandler).Method("GET", "POST", "PUT")

	http.ListenAndServe(":8080", r)
}
