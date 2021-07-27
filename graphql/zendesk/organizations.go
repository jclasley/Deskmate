package zendesk

import (
	"context"

	"github.com/tylerconlee/Deskmate/graphql/model"
)

// GetOrganization takes the client, c, and requests the details for the
// organization provided by the context to the Zendesk API wrapper. Once it
// retreives that data from Zendesk, it converts the output into a model.
// Organization.
func (c *Client) GetOrganization(ctx context.Context, id int) (output *model.Organization, err error) {
	o, err := c.client.GetOrganization(ctx, int64(id))
	output = &model.Organization{
		URL:         o.URL,
		ID:          int(o.ID),
		Name:        o.Name,
		CreatedAt:   o.CreatedAt.String(),
		UpdatedAt:   o.UpdatedAt.String(),
		DomainNames: o.DomainNames,
		Tags:        o.Tags,
	}
	return
}
