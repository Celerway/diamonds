package slapp

import (
	log "github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
)

func (a Slapp) PostDailyReviews() {

	channelID, timestamp, err := a.client.PostMessage(
		a.SlackChannelId,
		slack.MsgOptionText("Some text", false),
		slack.MsgOptionAsUser(false),
	)
	if err != nil {
		log.Errorf("%s\n", err)
		return
	}
	log.Infof("Message successfully sent to channel %s at %s", channelID, timestamp)

}
