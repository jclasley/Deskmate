package slack

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"time"

	"github.com/circleci/Deskmate/server/datastore"
)

// Triage outlines the various users that are currently in the
// triage role within Slack. Because of multi-channel triage,
// the channel and user structs are included here to represent the
// multiple channels that could have an active triager
type Triage struct {
	Channel Channel
	User    User
	Started time.Time
}

// Channel struct defines the details about a specific channel in Slack
// that a triage role can be assumed from. The channel ID is Slack's internal
// ID that allows for the channel to be referenced programmatically. The name
// included here allows for the human-friendly channel name to be displayed
// in the frontend UI
type Channel struct {
	Name string
	ID   string
}

// User struct defines the details about a user within Slack that is currently
// in a triage role. The name is the human readable value that can be used as a
// display in the UI/Slack, and the ID is the Slack internal ID that allows for
// that user to be accessed/called programmatically
type User struct {
	Name string
	ID   string
}

// T represents the users that are currently in the triage role
var T []Triage

// DeleteTriage takes the request URI which has a channel ID in it,
// and removes the triage role associated with that channel.
func DeleteTriage(w http.ResponseWriter, r *http.Request) (n Triage) {
	user := path.Base(r.RequestURI)
	fmt.Println("Removing active triager for channel: ", user)
	removeTriage(user)
	return
}

// GetAllTriage returns the current triage object to be used on the frontend UI
// to show all users in every active channel that currently has a user in the
// triage role.
// Endpoint: GET /api/triage
func GetAllTriage(w http.ResponseWriter, r *http.Request) {
	// Add LoadTriage to retrieve data from database
	if T == nil {
		loadTriage()
	}
	triage, err := json.Marshal(T)
	if err != nil {
		fmt.Println("Error marshalling JSON for config")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(triage)

}

func activeTriage(channel string) (triage string) {
	for _, role := range T {
		if channel == role.Channel.ID {
			return role.User.ID
		}
	}
	return ""
}

func setTriage(channel string, user string) {
	// Remove the current triager for this channel if it exists
	removeTriage(channel)
	addTriage(channel, user, time.Now(), true)

}

func saveTriage(user Triage) {
	fmt.Println("Saving triage role to database")
	datastore.SaveTriage(user.User.ID, user.Channel.ID)
}

func loadTriage() {

	rows := datastore.LoadAllTriage()
	for _, row := range rows {
		channel := fmt.Sprintf("%v", row["channel"])
		user := fmt.Sprintf("%v", row["user"])
		if T == nil {
			addTriage(channel, user, row["started"].(time.Time), false)
		} else {
			exists := false
			for _, item := range T {
				if item.Channel.ID == channel {
					exists = true
				}
			}
			if !exists {
				addTriage(channel, user, row["started"].(time.Time), false)
			}
		}
	}

}

func addTriage(channel string, user string, started time.Time, save bool) {
	userInfo := getUserInfo(user)
	channelInfo := getChannelInfo(channel)
	triager := Triage{
		Channel: channelInfo,
		User:    userInfo,
		Started: started,
	}
	T = append(T, triager)
	if save {
		saveTriage(triager)
	}
	fmt.Println("Added triage: ", triager)
}

func removeTriage(channel string) {
	for i := range T {
		if T[i].Channel.ID == channel {
			T = append(T[:i], T[i+1:]...)
			break
		}
	}
}
