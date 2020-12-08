package server

import (
	"fmt"
	"github.com/tylerconlee/Deskmate/config"
	"net/http"
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
	if r.Method == http.MethodGet {
		config.LoadConfig(w, r)
	}

}
