package config

type Config struct {
	Slack   Slack
	Zendesk Zendesk
}

type Slack struct {
	SlackURL NullString `json:"slackurl,omitempty"`
	SlackAPI NullString `json:"slackapi,omitempty"`
}

type Zendesk struct {
	ZendeskUser NullString
	ZendeskAPI  NullString
	ZendeskURL  NullString
}
