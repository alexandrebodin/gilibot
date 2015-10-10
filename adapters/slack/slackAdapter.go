package slack

import (
	"log"
	"os"
	"strings"

	"github.com/alexandrebodin/gilibot"
	slack "github.com/alexandrebodin/slack_rtm"
)

// Adapter defines a bot adpater to receive data from slack chat
type Adapter struct {
	bot    *gilibot.Bot
	client *slack.SlackClient
}

// New returns a new slack adapter
func New(b *gilibot.Bot) *Adapter {
	return &Adapter{bot: b}
}

type slackHandler struct {
	bot *gilibot.Bot
}

func (h *slackHandler) OnMessage(c *slack.SlackContext, m *slack.MessageType) error {

	msg := gilibot.Message{
		Channel: m.Channel,
		User:    m.User,
		Text:    strings.Replace(m.Text, "&amp;", "&", -1),
	}
	h.bot.ReceiveMessage(msg)
	return nil
}

// Start starts slack adapter
func (s *Adapter) Start() error {

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

// Reply send back and answer to slack
func (s *Adapter) Reply(msg gilibot.Message, message string) error {

	resp := slack.ResponseMessage{
		Id:      "1",
		Type:    "message",
		Text:    ">>>" + message,
		Channel: msg.Channel,
	}

	err := s.client.WriteMessage(resp)
	if err != nil {
		return err
	}

	return nil
}
