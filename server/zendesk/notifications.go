package zendesk

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/circleci/Deskmate/server/slack"
	"github.com/circleci/Deskmate/server/tags"
)

// Sent is a collection of all NotifySent tickets that is checked before each // notification is sent.
var Sent = []NotifySent{}

type Notify struct {
	channel          string
	tag              string
	notificationType string
}

// NotifySent is represetative of an individual ticket, what kind of
// notification was last sent for that ticket, and when the SLA breach time is.
type NotifySent struct {
	ID         int
	Type       int64
	Expire     time.Time
	LastUpdate time.Time
	Channel    string
}

func sendSLANotification(ticket Ticket, channel string, tag string) {
	url := fmt.Sprintf("https://%s.zendesk.com/agent/tickets/%d", string(c.url), ticket.ID)
	if ticket.SLA != "" {
		send, notify := UpdateCache(ticket, channel)
		log.Debugw("Processing SLA notification", "ticket", ticket.ID, "sla", ticket, "send", send)
		if send {
			message, _ := prepSLANotification(ticket, notify, tag)
			getUser(&ticket)
			notification := map[string]interface{}{
				"ID":            ticket.ID,
				"Subject":       ticket.Subject,
				"CreatedAt":     ticket.CreatedAt,
				"Assignee":      ticket.User,
				"Email":         ticket.Email,
				"TimeRemaining": message,
				"Channel":       channel,
				"Tag":           tag,
				"SLA":           ticket.SLA,
				"URL":           url,
			}
			slack.SLANotification(notification)

		}
	}
}

func sendNewNotification(ticket Ticket, channel string, tag string) {
	log.Debugw("Processing New notification", "ticket", ticket.ID, "created", ticket.CreatedAt, "last-ran", lastRan)
	if ticket.CreatedAt.After(lastRan.Add(-(2 * time.Minute))) {
		url := fmt.Sprintf("https://%s.zendesk.com/agent/tickets/%d", string(c.url), ticket.ID)
		notification := map[string]interface{}{
			"ID":        ticket.ID,
			"Subject":   ticket.Subject,
			"CreatedAt": ticket.CreatedAt,
			"Channel":   channel,
			"Tag":       tag,
			"SLA":       ticket.SLA,
			"URL":       url,
		}
		slack.NewNotification(notification)
	}
}

func sendUpdatedNotification(ticket Ticket, channel string, tag string) {

	log.Debugw("Processing Updated notification", "ticket", ticket.ID, "created", ticket.CreatedAt, "last-ran", lastRan)
	if ticket.UpdatedAt.After(lastRan.Add(-(2 * time.Minute))) {
		url := fmt.Sprintf("https://%s.zendesk.com/agent/tickets/%d", string(c.url), ticket.ID)
		getUser(&ticket)
		notification := map[string]interface{}{
			"ID":        ticket.ID,
			"Subject":   ticket.Subject,
			"CreatedAt": ticket.CreatedAt,
			"UpdatedAt": ticket.UpdatedAt,
			"Assignee":  ticket.User,
			"Email":     ticket.Email,
			"Channel":   channel,
			"Tag":       tag,
			"SLA":       ticket.SLA,
			"URL":       url,
		}
		slack.UpdatedNotification(notification)
	}
}

// PrepSLANotification takes a given ticket and what notification level and returns a string to be sent to Slack.
func prepSLANotification(ticket Ticket, notify int64, tag string) (notification string, color string) {
	var t, c string

	switch notify {
	case 1:
		t = "15 minutes"
		c = "danger"
	case 2:
		t = "30 minutes"
		c = "warning"
	case 3:
		t = "1 hour"
		c = "#ffec1e"
	case 4:
		t = "2 hours"
		c = "#439fe0"
	case 5:
		t = "3 hours"
		c = "#43e0d3"
	}

	return t, c

}

// GetTimeRemaining takes an instance of a ticket and returns the value of the
// next SLA breach.
func GetTimeRemaining(ticket Ticket) (remain time.Time) {

	breach, err := time.Parse(time.RFC3339, ticket.SLA)
	if nil != err {
		log.Errorw("Error parsing time on ticket", "error", err.Error())
		return
	}
	return breach
}

// GetNotifyType - Based off of the time remaining on the ticket, return a
// integer representing the closest time marker to a notification time.
func GetNotifyType(remain time.Duration) (notifyType int64) {
	p, _ := time.ParseDuration("3h")
	q, _ := time.ParseDuration("2h")
	r, _ := time.ParseDuration("1h")
	s, _ := time.ParseDuration("30m")
	t, _ := time.ParseDuration("15m")
	u, _ := time.ParseDuration("0m")

	switch {
	case remain < t:
		return 1
	case remain < s:
		return 2
	case remain < r:
		return 3
	case remain < q:
		return 4
	case remain < p:
		return 5
	case remain < u:
		return 9
	default:
		return 0
	}
}

// UpdateCache checks the time remaining on a ticket, what the closest marker
// for notifications is, and then checks to see if that ticket ID and
// notification type have been sent already. If yes, it returns True,
// indicating a notifcation needs to be sent.
func UpdateCache(ticket Ticket, channel string) (bool, int64) {
	cleanCache(ticket)
	// get the expiration timestamp
	expire := GetTimeRemaining(ticket)
	notify := GetNotifyType(time.Until(expire))

	// take the ticket expiration time and add 15 minutes
	expireTime := expire.Add(15 * time.Minute)

	// if the ticket expiration time is after 15 minutes from now and there's a
	// valid notification type
	if expireTime.After(time.Now()) && notify != 0 {
		rangeOnMe := reflect.ValueOf(Sent)
		for i := 0; i < rangeOnMe.Len(); i++ {
			s := rangeOnMe.Index(i)
			f := s.FieldByName("ID")
			if f.IsValid() {
				if f.Interface() == ticket.ID && s.FieldByName("Type").Int() == notify && s.FieldByName("Channel").String() == channel {
					log.Debugf("%d has already received a notification", ticket.ID)
					return false, 0
				}

			}

		}
		Sent = append(Sent, NotifySent{ticket.ID, notify, expire, ticket.UpdatedAt, channel})

		return true, notify
	}

	return false, 0

}

// cleanCache checks the Sent slice and loops through the tickets listed. If
// any have gone 15 minutes past the expiration time, they are removed from the
// slice and the length of the slice is shortened.
func cleanCache(ticket Ticket) {
	for i := 0; i < len(Sent); i++ {
		item := Sent[i]
		if ticket.ID == item.ID {
			expireTime := item.Expire.Add(15 * time.Minute)

			d := 1 * time.Minute
			sentupdate := item.LastUpdate.Truncate(d)
			ticketupdate := ticket.UpdatedAt.Truncate(d)

			if expireTime.Before(time.Now()) || sentupdate.Before(ticketupdate) {

				Sent = append(Sent[:i], Sent[i+1:]...)
				i--
			}
		}
	}

}

func checkTag(ticket Ticket) (n []Notify) {
	for _, tag := range tags.T {

		if contains(ticket.Tags, tag.Tag) {
			if tag.GroupID == strconv.Itoa(ticket.GroupID) || strconv.Itoa(ticket.GroupID) == "" {
				n = append(n, Notify{
					channel:          tag.Channel,
					notificationType: tag.NotificationType,
					tag:              tag.Tag,
				})
			} else {
				log.Debugf("Group does not match. have: %v, want: %v", tag.GroupID, ticket.GroupID)
			}
		}
	}
	return n
}
func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
