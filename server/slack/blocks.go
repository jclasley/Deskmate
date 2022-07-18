package slack

import (
	"fmt"

	"github.com/slack-go/slack"
)

func SLANotification(notification map[string]interface{}) {
	user := "Unassigned"
	if notification["Email"] != "" {
		email := fmt.Sprintf("%s", notification["Email"])
		user = getUserID(email)
	}
	divSection := slack.NewDividerBlock()
	t, _ := ActiveTriage(notification["Channel"].(string))
	// Build Message with blocks created above
	var alertmsg string
	if t != "" {
		alertmsg = fmt.Sprintf("<@%s> - Upcoming SLA Alert on #%d - Less than %s remaining", t, notification["ID"], notification["TimeRemaining"])
	} else {
		alertmsg = fmt.Sprintf("<!here> - Upcoming SLA Alert on #%d - Less than %s remaining", notification["ID"], notification["TimeRemaining"])
	}
	// Header Section
	headerText := slack.NewTextBlockObject("mrkdwn", alertmsg, false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)

	// Main alert body
	ticketDetails := fmt.Sprintf("*<%s|%s>*\nTicket Created: %s\nTag: %s\nSLA Level: %s\nAssignee: <@%s>", notification["URL"], notification["Subject"], notification["CreatedAt"], notification["Tag"], notification["SLA"], user)

	scheduleText := slack.NewTextBlockObject("mrkdwn", ticketDetails, false, false)
	schedeuleSection := slack.NewSectionBlock(scheduleText, nil, nil)

	// Conflict Section
	conflictImage := slack.NewImageBlockElement("https://api.slack.com/img/blocks/bkb_template_images/notificationsWarningIcon.png", "notifications warning icon")

	expireText := fmt.Sprintf("*Ticket SLA expires at %s*", notification["SLA"])
	conflictText := slack.NewTextBlockObject("mrkdwn", expireText, false, false)

	conflictSection := slack.NewContextBlock(
		"",
		[]slack.MixedElement{conflictImage, conflictText}...,
	)

	// Approve and Deny Buttons
	acknowledgeBtnTxt := slack.NewTextBlockObject("plain_text", "Acknowledge", false, false)
	acknowledgeBtn := slack.NewButtonBlockElement("", "sla_ticket_ack", acknowledgeBtnTxt)
	acknowledgeSection := slack.NewActionBlock("", acknowledgeBtn)

	message := fmt.Sprintf("%s left on ticket #%v", notification["TimeRemaining"], notification["ID"])
	api.PostMessage(notification["Channel"].(string), slack.MsgOptionText(message, false), slack.MsgOptionBlocks(headerSection,
		divSection,
		schedeuleSection,
		conflictSection,
		divSection,
		acknowledgeSection))
}

func NewNotification(notification map[string]interface{}) {

	t, _ := ActiveTriage(notification["Channel"].(string))
	// Build Message with blocks created above
	var message string
	if t != "" {
		message = fmt.Sprintf("<@%s> - New ticket received - #%d", t, notification["ID"])
	} else {
		message = fmt.Sprintf("New ticket received - #%d", notification["ID"])
	}

	// Header Section
	url := fmt.Sprintf("*<%s|%s>*", notification["URL"], notification["Subject"])
	header := fmt.Sprintf("%s:\n%s", message, url)
	headerText := slack.NewTextBlockObject("mrkdwn", header, false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)

	// Fields
	ca := fmt.Sprintf("*Created At:*\n%s", notification["CreatedAt"])
	typeField := slack.NewTextBlockObject("mrkdwn", ca, false, false)

	fieldSlice := make([]*slack.TextBlockObject, 0)
	fieldSlice = append(fieldSlice, typeField)

	fieldsSection := slack.NewSectionBlock(nil, fieldSlice, nil)

	// Approve and Deny Buttons
	approveBtnTxt := slack.NewTextBlockObject("plain_text", "Acknowledge", false, false)
	approveBtn := slack.NewButtonBlockElement("", "new_ticket_ack", approveBtnTxt)

	actionBlock := slack.NewActionBlock("", approveBtn)

	api.PostMessage(notification["Channel"].(string), slack.MsgOptionText(message, false), slack.MsgOptionBlocks(headerSection,
		fieldsSection,
		actionBlock))

}

func UpdatedNotification(notification map[string]interface{}) {
	user := "Unassigned"
	if notification["Email"] != "" {
		email := fmt.Sprintf("%s", notification["Email"])
		user = getUserID(email)
	}
	t, _ := ActiveTriage(notification["Channel"].(string))
	// Build Message with blocks created above

	divSection := slack.NewDividerBlock()
	var message string
	if t != "" {
		message = fmt.Sprintf("<@%s> - Ticket update received - #%d", t, notification["ID"])
	} else {
		message = fmt.Sprintf("Ticket update received - #%d", notification["ID"])
	}

	// Header Section
	url := fmt.Sprintf("*<%s|%s>*", notification["URL"], notification["Subject"])
	header := fmt.Sprintf("%s:\n%s", message, url)
	headerText := slack.NewTextBlockObject("mrkdwn", header, false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)

	// Fields
	ca := fmt.Sprintf("*Created At:*\n%s\n*Updated At:*\n%s\n*Assignee*\n<@%s>\n", notification["CreatedAt"], notification["UpdatedAt"], user)
	typeField := slack.NewTextBlockObject("mrkdwn", ca, false, false)

	fieldSlice := make([]*slack.TextBlockObject, 0)
	fieldSlice = append(fieldSlice, typeField)

	fieldsSection := slack.NewSectionBlock(nil, fieldSlice, nil)

	// Approve and Deny Buttons
	approveBtnTxt := slack.NewTextBlockObject("plain_text", "Acknowledge", false, false)
	approveBtn := slack.NewButtonBlockElement("", "updated_ticket_ack", approveBtnTxt)

	actionBlock := slack.NewActionBlock("", approveBtn)

	api.PostMessage(notification["Channel"].(string), slack.MsgOptionText(message, false), slack.MsgOptionBlocks(headerSection, divSection,
		fieldsSection, divSection,
		actionBlock))

}
