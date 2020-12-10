package config

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tylerconlee/Deskmate/server/datastore"
)

var c Config

func GetConfig(w http.ResponseWriter, r *http.Request) {
	rows := datastore.LoadConfig()
	err := rows.Scan(&c.Slack.SlackURL, &c.Slack.SlackAPI)
	if err != nil {
		fmt.Println("Error scanning config into struct", err.Error())
	}
	js, err := json.Marshal(c)
	if err != nil {
		fmt.Println("Error marshalling JSON for config")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func PutConfig(w http.ResponseWriter, r *http.Request) {

}
func PostConfig(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err.Error())
	}
	datastore.SaveConfig(data)
}
