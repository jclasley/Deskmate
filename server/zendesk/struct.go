package zendesk

import "time"

type Ticket struct {
	ID             int
	SLA            string
	Tags           []string
	UpdatedAt      time.Time
	CreatedAt      time.Time
	Status         string
	Subject        string
	AdditionReason int
}
