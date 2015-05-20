package listeners

import (
	"github.com/alexandrebodin/gilibot"
	"time"
)

type UtilsListener struct{}

func (u *UtilsListener) GetHandlers() []*gilibot.ListenerHandler {

	return []*gilibot.ListenerHandler{
		{
			Regex: "!time",
			HandlerFunc: func(c *gilibot.Context) {
				layout := "Mon Jan 2 2006 15:04:05"
				t := time.Now()
				c.Reply("Time : \n" + t.Format(layout))
			},
		},
	}
}
