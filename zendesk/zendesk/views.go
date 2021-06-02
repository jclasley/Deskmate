package zendesk

import (
	"context"

	"github.com/tylerconlee/Deskmate/zendesk/model"
	"go.uber.org/zap"
)

// GetViews takes the client, c, and requests a list of all
// of the currently active views in Zendesk
// View.
func (c *Client) GetViews(ctx context.Context) ([]*model.View, error) {
	var output []*model.View
	o, _, err := c.client.GetViews(ctx)
	if err != nil {
		log.Error("Error retrieving views", zap.String("Error", err.Error()))
		return nil, err
	}

	for _, v := range o {
		output = append(output, &model.View{
			ID:        int(v.ID),
			Title:     v.Title,
			CreatedAt: v.CreatedAt.String(),
			UpdatedAt: v.UpdatedAt.String(),
		})
	}
	return output, nil
}

// GetView takes the client, c, and a View ID, and pulls the details for a
// specific view
func (c *Client) GetView(ctx context.Context, viewID int) (view *model.View, err error) {
	var output *model.View
	o, err := c.client.GetView(ctx, int64(viewID))
	output = &model.View{
		ID:        int(o.ID),
		Title:     o.Title,
		CreatedAt: o.CreatedAt.String(),
		UpdatedAt: o.UpdatedAt.String(),
	}
	return output, nil
}

// GetViewCount takes the client, c, and a View ID, and retrieves the count of
// tickets that match the criteria of that view
func (c *Client) GetViewCount(ctx context.Context, viewID int) (viewCount *model.ViewCount, err error) {
	var output *model.ViewCount
	o, err := c.client.GetViewCount(ctx, int64(viewID))
	output = &model.ViewCount{
		ViewID: int(o.ViewID),
		Value:  int(o.Value),
		Pretty: o.Pretty,
	}
	return output, nil
}
