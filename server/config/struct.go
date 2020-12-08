package config

type Config struct {
	Slack   Slack
	Zendesk Zendesk
}

type Slack struct {
	SlackURL string
	SlackAPI string
}

type Zendesk struct {
	ZendeskUser string
	ZendeskAPI  string
	ZendeskURL  string
}
