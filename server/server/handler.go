package server

import (
	"fmt"
	"net/http"

	"github.com/tylerconlee/Deskmate/server/config"
	"github.com/tylerconlee/Deskmate/server/slack"
)

// APIHandler is a base path for all API related requests
func APIHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received for /api endpoint")
}

// SlackHandler routes all event callbacks from Slack
func SlackHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received for /slack endpoint")
	slack.EventHandler(w, r)
}

// SlackStatusHandler returns a health check if Slack is connected
func SlackStatusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received for /slack/status endpoint")
	slack.StatusHandler(w, r)
}

// SlackConnectHandler routes the request to start a connection
// to the configured Slack instance
func SlackConnectHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received for /slack/connect endpoint")
	slack.ConnectHandler(w, r)
}

// TriageHandler routes the request for the triage delete endpoint
// to the DeleteTriage function
func TriageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	switch r.Method {
	case http.MethodDelete:
		fmt.Println("DELETE method request for /triage/{channel} endpoint")
		slack.DeleteTriage(w, r)
	}

}

// TriageAllHandler routes the incoming request to the triage endpoint to
// return all current triagers stored in slack.T
func TriageAllHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println("GET method request for /triage endpoint")
	slack.GetAllTriage(w, r)

}

// ConfigHandler routes requests to the various config endpoints.
// The GET request goes to config.GetConfig, which returns the current
// configuration state. The POST request goes to the config.PostConfig function
// which saves the incoming configuration
func ConfigHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	switch r.Method {
	case http.MethodGet:
		fmt.Println("GET method request for /config endpoint")
		config.GetConfig(w, r)
	case http.MethodPost:
		fmt.Println("POST method request for /config endpoint")
		config.PostConfig(w, r)
	}

}
