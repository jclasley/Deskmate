package server

import (
	"fmt"
	"net/http"
)

// IndexRouter serves the frontend for Deskmate
func IndexRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received for / endpoint")
}

// SlackHandler routes all callbacks from Slack
func SlackHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received for /slack endpoint")
}
