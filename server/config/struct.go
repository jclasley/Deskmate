package config

type Config struct {
	Slack   Slack
	Zendesk Zendesk
}

type Slack struct {
	SlackURL string `json:slackurl`
	SlackAPI string `json:slackapi`
}

type Zendesk struct {
	ZendeskUser string
	ZendeskAPI  string
	ZendeskURL  string
}
