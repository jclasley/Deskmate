package slack

import (
	"fmt"

	"github.com/nlopes/slack"
	"github.com/tylerconlee/Deskmate/server/config"
)

var (
	api    = slack.New(c.Slack.SlackAPI)
	c      config.Config
	status bool
)

func LoadConfig() {
	c = config.LoadConfig()
}

func Connect() {
	LoadConfig()
	api = slack.New(c.Slack.SlackAPI)
}

func Ping() bool {
	team, err := api.GetTeamInfo()
	if err != nil {
		status = false
		fmt.Println("Slack Disconnected. Unable to retreive Slack info.")
		return status
	}
	status = true
	fmt.Println("Connected to Slack. Team: ", team.Name)
	return status

}
