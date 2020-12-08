package config

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Config struct {
	SlackAPI string
}

func GetConfig(w http.ResponseWriter, r *http.Request) {
	config := Config{SlackAPI: "test"}
	js, err := json.Marshal(config)
	if err != nil {
		fmt.Println("Error marshalling JSON for config")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func PutConfig(w http.ResponseWriter, r *http.Request) {

}
func PostConfig(w http.ResponseWriter, r *http.Request) {

}
