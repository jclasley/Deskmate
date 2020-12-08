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
