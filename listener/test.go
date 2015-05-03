package listener

import (
	"fmt"
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
				fmt.Println("matched \"mon test\"")
			},
		},
		{
			Regex: "merci (.*)",
			HandlerFunc: func(c *gilibot.Context) {
				fmt.Println(c.Matches[1])
			},
		},
	}
}
