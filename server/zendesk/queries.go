package zendesk

import "github.com/shurcooL/graphql"

var TicketQuery struct {
	Tickets struct {
		Tickets []struct {
			ID        graphql.Int
			SLA       graphql.String
			Tags      []graphql.String
			Updatedat graphql.String
			Createdat graphql.String
			Subject   graphql.String
			Status    graphql.String
		}
	} `graphql:"getAllTickets(user: $user, apikey: $apikey, url: $url)"`
}
