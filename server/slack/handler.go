package slack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
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

func ChannelListHandler(w http.ResponseWriter, r *http.Request) {
	channels := ListChannels()
	js, err := json.Marshal(channels)
	if err != nil {
		fmt.Println("Error marshalling JSON for channels")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func GroupListHandler(w http.ResponseWriter, r *http.Request) {
	groups := ListGroups()
	js, err := json.Marshal(groups)
	if err != nil {
		fmt.Println("Error marshalling JSON for groups")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func UserListHandler(w http.ResponseWriter, r *http.Request) {
	user := ListUsers()
	js, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error marshalling JSON for user")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// EventHandler processes the incoming callbacks set from Slack to the
// /api/slack endpoint. Slack sends an event back to this endpoint
// and it matches up to one of the events listed on their Events API
// page: https://api.slack.com/events
// Depending on the event type, Deskmate either verifies the URL or
// processes incoming text
func EventHandler(w http.ResponseWriter, r *http.Request) {
	if r.PostFormValue("payload") != "" {
		CallbackHandler(w, r)
		return
	}
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
			fmt.Println("Handling mention event", ev)
			HandleMentionEvent(ev)
		}
	}
}

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	payload := &slack.InteractionCallback{}
	err := json.Unmarshal([]byte(r.PostFormValue("payload")), payload)
	if err != nil {

	}
	fmt.Println("Handling Slack callback", payload)
	AcknowledgeTicket(payload)

	return
}

func AcknowledgeTicket(payload *slack.InteractionCallback) {
	f, _ := strconv.ParseFloat(payload.ActionCallback.BlockActions[0].ActionTs, 10)
	i := int64(f)
	ts := time.Unix(i, 0)
	fmt.Println("Ticket alert acknowledged by ", payload.User.Name)
	ts.Format(time.RFC822Z)
	t := fmt.Sprintf("<@%s> acknowledged this alert at %s", payload.User.Name, ts.String())

	ackImage := slack.NewImageBlockElement("https://emojipedia-us.s3.amazonaws.com/thumbs/120/apple/114/white-heavy-check-mark_2705.png", "white checkmark icon")
	ackText := slack.NewTextBlockObject("mrkdwn", t, false, false)

	ackSection := slack.NewContextBlock(
		"",
		[]slack.MixedElement{ackImage, ackText}...,
	)
	var respBlocks []slack.Block
	for _, block := range payload.Message.Msg.Blocks.BlockSet {
		respBlocks = append(respBlocks, block)
	}
	respBlocks = respBlocks[:len(respBlocks)-1]
	respBlocks = append(respBlocks, ackSection)
	replaceOriginal := slack.MsgOptionReplaceOriginal(payload.ResponseURL)

	opts := []slack.MsgOption{}
	opts = append(opts, slack.MsgOptionText(payload.Message.Msg.Text, false))
	opts = append(opts, slack.MsgOptionBlocks(respBlocks...))
	opts = append(opts, replaceOriginal)
	api.SendMessage(payload.Channel.ID, opts...)
}
