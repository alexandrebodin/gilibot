package listener

import (
	"github.com/alexandrebodin/gilibot"
)

type testListener struct {
	bot *gilibot.Bot
}

func NewTestListener() *testListener {
	return &testListener{}
}

func (t *testListener) GetHandlers() []*gilibot.ListenerHandler {

	return []*gilibot.ListenerHandler{
		{
			Regex: "mon test",
			HandlerFunc: func(c *gilibot.Context) {
				c.Reply([]string{"matched \"mon test\""})
			},
		},
		{
			Regex: "merci (.*)",
			HandlerFunc: func(c *gilibot.Context) {
				c.Reply([]string{c.Matches[1]})
			},
		},
	}
}
