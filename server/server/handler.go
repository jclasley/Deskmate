package server

import (
	"fmt"
	"net/http"

	"github.com/tylerconlee/Deskmate/server/config"
)

// SlackHandler routes all callbacks from Slack
func SlackHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received for /slack endpoint")
}

func APIHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received for /slack endpoint")
}
func ConfigHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received for /config endpoint")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	switch r.Method {
	case http.MethodGet:
		config.GetConfig(w, r)
	case http.MethodPut:
		config.PutConfig(w, r)
	case http.MethodPost:
		config.PostConfig(w, r)
	}

}
