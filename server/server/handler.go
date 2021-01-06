package server

import (
	"fmt"
	"net/http"

	"github.com/tylerconlee/Deskmate/server/config"
	"github.com/tylerconlee/Deskmate/server/slack"
)

func APIHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received for /api endpoint")
}

// SlackHandler routes all event callbacks from Slack
func SlackHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received for /slack endpoint")
	slack.EventHandler(w, r)
}

func SlackStatusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received for /slack/status endpoint")
	slack.StatusHandler(w, r)
}

func SlackConnectHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received for /slack/connect endpoint")
	slack.ConnectHandler(w, r)
}

func TriageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received for /triage endpoint")
}

func ConfigHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received for /config endpoint")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	switch r.Method {
	case http.MethodGet:
		fmt.Println("GET method request for /config endpoint")
		config.GetConfig(w, r)
	case http.MethodPut:
		config.PutConfig(w, r)
	case http.MethodPost:
		fmt.Println("POST method request for /config endpoint")
		config.PostConfig(w, r)
	}

}
