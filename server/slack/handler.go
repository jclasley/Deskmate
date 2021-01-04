package slack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/nlopes/slack"
	"github.com/nlopes/slack/slackevents"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(status)
	if err != nil {
		fmt.Println("Error marshalling JSON for config")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func ConnectHandler(w http.ResponseWriter, r *http.Request) {
	Connect()
}

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
	status = true
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
