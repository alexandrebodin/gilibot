package main

import (
	"github.com/alexandrebodin/gilibot"
	"log"
	"time"
)

func main() {

	bot := gilibot.New("slack")

	bot.ListenFunc("!time", func(c *gilibot.Context) {
		layout := "Mon Jan 2 2006 15:04:05"
		t := time.Now()
		c.Reply([]string{t.Format(layout)})
	})

	err := bot.Start()
	if err != nil {
		log.Fatal(err)
	}
}
