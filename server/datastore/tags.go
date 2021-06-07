package datastore

import (
	"fmt"
	"time"
)

// LoadTags returns an array of map[string]interfaces that
// contain all of the tags current stored in the `tags`
// table of the database
func LoadTags() (t []map[string]interface{}) {
	row, err := db.Query("SELECT * from tags ORDER BY added DESC;")
	if err != nil {
		fmt.Println(err)
	}
	defer row.Close()
	for row.Next() {
		var (
			id               int
			tag              string
			slackID          string
			groupID          string
			channel          string
			notificationType string
			added            time.Time
		)
		if err := row.Scan(&id, &tag, &slackID, &groupID, &channel, &notificationType, &added); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			log.Fatalw("Unable to load tags from datastore", "error", err.Error())
			return
		}
		t = append(t, map[string]interface{}{"id": id, "tag": tag, "slackID": slackID, "groupID": groupID, "channel": channel, "notificationType": notificationType, "added": added})

	}

	err = row.Err()
	if err != nil {
		log.Fatalw("Unable to load tags from datastore", "error", err.Error())
		return
	}
	return t
}

// CreateTag takes the metadata needed for saving a tag and
// saves the data to the `tags` table in the database
func CreateTag(t map[string]interface{}) {
	_, err = db.Query("INSERT INTO tags(tag, slack_id, group_id, channel, notification_type, added) VALUES ($1, $2, $3, $4, $5, $6)", t["tag"], t["slackID"], t["groupID"], t["channel"], t["notificationType"], t["added"])
	if err != nil {
		log.Errorw("Unable to save tags to datastore", "error", err.Error())
		return
	}
}

// RemoveTag removes the specified tag from the database so that
// it isn't loaded on the next restart
func RemoveTag(tagID int) {
	_, err = db.Query("DELETE FROM tags WHERE id = $1", tagID)
	if err != nil {
		log.Errorw("Unable to remove tags from datastore", "error", err.Error())
		return
	}
}

// UpdateTag updates the specified tag with the provided metadata
func UpdateTag(t map[string]interface{}) {
	_, err = db.Query("UPDATE tags SET slack_id = $1, group_id = $2, channel =$3, notification_type = $4, added = $5 WHERE id = $6", t["slackID"], t["groupID"], t["channel"], t["notificationType"], t["added"], t["id"])
	if err != nil {
		log.Errorw("Unable to update tags in datastore", "error", err.Error())
		return
	}

}
