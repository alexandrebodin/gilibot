package main

import (
	"fmt"
	"github.com/alexandrebodin/gilibot"
	"github.com/alexandrebodin/gilibot/listener"
)

func main() {

	bot := gilibot.New("slack")

	bot.ListenFunc(".*", func(c *gilibot.Context) {
		c.Reply([]string{"coucou match everything"})
	})

	bot.RegisterListener(listener.NewTestListener())

	err := bot.Start()
	if err != nil {
		fmt.Println(err)
	}
}
