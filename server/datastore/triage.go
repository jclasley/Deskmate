package datastore

import (
	"database/sql"
	"time"
)

func LoadTriage() (row *sql.Rows) {
	// Load triage from database
	row, err = db.Query("SELECT slack_id, channel, started FROM triage ORDER BY started DESC; ")
	return row
}

func SaveTriage(slackID string, channel string) {
	_ = db.QueryRow("INSERT INTO triage(slack_id, channel, started) VALUES ($1, $2, $3) RETURNING id", slackID, channel, time.Now())

}
