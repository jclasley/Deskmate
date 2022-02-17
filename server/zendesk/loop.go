package zendesk

import (
	"fmt"
	"os"
	"strings"
	"time"

	l "github.com/circleci/Deskmate/server/log"
	"github.com/circleci/Deskmate/server/slack"
	"github.com/shurcooL/graphql"
)

var active = false
var activeTickets []Ticket
var lastRan time.Time
var log = l.Log

func Connect(host string) {
	var a string
	if os.Getenv("APP_ENV") == "development" {
		host = strings.Replace(host, "3", "6", 1)

		a = fmt.Sprintf("%squery", host)
	} else {
		a = fmt.Sprintf("%szendesk/query", host)
	}
	client = graphql.NewClient(a, nil)

	variables = map[string]interface{}{
		"user":   c.user,
		"apikey": c.apikey,
		"url":    c.url,
	}
	if !active {
		go RunTimer(time.Minute)
	}
}

func RunTimer(interval time.Duration) {
	t := time.NewTicker(interval)
	active = true
	for {
		iteration(t, interval)
		<-t.C

	}
}

func iteration(t *time.Ticker, interval time.Duration) {

	lastRan = time.Now()
	getAllTickets()
	for _, ticket := range activeTickets {

		notify := checkTag(ticket)
		for _, t := range notify {

			switch t.notificationType {
			case "breaches":

				sendSLANotification(ticket, t.channel, t.tag)
			case "new":

				sendNewNotification(ticket, t.channel, t.tag)
			case "updates":

				sendUpdatedNotification(ticket, t.channel, t.tag)
			}
		}
		for _, reminder := range slack.R {
			// GetActiveTriager for the channel to be reminded
			triage := slack.ActiveTriage(reminder.Channel.ID)

			// Check if the last time a reminder was sent was 15 minutes
			// before the current iteration
			// If it was, send another notification and update the last sent time
			if reminder.LastSent.Before(lastRan.Add(-(15 * time.Minute))) && triage == "" {
				slack.SendReminder(reminder.Channel.ID)
				reminder.LastSent = time.Now()
			}
		}

	}

	activeTickets = nil
	<-t.C
}
