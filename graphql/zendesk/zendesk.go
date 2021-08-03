package zendesk

import (
	"net/http"
	"time"

	logger "github.com/tylerconlee/Deskmate/graphql/log"
	"github.com/tylerconlee/Deskmate/graphql/model"
	"github.com/tylerconlee/zendesk-go/zendesk"
	"go.uber.org/zap"
)

var (
	log = logger.Log
)

// Client contains the instance of the Zendesk API wrapper client from
// go-zendesk/zendesk, as well as the configuration for connecting to
// Zendesk.
type Client struct {
	config *model.ZendeskConfigInput
	client *zendesk.Client
}

// Connect takes a configuration for Zendesk and uses it to create a new
// instance of a Zendesk client. c is then used as a receiver for any other
// interaction with Zendesk, as it's preconfigured, and contains an http.Client
func Connect(config *model.ZendeskConfigInput) *Client {

	var err error
	// c is the instance of a Client that gets used for all functions.
	c := &Client{config: config}
	client := http.Client{
		Timeout: time.Second * 480,
		Transport: &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
		},
	}
	c.client, err = zendesk.NewClient(&client)
	if err != nil {
		log.Error("error creating zendesk client", zap.String("Error", err.Error()))
	}
	c.client.SetSubdomain(config.URL)
	c.client.SetCredential(zendesk.NewAPITokenCredential(config.User, config.Apikey))
	//log.Info("Zendesk credentials set. Client successfully created", zap.String("subdomain", config.URL))
	return c
}
