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

type slackHandler struct {
	bot *Bot
}

func (h *slackHandler) OnMessage(c *slack.SlackContext, m *slack.MessageType) error {

	h.bot.ReceiveMessage(m.Text)
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

	h := &slackHandler{bot: s.bot}
	slackClient.AddListener(slack.MessageEvent, h)

	err = slackClient.Run()
	if err != nil {
		return err
	}

	return nil
}
