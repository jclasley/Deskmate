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
	for rows.Next() {
		err := rows.Scan(&c.Slack.SlackURL, &c.Slack.SlackAPI, &c.Zendesk.ZendeskUser, &c.Zendesk.ZendeskAPI, &c.Zendesk.ZendeskURL)
		if err != nil {
			fmt.Println("Error scanning config into struct")
		}
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
	fmt.Println("received POST request")
	r.ParseForm()
	fmt.Println(r.Form.Get("slackapi"))
	fmt.Println(r.FormValue("slackapi"))
	for k, v := range r.Form {
		fmt.Printf("%s = %s\n", k, v)
	}
}
