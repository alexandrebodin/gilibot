package shell

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
    "github.com/alexandrebodin/gilibot"
)

type shellHistory struct {
	list []string
	currentIndex int
	maxLen int
}

type shellAdapter struct {
	bot     *gilibot.Bot
	history *shellHistory
}

func (h *shellHistory) Add(line string) {
	if h.currentIndex >= h.maxLen {
		h.list = append(h.list[1:h.maxLen], line)
	} else {
		h.list[h.currentIndex] = line
		h.currentIndex++
	}
}

func New(b *gilibot.Bot) *shellAdapter {

	maxLen := 10000
	return &shellAdapter{
		bot: b,
		history: &shellHistory{
			list: make([]string, maxLen),
			currentIndex: 0,
			maxLen: maxLen,
		},
	}
}

func (s *shellAdapter) Start() error {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		os.Stdout.WriteString(s.bot.Name + " > ")
		scanner.Scan()

		line := scanner.Text()
		line = strings.TrimSpace(line)

		//add line to history
		s.history.Add(line)

		if line == "quit" || line == "q" || line == "exit" {
			os.Stdout.WriteString("GoodBye!")
			return nil
		}

		//TODO
		if line == "help" {
			//display help
		}

		if line == "history" || line == "h" || line == "hist" {

			for index, line := range s.history.list {
				fmt.Fprintf(os.Stdout, "%4s %v\n", strconv.Itoa(index + 1), line)
			}

			continue
		}

		v := gilibot.Message{Text: line}
		s.bot.ReceiveMessage(v)
	}
}

func (s *shellAdapter) Reply(msg gilibot.Message, message string) error {

	os.Stdout.WriteString(message + "\n")
	return nil
}
