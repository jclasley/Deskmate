package datastore

import (
	"fmt"
)

// checkTable is called after the connection to Postgres is made. It
// checks to see if the necessary tables for Deskmate are created, and
// creates them if they're not available.
func checkTable() {
	createConfigTable()
	createTriageTable()
	createTagsTable()
	fmt.Println("Tables successfully loaded/created")
}

// Create the table that Deskmate's configuration is stored in if the table
// does not already exist. This configuration contains the Slack API key,
// Zendesk connection details,
func createConfigTable() {
	const query = `
	CREATE TABLE IF NOT EXISTS configuration (
		id serial PRIMARY KEY,
		slack_api text,
		slack_url text,
		slack_signing text,
		zendesk_url text,
		zendesk_user text,
		zendesk_api text
	)`
	// Exec executes a query without returning any rows.
	result, err := db.Exec(query)
	if err != nil {
		fmt.Println("Error creating configuration table", err.Error())
		return
	}
	a, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err.Error())
	}
	if a != 0 {
		fmt.Println("Configuration table successfully created.", a)
	}
	return
}

func createTriageTable() {
	const query = `
	CREATE TABLE IF NOT EXISTS triage (
		id serial PRIMARY KEY,
		name text,
		slack_id text,
		channel text,
		started timestamp
	)`
	// Exec executes a query without returning any rows.
	result, err := db.Exec(query)
	if err != nil {
		fmt.Println("Error creating triage table", err.Error())
		return
	}
	a, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err.Error())
	}
	if a != 0 {
		fmt.Println("Triage table successfully created.", a)
	}
	return
}

func createTagsTable() {
	const query = `
	CREATE TABLE IF NOT EXISTS tags (
		id serial PRIMARY KEY,
		tag text,
		slack_id text,
		group_id text,
		channel text,
		notification_type int,
		added timestamp
	)`
	// Exec executes a query without returning any rows.
	result, err := db.Exec(query)
	if err != nil {
		fmt.Println("Error creating tags table", err.Error())
		return
	}
	a, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err.Error())
	}
	if a != 0 {
		fmt.Println("Tags table successfully created.", a)
	}
	return
}
