package server

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/tylerconlee/Deskmate/server/slack"
)

// Launch starts the webserver for Deskmate and waits for incoming requests
func Launch() {
	slack.LoadConfig()

	r := mux.NewRouter()

	// Web App Endpoints
	s := r.PathPrefix("/api").Subrouter()
	// "/api/"
	s.HandleFunc("/", APIHandler)
	// "/api/config"
	s.HandleFunc("/config", ConfigHandler).Methods("GET", "POST", "PUT", http.MethodOptions)

	s.HandleFunc("/slack", SlackHandler).Methods("GET", "POST", http.MethodOptions)

	s.HandleFunc("/slack/status", SlackStatusHandler)

	r.Use(mux.CORSMethodMiddleware(r))
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	origin := handlers.AllowedOrigins([]string{"*"})
	http.ListenAndServe(":8080", handlers.CORS(headers, origin)(r))
}
