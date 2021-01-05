package datastore

import (
	"database/sql"
	"time"
)

func GetTriage(channel string) (row *sql.Row) {
	// Load triage from database
	row = db.QueryRow("SELECT slack_id, channel FROM triage WHERE channel = $1 ORDER BY started DESC LIMIT 1; ", channel)
	return row
}

func GetAllTriage() {

}

func SaveTriage(slackID string, channel string) {
	_ = db.QueryRow("INSERT INTO triage(slack_id, channel, started) VALUES ($1, $2, $3) RETURNING id", slackID, channel, time.Now())

}
