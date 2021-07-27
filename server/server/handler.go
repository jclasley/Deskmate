package server

import (
	"net/http"

	"github.com/circleci/Deskmate/server/config"
	l "github.com/circleci/Deskmate/server/log"
	"github.com/circleci/Deskmate/server/slack"
	"github.com/circleci/Deskmate/server/tags"
	"github.com/circleci/Deskmate/server/zendesk"
)

var log = l.Log

// APIHandler is a base path for all API related requests
func APIHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug("Request received for /handler endpoint")
}

// SlackHandler routes all event callbacks from Slack
func SlackHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug("Request received for /slack endpoint")
	slack.EventHandler(w, r)
}

func SlackCallbackHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug("Request received for /slack/callback endpoint")
	slack.CallbackHandler(w, r)
}

// ZendeskStatusHandler returns a health check if Zendesk is connected
func ZendeskStatusHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug("Request received for /zendesk/status endpoint")
	zendesk.StatusHandler(w, r)
}

// ZendeskConnectHandler routes the request to start a connection
// to the configured Zendesk instance
func ZendeskConnectHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug("Request received for /Zendesk/connect endpoint")
	zendesk.ConnectHandler(w, r)
}

// TriageHandler routes the request for the triage delete endpoint
// to the DeleteTriage function
func TriageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	switch r.Method {
	case http.MethodDelete:
		log.Debug("DELETE method request for /triage/{channel} endpoint")
		slack.DeleteTriage(w, r)
	}

}

// TriageAllHandler routes the incoming request to the triage endpoint to
// return all current triagers stored in slack.T
func TriageAllHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	log.Debug("GET method request for /triage endpoint")
	slack.GetAllTriage(w, r)

}

// ConfigHandler routes requests to the various config endpoints.
// The GET request goes to config.GetConfig, which returns the current
// configuration state. The POST request goes to the config.PostConfig function
// which saves the incoming configuration
func ConfigHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	switch r.Method {
	case http.MethodPost:
		log.Debug("POST method request for /config endpoint")
		config.PostConfig(w, r)
		zendesk.SetConfig()
	}

}

// TagsHandler handles all incoming requests for any of the tag
// based routes. GET goes to tags.GetAllTagsHandler, which returns
// all tags, POST goes to tags.PostTagHandler, which saves new tags,
// PUT goes to tags.UpdateTagHandler, which saves an existing tag and
// DELETE goes to tags.DeleteTagHandler which removes the tag from the
// database
func TagsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	switch r.Method {
	case http.MethodGet:
		log.Debug("GET method request for /tags endpoint")
		tags.GetAllTagsHandler(w, r)
	case http.MethodPost:
		log.Debug("POST method request for /tags/{id} endpoint")
		tags.PostTagHandler(w, r)

	}
}

// TagHandler handles the requests related to a specific tag, such as updating
// or deleting
func TagHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	switch r.Method {

	case http.MethodPut:
		log.Debug("PUT method request for /tags endpoint")
		tags.UpdateTagHandler(w, r)
	case http.MethodDelete:
		log.Debug("DELETE method request for /tags endpoint")
		tags.DeleteTagHandler(w, r)
	}
}
