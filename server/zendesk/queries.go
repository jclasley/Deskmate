package zendesk

import "github.com/shurcooL/graphql"

var TicketQuery struct {
	Tickets struct {
		Tickets []struct {
			ID         graphql.Int
			SLA        graphql.String
			GroupID    graphql.String
			Assigneeid graphql.String
			Tags       []graphql.String
			Updatedat  graphql.String
			Createdat  graphql.String
			Subject    graphql.String
			Status     graphql.String
		}
	} `graphql:"getAllTickets(user: $user, apikey: $apikey, url: $url)"`
}

var AssigneeQuery struct {
	User struct {
		ID    graphql.String
		Name  graphql.String
		Email graphql.String
	} `graphql:"getUser(user: $user, apikey: $apikey, url: $url, id: $id)"`
}
