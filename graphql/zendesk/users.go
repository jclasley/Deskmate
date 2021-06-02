package zendesk

import (
	"context"
	"strconv"

	"github.com/tylerconlee/Deskmate/zendesk/model"
)

// GetUser takes the client, c, and requests the details for the
// user provided by the context to the Zendesk API wrapper. Once it
// retreives that data from Zendesk, it converts the output into a model.
// user.
func (c *Client) GetUser(ctx context.Context, id string) (output *model.User, err error) {
	userID, err := strconv.Atoi(id)
	o, err := c.client.GetUser(ctx, int64(userID))
	user := strconv.FormatInt(o.ID, 10)
	defaultGroup := strconv.FormatInt(o.DefaultGroupID, 10)
	output = &model.User{
		Active:       o.Active,
		Defaultgroup: defaultGroup,
		ID:           user,
		Email:        o.Email,
		Name:         o.Name,
		Createdat:    o.CreatedAt.String(),
		Updatedat:    o.UpdatedAt.String(),
		Lastlogin:    o.LastLoginAt.String(),
		Timezone:     o.Timezone,
	}
	return
}
