package slapp

import (
	"github.com/celerway/diamonds/service"
	log "github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
	"os"
)

func (s Slapp) Say(payload string) {

	_, _, err := s.client.PostMessage(
		s.SlackChannelId,
		slack.MsgOptionText(payload, false),
	)
	if err != nil {
		log.Errorf("[slack] Saying '%s' provoked error: %v",
			payload, err)
		return
	}
	log.Debugf("[slack] Message posted to channel(%s)", s.SlackChannelId)
}

func getEnv() (string, string) {
	token, ok := os.LookupEnv("SLACK_TOKEN")
	if !ok {
		log.Fatal("Missing SLACK_TOKEN in environment")
	}
	channel, ok := os.LookupEnv("SLACK_CHANNELID")
	if !ok {
		log.Fatal("Missing SLACK_CHANNELID in environment")
	}

	return token, channel
}

func Initialize(service service.DiamondService) Slapp {
	s := Slapp{
		Service: service,
	}
	token, channelid := getEnv()
	s.client = slack.New(token)
	s.SlackChannelId = channelid
	log.Infof("Joining %s", channelid)
	_, _, _, err := s.client.JoinConversation(channelid)
	if err != nil {
		log.Fatalf("[slack] Could not join conversation: %v", err)
	}
	s.Say("Hello! I have 💎")
	log.Info("Slack app initialized")
	return s
}
