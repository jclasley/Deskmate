package resolver

import (
	"context"

	"github.com/tylerconlee/Deskmate/zendesk/model"
	"github.com/tylerconlee/Deskmate/zendesk/zendesk"
)

var c *zendesk.Client

// ***** GET ticket functions ***** //
// GetAllTickets takes the ZendeskConfig object of username, APIkey and URL and
// makes a request to Zendesk to the /tickets.json endpoint. This returns all
// tickets in the Tickets type, found in the schema.
// Endpoint: /tickets.json
func (r *queryResolver) GetAllTickets(ctx context.Context, user string, apikey string, url string) (*model.Tickets, error) {
	config := model.ZendeskConfigInput{
		User:   user,
		Apikey: apikey,
		URL:    url,
	}
	c = zendesk.Connect(&config)
	output, err := c.GetTickets(ctx)
	if err != nil {
		return nil, err
	}
	tickets := &model.Tickets{
		Tickets: output,
		Count:   len(output),
	}

	return tickets, nil
}

// ***** GET organization functions ***** //
// GetOrganization takes the ZendeskConfig object of username, APIkey and URL,
// as well as an organization ID and makes a request to Zendesk to the /
// organization.json endpoint. This returns the information related to that
// organization.
// Endpoint: /organization.json
func (r *queryResolver) GetOrganization(ctx context.Context, config model.ZendeskConfigInput, id int) (*model.Organization, error) {
	c = zendesk.Connect(&config)
	org, err := c.GetOrganization(ctx, id)
	if err != nil {
		return nil, err
	}

	return org, nil
}

// ***** GET user functions ***** //
//

func (r *queryResolver) GetUser(ctx context.Context, user string, apikey string, url string, id string) (*model.User, error) {
	config := model.ZendeskConfigInput{
		User:   user,
		Apikey: apikey,
		URL:    url,
	}
	c = zendesk.Connect(&config)
	org, err := c.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return org, nil
}
