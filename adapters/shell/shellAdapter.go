package shell

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/alexandrebodin/gilibot"
)

// A simple history structure
type shellHistory struct {
	list         []string
	currentIndex int
	maxLen       int
}

// Adapter defines an bot adapter to send command from the shell
type Adapter struct {
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

// New returns a new shellAdapter
func New(b *gilibot.Bot) *Adapter {

	maxLen := 10000
	return &Adapter{
		bot: b,
		history: &shellHistory{
			list:         make([]string, maxLen),
			currentIndex: 0,
			maxLen:       maxLen,
		},
	}
}

// Start starts the adapter
func (s *Adapter) Start() error {

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
				fmt.Fprintf(os.Stdout, "%4s %v\n", strconv.Itoa(index+1), line)
			}

			continue
		}

		v := gilibot.Message{Text: line}
		s.bot.ReceiveMessage(v)
	}
}

// Reply sends back an answer through the adapter
func (s *Adapter) Reply(msg gilibot.Message, message string) error {

	os.Stdout.WriteString(message + "\n")
	return nil
}
