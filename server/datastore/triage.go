package datastore

import (
	"fmt"
	"time"
)

// LoadAllTriage pulls all triage roles from the database in the triage table
// and returns it as a slice of map[string]interface
func LoadAllTriage() (rows []map[string]interface{}) {
	// Load triage from database
	row, err := db.Query("SELECT slack_id, channel, started FROM triage ORDER BY started DESC; ")
	if err != nil {
		log.Fatalw("Unable to load triage from datastore", "error", err.Error())
	}
	defer row.Close()

	for row.Next() {
		var (
			channel string
			user    string
			started time.Time
		)
		if err := row.Scan(&user, &channel, &started); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			log.Fatalw("Unable to scan rows for triage", "error", err.Error())

		}
		rows = append(rows, map[string]interface{}{"channel": channel, "user": user, "started": started})

	}

	err = row.Err()
	if err != nil {
		log.Fatalw("Unable to scan rows for triage", "error", err.Error())
	}
	return rows
}

// SaveTriage takes a Slack ID of a user, and the channel that they submitted
// from and saves them as the triage role for that channel
func SaveTriage(slackID string, channel string) {
	_, err = db.Query("INSERT INTO triage(slack_id, channel, started) VALUES ($1, $2, $3)", slackID, channel, time.Now())
	if err != nil {
		log.Fatalw("Unable to save triage to datastore", "error", err.Error())
	}
}

const durationQuery = `select current_timestamp - (select started from triage
	where slack_id='%s' order by started desc limit 1)`

// SetTriageDuration is intended to be called every time the active triager changes.
// This will change if there is either a call to the `unset` command or if a new triager
// comes online.
func SetTriageDuration(slackID string, channel string) {
	durQuery := fmt.Sprintf(durationQuery, slackID)
	query := "update triage set triage_interval=(%s) where triage.slack_id=%s and triage.channel=%s"
	query = fmt.Sprintf(query, durQuery, slackID, channel)
	_, err = db.Query(query)
	if err != nil {
		log.Fatalf("error in updating duration: %q", err.Error())
	}
}
