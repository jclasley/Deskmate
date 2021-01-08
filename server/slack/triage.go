package slack

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Triage outlines the various users that are currently in the
// triage role within Slack. Because of multi-channel triage,
// the channel and user structs are included here to represent the
// multiple channels that could have an active triager
type Triage struct {
	Channels []Channel
	Users    []User
	Started  []time.Time
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
var T Triage

// GetTriage gets the triage user details for a specific channel and returns it
// in a JSON format so it can be parsed in Slack and the frontend.
// Endpoint: GET /api/triage/{channel-id}
func GetTriage(w http.ResponseWriter, r *http.Request) (T Triage) {
	return
}

// GetAllTriage returns the current triage object to be used on the frontend UI
// to show all users in every active channel that currently has a user in the
// triage role.
// Endpoint: GET /api/triage
func GetAllTriage(w http.ResponseWriter, r *http.Request) {

}

// PostTriage submits a new user for the triage role in a specific channel and
// saves it to the database. It also updates the main Triage object, T, so that
// the new user is reflected as on the triage role across the board.
// Endpoint: POST /api/triage/{channel-id}
func PostTriage(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error processing channel ID")
	}
	fmt.Println(body)
}
