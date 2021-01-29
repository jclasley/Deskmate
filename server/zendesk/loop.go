package zendesk

import (
	"fmt"
	"time"

	"github.com/shurcooL/graphql"
)

var active = false
var activeTickets []Ticket
var lastRan time.Time

func Connect(host string) {

	a := fmt.Sprintf("%szendesk/query", host)
	client = graphql.NewClient(a, nil)

	fmt.Println(c.user)
	variables = map[string]interface{}{
		"user":   c.user,
		"apikey": c.apikey,
		"url":    c.url,
	}
	if !active {
		RunTimer(time.Minute)
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
		if notify != nil {
			fmt.Println("Ticket ", ticket.ID, " is processing the following notifications: ", notify, " Ticket created: ", ticket.CreatedAt, " Ticket Updated: ", ticket.UpdatedAt, " Last Loop: ", lastRan)
			for _, t := range notify {
				fmt.Println("Sorting notification type: ", t.notificationType)
				switch t.notificationType {
				case "breaches":
					fmt.Println("Processing SLA breach notification")
					sendSLANotification(ticket, t.channel, t.tag)
				case "new":
					fmt.Println("Processing new ticket notification")
					sendNewNotification(ticket, t.channel, t.tag)
				case "updates":
					fmt.Println("Processing updated ticket notification")
					sendUpdatedNotification(ticket, t.channel, t.tag)
				}
			}
		}

	}
	activeTickets = nil
	<-t.C
}
