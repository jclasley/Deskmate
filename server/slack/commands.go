package slack

import (
	"fmt"
	"regexp"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

var scripts []Script

type ScriptFunction func(*slackevents.AppMentionEvent)

type Script struct {
	Name               string
	Matcher            string
	Description        string
	CommandDescription string
	Function           ScriptFunction
}

func init() {
	RegisterScript(Script{
		Name:               "Help",
		Matcher:            "(?i)^help$",
		Description:        "show description for all commands",
		CommandDescription: "help",
		Function:           helpScriptFunc,
	})
	RegisterScript(Script{
		Name:               "Set Triage Role",
		Matcher:            "(?i)^set$",
		Description:        "sets the current user as the triage role for this channel",
		CommandDescription: "set",
		Function:           setTriageFunc,
	})
	RegisterScript(Script{
		Name:               "Unset Triage Role",
		Matcher:            "(?i)^unset$",
		Description:        "removes the current user from the triage role for this channel",
		CommandDescription: "unset",
		Function:           unsetTriageFunc,
	})

}

func RegisterScript(script Script) {
	scripts = append(scripts, script)
}

func HandleMentionEvent(event *slackevents.AppMentionEvent) {

	// Strip @bot-name out
	re, err := regexp.Compile(`^<@.*> *`)
	if err != nil {
		fmt.Println("error parsing command", err.Error())
	}
	event.Text = re.ReplaceAllString(event.Text, "")

	for _, script := range scripts {
		if match(script.Matcher, event.Text) {
			script.Function(event)
			return
		}
	}

	api.PostMessage(event.Channel, slack.MsgOptionText("Sorry, I don't know that command.", false))

}

func match(matcher string, content string) bool {
	re := regexp.MustCompile(matcher)
	return re.MatchString(content)
}

func helpScriptFunc(event *slackevents.AppMentionEvent) {
	helpMsg := "Prefix @deskmate to any command you would like to execute. \n\n"
	for i, script := range scripts {
		if i != 0 {
			helpMsg += "\n"
		}
		if script.CommandDescription != "" {
			helpMsg += "@deskmate " + script.CommandDescription
			if script.Description != "" {
				helpMsg += fmt.Sprintf(" - %s", script.Description)
			}
		} else {
			helpMsg += fmt.Sprintf("Missing help command description for %s", script.Name)
		}
	}
	api.PostMessage(event.Channel, slack.MsgOptionText(fmt.Sprintf("```%s```", helpMsg), false))
}

func setTriageFunc(event *slackevents.AppMentionEvent) {
	setTriage(event.Channel, event.User)
}

func unsetTriageFunc(event *slackevents.AppMentionEvent) {
	removeTriage(event.Channel)
}
