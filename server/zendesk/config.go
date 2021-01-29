package zendesk

import (
	"fmt"

	"github.com/shurcooL/graphql"
	con "github.com/tylerconlee/Deskmate/server/config"
)

type ZendeskConfigInput struct {
	user   graphql.String
	apikey graphql.String
	url    graphql.String
}

var c ZendeskConfigInput

func SetConfig() {
	global := con.LoadConfig()
	c = ZendeskConfigInput{
		user:   graphql.String(global.Zendesk.ZendeskUser),
		apikey: graphql.String(global.Zendesk.ZendeskAPI),
		url:    graphql.String(global.Zendesk.ZendeskURL),
	}
	fmt.Println("config: ", c)
}
