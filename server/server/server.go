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
	s.HandleFunc("/config", ConfigHandler).Methods("GET", "POST", "PUT", http.MethodOptions)
	r.Use(mux.CORSMethodMiddleware(r))
	http.ListenAndServe(":8080", r)
}
