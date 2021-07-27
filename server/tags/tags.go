package tags

import (
	"fmt"
	"time"

	"github.com/circleci/Deskmate/server/datastore"
)

// T represents all of the currently active tags that have been
// loaded into memory from the `loadTags` function
var T []Tag

// Tag is an individual tag used to power the notification system
type Tag struct {
	ID               int
	Tag              string
	SlackID          string
	GroupID          string
	Channel          string
	NotificationType string
	Added            time.Time
}

func loadTags() {
	tags := datastore.LoadTags()
	T = nil
	for _, tag := range tags {
		T = append(T, Tag{
			ID:               tag["id"].(int),
			Tag:              fmt.Sprintf("%v", tag["tag"]),
			SlackID:          fmt.Sprintf("%v", tag["slackID"]),
			GroupID:          fmt.Sprintf("%v", tag["groupID"]),
			Channel:          fmt.Sprintf("%v", tag["channel"]),
			NotificationType: fmt.Sprintf("%v", tag["notificationType"]),
			Added:            tag["added"].(time.Time),
		})
	}

}

func removeTag(ID int) {
	for i := range T {
		if T[i].ID == ID {
			T = append(T[:i], T[i+1:]...)
			break
		}
	}
}
