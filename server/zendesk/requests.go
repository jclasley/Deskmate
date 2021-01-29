package zendesk

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/shurcooL/graphql"
)

var client *graphql.Client
var variables map[string]interface{}

func getAllTickets() {
	err := client.Query(context.Background(), &TicketQuery, variables)
	if err != nil {
		fmt.Println("Error parsing tickets: ", err)
	} else {
		t := TicketQuery.Tickets
		fmt.Println("Tickets retrieved: ", len(t.Tickets))
		for _, ticket := range t.Tickets {

			var tags []string
			for _, tag := range ticket.Tags {
				tags = append(tags, string(tag))
			}
			created := fmt.Sprintf("%s%s", strings.Replace(string(ticket.Createdat[0:19]), " ", "T", 1), "Z")
			createdAt, err := time.Parse(time.RFC3339, created)
			if err != nil {
				fmt.Println("Error converting createdat string to time", err)
			}
			var updated string
			var updatedAt time.Time
			if ticket.Updatedat != "" {
				updated = fmt.Sprintf("%s%s", strings.Replace(string(ticket.Updatedat[0:19]), " ", "T", 1), "Z")
				updatedAt, err = time.Parse(time.RFC3339, updated)
				if err != nil {
					fmt.Println("Error converting updatedat string to time", err)
				}
			}

			activeTickets = append(activeTickets, Ticket{
				ID:        int(ticket.ID),
				SLA:       string(ticket.SLA),
				Tags:      tags,
				Status:    string(ticket.Status),
				Subject:   string(ticket.Subject),
				UpdatedAt: updatedAt,
				CreatedAt: createdAt,
			})

		}
	}
}
