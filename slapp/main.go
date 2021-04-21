package slapp

import (
	"github.com/celerway/diamonds/service"
	log "github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
	"os"
)

func Ping() {

}
func getEnv() (string, string) {
	token, ok := os.LookupEnv("SLACK_TOKEN")
	if !ok {
		log.Fatal("Missing SLACK_TOKEN in environment")
	}
	channel, ok := os.LookupEnv("SLACK_CHANNEL")
	if !ok {
		log.Fatal("Missing SLACK_CHANNEL in environment")
	}

	return token, channel
}

func Initialize(service service.DiamondService) Slapp {
	s := Slapp{
		Service: service,
	}
	token, channel := getEnv()
	s.client = slack.New(token)
	s.SlackChannel = channel
	log.Info("Slack app initialized")
	return s
}
