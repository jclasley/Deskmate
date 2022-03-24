package datastore

// checkTable is called after the connection to Postgres is made. It
// checks to see if the necessary tables for Deskmate are created, and
// creates them if they're not available.
func checkTable() {
	createConfigTable()
	createTriageTable()
	createTagsTable()
	log.Info("Datastore tables successfully processed")
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

		log.Fatalw("Table creation error for Config", "error", err.Error())
		return
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		log.Fatalw("Table creation error for Config", "error", err.Error())
		return
	}
	if affectedRows != 0 {
		log.Debug("Configuration table successfully created")
	}

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
		log.Fatalw("Table creation error for Triage", "error", err.Error())
		return

	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		log.Fatalw("Table creation error for Triage", "error", err.Error())
		return
	}
	if affectedRows != 0 {
		log.Debug("Triage table successfully created")
	}

}

func createTagsTable() {
	const query = `
	CREATE TABLE IF NOT EXISTS tags (
		id serial PRIMARY KEY,
		tag text,
		slack_id text,
		group_id text,
		channel text,
		notification_type text,
		added timestamp
	)`
	// Exec executes a query without returning any rows.
	result, err := db.Exec(query)
	if err != nil {
		log.Fatalw("Table creation error for Tags", "error", err.Error())
		return

	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		log.Fatalw("Table creation error for Tagts", "error", err.Error())

		return
	}
	if affectedRows != 0 {
		log.Debug("Tags table successfully created")
	}

}
