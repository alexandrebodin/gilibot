package gilibot

import (
	slack "github.com/alexandrebodin/slack_rtm"
	"log"
	"os"
)

type slackAdapter struct {
	bot    *Bot
	client *slack.SlackClient
}

func NewSlackAdapter(b *Bot) *slackAdapter {
	return &slackAdapter{bot: b}
}

type slackHandler struct {
	bot *Bot
}

func (h *slackHandler) OnMessage(c *slack.SlackContext, m *slack.MessageType) error {

	msg := &Message{
		channel: m.Channel,
		user:    m.User,
		text:    m.Text,
	}
	h.bot.ReceiveMessage(msg)
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

	s.client = slackClient

	h := &slackHandler{bot: s.bot}
	slackClient.AddListener(slack.MessageEvent, h)

	err = slackClient.Run()
	if err != nil {
		return err
	}

	return nil
}

func (s *slackAdapter) Reply(msg MessageInterface, message string) error {

	resp := slack.ResponseMessage{
		Id:      "1",
		Type:    "message",
		Text:    ">>>" + message,
		Channel: msg.Channel(),
	}

	err := s.client.WriteMessage(resp)
	if err != nil {
		return err
	}

	return nil
}
