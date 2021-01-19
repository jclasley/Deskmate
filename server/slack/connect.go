package slack

import (
	"fmt"

	"github.com/slack-go/slack"
	"github.com/tylerconlee/Deskmate/server/config"
)

var (
	api    = slack.New(c.Slack.SlackAPI)
	c      config.Config
	status bool
)

// LoadConfig is called by the Connect() function and requests
// the LoadConfig function from the config package. It sets the
// loaded configuration to the package-wide variable 'c'
func LoadConfig() {
	c = config.LoadConfig()
	api = slack.New(c.Slack.SlackAPI)
}

// Connect loads the configuration needed to connect to a Slack instance,
// and then uses the OAuth Bot API key for Slack to establish a connection.
// TODO: Add in a catch for if the connection is unable to be established.
func Connect() {
	LoadConfig()
	api = slack.New(c.Slack.SlackAPI)
}

// Ping checks to see if there's a valid connection to a Slack instance by
// requesting the Team information from Slack and returning a boolean value.
// If TRUE, it logs the connection as successful, and outputs the connected
// Slack's team name.
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
