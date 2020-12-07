package main

import (
	"github.com/tylerconlee/Deskmate/datastore"
	"github.com/tylerconlee/Deskmate/server"
)

func main() {
	// On launch, Deskmate should connect to a local Postgres database using
	// environment variables for the Postgres connection details
	// It should then look for a configuration table. If no configuration table
	// is found, it should prompt the user at the command line to enter in the
	// Slack API authentication token.
	// Further configuration would then be handled within the Deskmate app in
	// Slack.
	datastore.ConnectPostgres()

	// Deskmate will use heavy use of the SlabAPI GraphQL API project. In order
	// to do so, however, the SlabAPI project will have to be revamped to not
	// launch its own webserver, but to serve as a general package for use in
	// other projects.
	// That will allow Deskmate to launch a webserver with the GraphQL API for
	// Zendesk.
	server.Launch()

	// TODO: Establish connection to Slack
	// With the webserver up and running, Deskmate will need an endpoint that
	// will be used for the Events API (https://api.slack.com/events-api). This
	// will likely be `/events`, which Slack will send POST data to. From there,
	// the `slack-go/slack` package will have the handler for that endpoint.

}
