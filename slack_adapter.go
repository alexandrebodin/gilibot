package gilibot

import (
	slack "github.com/alexandrebodin/slack_rtm"
)

type slackAdapter struct {
	slackClient *slack.SlackClient
}

func (s *slackAdapter) Start() error {

}
