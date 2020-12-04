package datastore

import (
	"fmt"
	"strings"
)

// checkTable is called after the connection to Postgres is made. It
// checks to see if the necessary tables for Deskmate are created, and
// creates them if they're not available.
func checkTable() {
	var count int

	// Create a slice of table names to add to the query. When new tables are
	// added, they need to be added to this string slice to be verified.
	tables := []string{"config", "tags", "users"}
	t := "'" + strings.Join(tables, "',")

	// Prepare a query that will look at the information_schema list, which
	// has the list of tables in Postgres. As more tables are added, they'll
	// need to be added to the query.
	query := fmt.Sprintf("SELECT COUNT(*) AS tables_found_count FROM `information_schema`.`tables` WHERE `TABLE_SCHEMA` = '%s' AND	`TABLE_NAME` IN (%s)",
		config.host,
		t)

	// Execute the query and see if the query
	err = db.QueryRow(query).Scan(&count)
	if err != nil {
		// TODO: Log error and potentially exit Deskmate
	}

	// Check to see if the count returned from Postgres matches the number of
	// tables defined
	if count != len(tables) {
		createTables()
	}
}

// createTable creates the necessary tables for Deskmate to run if they're
// not already in the Postgres database. These tables include:
// - Config
// - Tags
// - Triager
func createTables() {
	createConfigTable()
}

// Create the table that Deskmate's configuration is stored in if the table
// does not already exist. This configuration contains the Slack API key,
// Zendesk connection details,
func createConfigTable() {
	const query = `
	CREATE TABLE IF NOT EXISTS configuration (
		slack_api text
		zendesk_url text
		zendesk_user text
		zendesk_api text	
	)`
	// Exec executes a query without returning any rows.
	if _, err := db.Exec(query); err != nil {
		return
	}

	return
}
