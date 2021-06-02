package resolver

import (
	"context"

	"github.com/tylerconlee/Deskmate/zendesk/model"
	"github.com/tylerconlee/Deskmate/zendesk/zendesk"
)

// ***** GET Trigger functions ***** //
// GetAllTriggers takes the ZendeskConfig object of username, APIkey and URL and
// makes a request to Zendesk to the /triggers.json endpoint. This returns all
// triggers in the Triggers type, found in the schema.
// Endpoint: /triggers.json
func (r *queryResolver) GetAllTriggers(ctx context.Context, config model.ZendeskConfigInput) (*model.Triggers, error) {
	c = zendesk.Connect(&config)
	output, err := c.GetTriggers(ctx)
	if err != nil {
		return nil, err
	}
	triggers := &model.Triggers{
		Triggers: output,
		Count:    len(output),
	}

	return triggers, nil
}

// GetTrigger takes the ZendeskConfig object of username, APIkey and URL and
// makes a request to Zendesk to the /triggers/{id}.json endpoint. This returns
// details about a specific trigger instead of all available triggers.
// Endpoint: /triggers/{id}.json
func (r *queryResolver) GetTrigger(ctx context.Context, config model.ZendeskConfigInput, id int) (*model.Trigger, error) {
	c = zendesk.Connect(&config)
	trigger, err := c.GetTrigger(ctx, id)
	if err != nil {
		return nil, err
	}

	return trigger, nil
}
