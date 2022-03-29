package slack

import (
	"fmt"
	"time"

	"github.com/slack-go/slack"
	"go.uber.org/ratelimit"
)

type Active struct {
	Channel Channel
	Enabled bool
}

var A []Active

func getChannelInfo(channel string) (info Channel) {
	c, err := api.GetConversationInfo(channel, false)
	if err != nil {
		fmt.Println("Error retrieving channel information")
	}
	info = Channel{
		Name: c.Name,
		ID:   c.ID,
	}
	return
}

func getUserInfo(user string) (info User) {
	u, err := api.GetUserInfo(user)
	if err != nil {
		fmt.Println("Error retrieving channel information")
	}
	info = User{
		Name: u.Name,
		ID:   u.ID,
	}
	return
}

func getUserID(email string) (info string) {
	u, err := api.GetUserByEmail(email)
	if err != nil {
		fmt.Println("Error retrieving user information", err.Error())
		return ""
	}
	info = u.ID
	return

}

func ListChannels() (channels []map[string]string) {
	params := slack.GetConversationsParameters{
		ExcludeArchived: "true",
		Limit:           1000,
		Types: []string{
			"public_channel",
		}}
	c, s, err := api.GetConversations(&params)
	rl := ratelimit.New(20, ratelimit.Per(time.Minute))
	for {
		rl.Take()
		if s != "" {
			var v []slack.Channel
			params = slack.GetConversationsParameters{Cursor: s,
				ExcludeArchived: "true",
				Limit:           1000,
				Types: []string{
					"public_channel",
				}}
			v, s, err = api.GetConversations(&params)
			c = append(c, v...)
		} else {
			break
		}
	}
	if err != nil {
		fmt.Println("Error retrieving channel list", err)
	}
	for _, channel := range c {
		channels = append(channels, map[string]string{
			"ChannelName": channel.Name,
			"ID":          channel.ID,
		})
	}
	return channels
}

func ListUsers() (users []map[string]string) {
	u, err := api.GetUsers()
	if err != nil {
		fmt.Println("Error retrieving user list")
	}
	for _, user := range u {
		users = append(users, map[string]string{
			"UserName": user.Name,
			"ID":       user.ID,
		})
	}
	return users
}

func toggleDeskmate(channel string) (active bool) {
	channelInfo := getChannelInfo(channel)
	current := deskmateActiveCheck(channel)
	enabled := !current

	// Loop through existing reminders and determine if they're
	// already set. If they are, update the value to the new
	// value.
	for _, active := range A {
		if channel == active.Channel.ID {
			active.Enabled = enabled
			return enabled
		}
	}
	// If no prior reminder was set, create a new entry
	a := Active{
		Channel: channelInfo,
		Enabled: enabled,
	}
	A = append(A, a)
	return enabled
}

func deskmateActiveCheck(channel string) (enabled bool) {
	for _, active := range A {
		if channel == active.Channel.ID {
			return active.Enabled
		}
	}
	return false
}
