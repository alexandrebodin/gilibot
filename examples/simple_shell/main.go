package main

import (
	"github.com/alexandrebodin/gilibot"
	listeners "github.com/alexandrebodin/gilibot/listeners"
	"log"
)

func main() {

	bot := gilibot.New("slack")

	j := listeners.NewJenkinsListener("http://jenkins.kilix.net/", "abodin", "2dafc5494f3df8e50317ecabbe15f936")
	bot.RegisterListener(j)
	bot.RegisterListener(&listeners.UtilsListener{})

	err := bot.Start()
	if err != nil {
		log.Fatal(err)
	}
}
