package slack

import (
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
