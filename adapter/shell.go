package adapter

import (
	"bufio"
	"fmt"
	"github.com/alexandrebodin/gilibot"
	"os"
	"strings"
)

type shellAdapter struct {
	bot     *gilibot.Bot
	history []string
}

func NewShellAdapter(b *gilibot.Bot) *shellAdapter {
	return &shellAdapter{bot: b}
}

func (s *shellAdapter) Start() error {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Fprint(os.Stdout, s.bot.Name+" > ")
		scanner.Scan()

		line := scanner.Text()
		line = strings.TrimSpace(line)

		//add line to history
		s.history = append(s.history, line)

		if line == "quit" || line == "q" || line == "exit" {
			fmt.Fprintln(os.Stdout, "GoodBye!")
			return nil
		}

		//TODO
		if line == "help" {
			//display help
		}

		if line == "history" || line == "h" || line == "hist" {
			for _, line := range s.history {
				fmt.Fprintln(os.Stdout, line)
			}
			continue
		}

		s.bot.ReceiveMessage(line)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
