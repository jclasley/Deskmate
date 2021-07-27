package zendesk

import (
	"context"
	"fmt"

	"github.com/tylerconlee/Deskmate/graphql/model"
	"github.com/tylerconlee/zendesk-go/zendesk"
)

// GetTriggers uses the preconfigured client, c, and sends a request for all
// triggers to the Zendesk API wrapper. Once it grabs the array of triggers, it
// makes sure that any pagination is handled, and converts the ticket output
// into an array of model.Trigger.
func (c *Client) GetTriggers(ctx context.Context) (output []*model.Trigger, err error) {
	opts := zendesk.TriggerListOptions{}
	o, _, err := c.client.GetTriggers(ctx, &opts)

	for _, trigger := range o {
		any := []*model.TriggerCondition{}
		all := []*model.TriggerCondition{}
		actions := []*model.TriggerAction{}

		// Convert the zendesk.Trigger.Conditions.Any slice to a slice of model.
		// TriggerCondition
		for _, a := range trigger.Conditions.Any {
			s := &model.TriggerCondition{
				Field:    a.Field,
				Operator: a.Operator,
				Value:    a.Value,
			}
			any = append(any, s)
		}
		// Convert the zendesk.Trigger.Conditions.All slice to a slice of model.
		// TriggerCondition
		for _, a := range trigger.Conditions.All {
			s := &model.TriggerCondition{
				Field:    a.Field,
				Operator: a.Operator,
				Value:    a.Value,
			}
			all = append(any, s)
		}
		// Convert the zendesk.Trigger.Actions slice to a slice of model.
		// TriggerCondition
		for _, a := range trigger.Actions {
			s := &model.TriggerAction{
				Field: a.Field,
				Value: fmt.Sprintf("%v", a.Value),
			}
			actions = append(actions, s)
		}

		save := &model.Trigger{
			Title:       trigger.Title,
			Description: trigger.Description,
			ID:          int(trigger.ID),
			Position:    int(trigger.Position),
			CreatedAt:   trigger.CreatedAt.String(),
			UpdatedAt:   trigger.UpdatedAt.String(),
			Active:      trigger.Active,
			Conditions: &model.TriggerConditions{
				Any: any,
				All: all,
			},
			Actions: actions,
		}
		output = append(output, save)
	}

	return output, nil
}

// GetTrigger uses the preconfigured client, c, and sends a request for all
// triggers to the Zendesk API wrapper. Once it grabs the specified trigger, it
// converts the output into a type of model.Trigger.
func (c *Client) GetTrigger(ctx context.Context, id int) (output *model.Trigger, err error) {
	o, err := c.client.GetTrigger(ctx, int64(id))
	var any []*model.TriggerCondition
	var all []*model.TriggerCondition
	actions := []*model.TriggerAction{}

	// Convert the zendesk.Trigger.Conditions.Any slice to a slice of model.
	// TriggerCondition
	for _, a := range o.Conditions.Any {
		c := &model.TriggerCondition{
			Field:    a.Field,
			Operator: a.Operator,
			Value:    a.Value,
		}
		any = append(any, c)
	}
	// Convert the zendesk.Trigger.Conditions.All slice to a slice of model.
	// TriggerCondition
	for _, a := range o.Conditions.All {
		c := &model.TriggerCondition{
			Field:    a.Field,
			Operator: a.Operator,
			Value:    a.Value,
		}
		all = append(all, c)
	}

	conditions := &model.TriggerConditions{
		Any: any,
		All: all,
	}

	// Convert the zendesk.Trigger.Actions slice to a slice of model.
	// TriggerCondition
	for _, a := range o.Actions {
		s := &model.TriggerAction{
			Field: a.Field,
			Value: fmt.Sprintf("%v", a.Value),
		}
		actions = append(actions, s)
	}

	output = &model.Trigger{
		ID:          int(o.ID),
		Title:       o.Title,
		CreatedAt:   o.CreatedAt.String(),
		UpdatedAt:   o.UpdatedAt.String(),
		Position:    int(o.Position),
		Active:      o.Active,
		Description: o.Description,
		Conditions:  conditions,
		Actions:     actions,
	}
	return
}
