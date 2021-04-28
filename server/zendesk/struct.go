package zendesk

import "time"

// Ticket is an individual instance of a ticket returned from the
// Zendesk GraphQL API (see /zendesk in this project). The fields
// listed here are just a subset of all available fields, as not all
// fields on the Ticket response are utilized by Deskmate
type Ticket struct {
	ID             int
	Assignee       int
	SLA            string
	Tags           []string
	UpdatedAt      time.Time
	CreatedAt      time.Time
	Status         string
	Subject        string
	AdditionReason int
}
