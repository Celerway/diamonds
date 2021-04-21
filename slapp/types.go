package slapp

import (
	"github.com/celerway/diamonds/service"
	"github.com/slack-go/slack"
)

type SlackMessage struct {
	payload string
}

type Slapp struct { // Slack App
	client         *slack.Client
	Service        service.DiamondService
	SlackChannel   string
	SlackChannelId string
}
