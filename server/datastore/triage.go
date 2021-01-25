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
		fmt.Println(err)
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
			fmt.Println(err)

		}
		rows = append(rows, map[string]interface{}{"channel": channel, "user": user, "started": started})

	}

	err = row.Err()
	if err != nil {
		fmt.Println(err)
	}
	return rows
}

// SaveTriage takes a Slack ID of a user, and the channel that they submitted
// from and saves them as the triage role for that channel
func SaveTriage(slackID string, channel string) {
	_, err = db.Query("INSERT INTO triage(slack_id, channel, started) VALUES ($1, $2, $3)", slackID, channel, time.Now())
	if err != nil {
		fmt.Println(err)
	}

}
