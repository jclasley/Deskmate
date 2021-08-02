package server

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/circleci/Deskmate/server/config"
	"github.com/circleci/Deskmate/server/slack"
)

// Launch starts the webserver for Deskmate and waits for incoming requests
func Launch(c *config.Config) {
	slack.LoadConfig(c)

	router := mux.NewRouter()

	// Web App Endpoints
	sub := router.PathPrefix("/handler").Subrouter()
	// "/api/"
	sub.HandleFunc("/", APIHandler)
	// "/api/config"
	sub.HandleFunc("/config", ConfigHandler).Methods("POST", "PUT", http.MethodOptions)

	sub.HandleFunc("/slack", SlackHandler).Methods("GET", "POST", http.MethodOptions)
	sub.HandleFunc("/slack/status", slack.StatusHandler).Methods("GET", http.MethodOptions)
	sub.HandleFunc("/slack/callback", SlackCallbackHandler).Methods("GET", "POST", http.MethodOptions)

	sub.HandleFunc("/zendesk/status", ZendeskStatusHandler)

	sub.HandleFunc("/zendesk/connect", ZendeskConnectHandler)

	sub.HandleFunc("/triage/{id}", TriageHandler).Methods("POST", "DELETE", http.MethodOptions)

	sub.HandleFunc("/triage", TriageAllHandler)

	// "/api/tags"
	sub.HandleFunc("/tags", TagsHandler).Methods("POST", "PUT", "DELETE", http.MethodOptions)

	sub.HandleFunc("/tags/{id}", TagHandler).Methods("PUT", "DELETE", http.MethodOptions)

	router.Use(mux.CORSMethodMiddleware(router))
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	// origin := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	http.ListenAndServe(":8080", handlers.CORS(headers)(router))
}
