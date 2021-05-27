package datastore

import (
	"database/sql"
	"fmt"
)

var configID = 0

// LoadConfig pulls the configuration details from the database and returns a
// pointer to a sql.Row
func LoadConfig() (rows *sql.Row) {
	// Load config from database
	row := db.QueryRow("SELECT slack_url,slack_api,slack_signing, zendesk_user, zendesk_api, zendesk_url from configuration where id = 1 ")
	return row
}

// SaveConfig takes a map[string]interface containing configuration details
// and saves the key/values to the database
func SaveConfig(data map[string]interface{}) {
	err := db.QueryRow(`INSERT INTO configuration
	(
		id, 
		slack_api, 
		slack_url, 
		slack_signing,
		zendesk_url, 
		zendesk_user,
		zendesk_api
	) VALUES (
		1, 
		$1, 
		$2,
		$3,
		$4,
		$5,
		$6
	) ON CONFLICT (id) 
	DO UPDATE SET 
	slack_api = $1, 
	slack_url = $2,
	slack_signing = $3,
	zendesk_url = $4,
	zendesk_user = $5,
	zendesk_api = $6
	RETURNING id `,
		data["slackapi"],
		data["slackurl"],
		data["slacksigning"],
		data["zendeskurl"],
		data["zendeskuser"],
		data["zendeskapi"]).Scan(&configID)
	if err != nil {
		fmt.Println("error saving configuration into database", err.Error())
	}

}
