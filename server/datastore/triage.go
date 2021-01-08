package datastore

import (
	"database/sql"
	"time"
)

func GetTriage(channel string) (row *sql.Row) {
	// Load triage from database
	row = db.QueryRow("SELECT name, slack_id, channel, channel_name, started FROM triage WHERE channel = $1 ORDER BY started DESC LIMIT 1; ", channel)
	return row
}

func GetAllTriage() {

}

func SaveTriage(name string, slackID string, channel string, channelName string) {
	_ = db.QueryRow("INSERT INTO triage(name, slack_id, channel, channel_name, started) VALUES ($1, $2, $3) RETURNING id", name, slackID, channel, channelName, time.Now())

}
