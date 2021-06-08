package config

type Config struct {
	Slack   Slack   `json:"slack"`
	Zendesk Zendesk `json:"zendesk"`
}

type Slack struct {
	SlackURL     string `json:"url,omitempty"`
	SlackAPI     string `json:"api,omitempty"`
	SlackSigning string `json:"sign,omitempty"`
}

type Zendesk struct {
	ZendeskUser string `json:"user"`
	ZendeskAPI  string `json:"api"`
	ZendeskURL  string `json:"url"`
}
