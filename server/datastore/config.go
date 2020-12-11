package datastore

import (
	"database/sql"
	"fmt"
)

var configID = 0

func LoadConfig() (rows *sql.Row) {
	// Load config from database
	row := db.QueryRow("SELECT slack_api,slack_url from configuration where id = 1 ")

	return row
}

func SaveConfig(data map[string]interface{}) {
	fmt.Println(data)

	err := db.QueryRow("INSERT INTO configuration(id, slack_api, slack_url) VALUES (1, $1, $2) ON CONFLICT (id) DO UPDATE SET slack_api = $1, slack_url = $2 RETURNING id ", data["slackapi"], data["slackurl"]).Scan(&configID)
	if err != nil {
		fmt.Println("error saving configuration into database", err.Error())
	}

}
