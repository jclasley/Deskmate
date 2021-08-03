package slack

import (
	"github.com/slack-go/slack"
	"github.com/circleci/Deskmate/server/config"
	l "github.com/circleci/Deskmate/server/log"
)

var (
	api    *slack.Client
	c      *config.Config
	status bool

	log = l.Log
)

// LoadConfig is called by the Connect() function and requests
// the LoadConfig function from the config package. It sets the
// loaded configuration to the package-wide variable 'c'
func LoadConfig() {
	c = config.LoadConfig()
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
		log.Error("Slack Disconnected. Unable to retreive Slack info.")
		return status
	}
	status = true
	log.Info("Connected to Slack. Team: ", team.Name)
	return status

}
