package resolver

import (
	"context"

	"github.com/circleci/Deskmate/graphql/model"
	"github.com/circleci/Deskmate/graphql/zendesk"
)

// *** GET View functions *** //
// GetView retreives the specified view from Zendesk, and uses the
// ZendeskConfig object to authenticate the request.
// Endpoint: /views/{id}.json
func (r *queryResolver) GetView(ctx context.Context, config model.ZendeskConfigInput, viewID int) (*model.View, error) {
	return nil, nil
}

func (r *queryResolver) GetAllViews(ctx context.Context, config model.ZendeskConfigInput) (*model.Views, error) {
	c = zendesk.Connect(&config)
	output, err := c.GetViews(ctx)
	if err != nil {
		return nil, err
	}
	views := &model.Views{
		Views: output,
		Count: len(output),
	}

	return views, nil
}

func (r *queryResolver) GetViewCount(ctx context.Context, config model.ZendeskConfigInput, viewID int) (*model.ViewCount, error) {
	c = zendesk.Connect(&config)
	count, err := c.GetViewCount(ctx, viewID)
	if err != nil {
		return nil, err
	}

	return count, nil
}
