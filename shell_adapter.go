package gilibot

import (
	"bufio"
	"os"
	"strings"
)

type shellAdapter struct {
	bot     *Bot
	history []string
}

func NewShellAdapter(b *Bot) *shellAdapter {
	return &shellAdapter{bot: b}
}

func (s *shellAdapter) Start() error {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		os.Stdout.WriteString(s.bot.Name + " > ")
		scanner.Scan()

		line := scanner.Text()
		line = strings.TrimSpace(line)

		//add line to history
		s.history = append(s.history, line)

		if line == "quit" || line == "q" || line == "exit" {
			os.Stdout.WriteString("GoodBye!")
			return nil
		}

		//TODO
		if line == "help" {
			//display help
		}

		if line == "history" || line == "h" || line == "hist" {
			for _, line := range s.history {
				os.Stdout.WriteString(line)
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

func (s *shellAdapter) Reply(e *Envelope, messages []string) {

	for m, _ := range messages {
		os.Stdout.WriteString(m)
	}
}
