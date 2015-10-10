package listeners

import (
	"time"

	"github.com/alexandrebodin/gilibot"
)

// UtilsListener defines a structure to hold a list of utility listeners
type UtilsListener struct{}

// GetHandlers Returns utility ListenerHandler
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
