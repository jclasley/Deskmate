package datastore

import (
	"database/sql"
	"fmt"
)

func LoadConfig() (rows *sql.Rows) {
	// Load config from database
	rows, err := db.Query("SELECT  slack_api,slack_url, zendesk_user, zendesk_api, zendesk_url from configuration")
	if err != nil {
		fmt.Println("Error retrieving configuration from database")
	}
	return rows
}

func SaveConfig(data map[string]interface{}) {
	fmt.Println(data)

	rows, err := db.Query("INSERT INTO configuration(slack_api, slack_url) VALUES ($1, $2);", data["slackapi"], data["slackurl"])
	if err != nil {
		fmt.Println("error saving configuration into database", err.Error())
	}

	defer rows.Close()

}
