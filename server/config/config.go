package config

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/tylerconlee/Deskmate/server/datastore"
	l "github.com/tylerconlee/Deskmate/server/log"
)

var (
	config *Config
	log    = l.Log
)

// GetConfig sends a request to the database to grab the
// contents of the configuration table. It scans it into
// an instance of 'c', the config for Deskmate. If it
// runs into any errors, it reports them to the logs.
func GetConfig(w http.ResponseWriter, r *http.Request) {
	config = LoadConfig()
	js, err := json.Marshal(config)
	if err != nil {
		log.Errorw("Error retrieving config from database", "error", err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func PostConfig(w http.ResponseWriter, r *http.Request) {
	payload := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Errorw("Error saving config to database", "error", err.Error())
		return
	}
	if payload["slackurl"] != nil {
		payload["zendeskuser"] = config.Zendesk.ZendeskUser
		payload["zendeskapi"] = config.Zendesk.ZendeskAPI
		payload["zendeskurl"] = config.Zendesk.ZendeskURL
		datastore.SaveConfig(payload)
		return
	}
	if payload["zendeskuser"] != nil {
		payload["slackurl"] = config.Slack.SlackURL
		payload["slackapi"] = config.Slack.SlackAPI
		payload["slacksigning"] = config.Slack.SlackSigning
		datastore.SaveConfig(payload)
		return
	}
}

// LoadConfig returns the configuration file as loaded from environment variables. Requires several environment variables that were previously fetched from the Postgres DB. Environment variables are set with the inclusion of a 'config.json' in the 'server' directory. This file can and should be overwritten via a bind-mount in the 'docker-compose.yml'. An example called 'config_example.json' is included for documentation of the JSON shape.
//
// SLACK_API, SLACK_SIGN, SLACK_URL, ZEN_API, ZEN_URL, ZEN_USER
func LoadConfig() *Config {
	// ! removed for the timebeing, will be re-implemented once entry to Deskmate is secured via authentication
	rows := datastore.LoadConfig()
	err := rows.Scan(&config.Slack.SlackURL, &config.Slack.SlackAPI, &config.Slack.SlackSigning, &config.Zendesk.ZendeskUser, &config.Zendesk.ZendeskAPI, &config.Zendesk.ZendeskURL)
	if err != nil {
		log.Errorw("Error retrieving config from database", "error", err.Error())
		return nil
	}
	return config
}

func unmarshalConfig(c *Config) {
	f, err := os.ReadFile("./config.json")
	if err != nil {
		log.Errorw("Error retrieving config.json", "error", err.Error())
	}
	if err := json.Unmarshal(f, c); err != nil {
		log.Errorw("Error unmarshaling json to struct", "error", err.Error())
	}
}
