package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Launch starts the webserver for Deskmate and waits for incoming requests
func Launch() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexRouter)
	r.HandleFunc("/slack", SlackHandler)
	http.ListenAndServe(":8000", r)
}
