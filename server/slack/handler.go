package slack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/nlopes/slack"
	"github.com/nlopes/slack/slackevents"
)

// StatusHandler is called from the front end to determine whether
// Deskmate is currently connected to a Slack instance or not.
// It calls slack.Ping(), and writes the result as JSON to be
// parsed by the frontend
// SEE: src/components/SlackConnect.js
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	info := Ping()
	js, err := json.Marshal(info)
	if err != nil {
		fmt.Println("Error marshalling JSON for config")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// ConnectHandler handles the request from the front end to establish
// a connection to Slack using the saved configuration. On the front end,
// this is triggered by the button in the top right corner.
// SEE: src/components/SlackConnect.js
// TODO: Add log or notification to indicate successful call to start connection
func ConnectHandler(w http.ResponseWriter, r *http.Request) {
	Connect()
}

// EventHandler processes the incoming callbacks set from Slack to the
// /api/slack endpoint. Slack sends an event back to this endpoint
// and it matches up to one of the events listed on their Events API
// page: https://api.slack.com/events
// Depending on the event type, Deskmate either verifies the URL or
// processes incoming text
func EventHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		status = false
		return
	}
	sv, err := slack.NewSecretsVerifier(r.Header, c.Slack.SlackSigning)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		status = false
		return
	}
	if _, err := sv.Write(body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		status = false
		return
	}
	if err := sv.Ensure(); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		status = false
		return
	}
	eventsAPIEvent, err := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionNoVerifyToken())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		status = false
		return
	}
	fmt.Println("Slack event received: ", eventsAPIEvent.InnerEvent.Type)

	if eventsAPIEvent.Type == slackevents.URLVerification {
		var r *slackevents.ChallengeResponse
		err := json.Unmarshal([]byte(body), &r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			status = false
			return
		}
		w.Header().Set("Content-Type", "text")
		w.Write([]byte(r.Challenge))

	}
	if eventsAPIEvent.Type == slackevents.CallbackEvent {
		innerEvent := eventsAPIEvent.InnerEvent
		switch ev := innerEvent.Data.(type) {
		case *slackevents.AppMentionEvent:
			api.PostMessage(ev.Channel, slack.MsgOptionText("Yes, hello.", false))
		}
	}
}
