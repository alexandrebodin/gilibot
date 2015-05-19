package main

import (
	"github.com/alexandrebodin/gilibot"
	listeners "github.com/alexandrebodin/gilibot/listeners"
	"log"
	"time"
)

func main() {

	bot := gilibot.New("slack")

	bot.ListenFunc("!time", func(c *gilibot.Context) {
		layout := "Mon Jan 2 2006 15:04:05"
		t := time.Now()
		c.Reply([]string{"Time : \n" + t.Format(layout)})
	})

	j := listeners.NewJenkinsListener("http://jenkins.kilix.net/", "abodin", "2dafc5494f3df8e50317ecabbe15f936")
	bot.RegisterListener(j)

	err := bot.Start()
	if err != nil {
		log.Fatal(err)
	}
}
