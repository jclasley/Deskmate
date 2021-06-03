package config

type Config struct {
	Slack   Slack
	Zendesk Zendesk
}

type Slack struct {
	SlackURL     string `json:"slackurl,omitempty"`
	SlackAPI     string `json:"slackapi,omitempty"`
	SlackSigning string `json:"slacksigning,omitempty"`
}

type Zendesk struct {
	ZendeskUser string
	ZendeskAPI  string
	ZendeskURL  string
}
