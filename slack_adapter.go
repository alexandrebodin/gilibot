package gilibot

import (
	slack "github.com/alexandrebodin/slack_rtm"
	"log"
	"os"
)

type slackAdapter struct {
	bot *Bot
}

func NewSlackAdapter(b *Bot) *slackAdapter {
	return &slackAdapter{bot: b}
}

type slackHandler struct{}

func (h *slackHandler) OnMessage(c *slack.SlackContext, m *slack.MessageType) error {

	r := slack.ResponseMessage{
		Id:      "1",
		Type:    "message",
		Text:    "Coucou les copains",
		Channel: m.Channel,
	}
	c.Client.WriteMessage(r)
	return nil
}

func (s *slackAdapter) Start() error {

	token := os.Getenv("GILIBOT_SLACK_TOKEN")
	if token == "" {
		log.Fatal("slack token is missing")
	}

	slackClient, err := slack.New(token)
	if err != nil {
		return err
	}

	h := &slackHandler{}
	slackClient.AddListener(slack.MessageEvent, h)

	err = slackClient.Run()
	if err != nil {
		return err
	}

	return nil
}
