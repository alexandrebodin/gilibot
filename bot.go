package gilibot

import (
	"errors"
	"regexp"
)

type Bot struct {
	Name      string
	Adapter   AdapterInterface
	listeners []ListenerInterface
	matcher   *Matcher
}

var (
	errNoAdapterSet   = errors.New("You must add at least one adapter")
)

const (
	defaultBotName    = "GiliBot"
)

func New(arguments ...string) *Bot {

	b := new(Bot)
	b.listeners = []ListenerInterface{}
	b.matcher =   &Matcher{Bot: b}

	if len(arguments) > 0 {
		b.Name = arguments[0]
	} else {
		b.Name = defaultBotName
	}

	return b
}

func (b *Bot) AddAdapter(a AdapterInterface) {
	b.Adapter = a
}

func (b *Bot) Start() error {

	if b.Adapter == nil {
		return errNoAdapterSet
	}

	return b.Adapter.Start()
}

func (b *Bot) ListenFunc(regex string, handler ListenerFunc) {

	regexp := regexp.MustCompile(regex)
	b.matcher.AddHandler(&Listener{regexp, handler})
}

func (b *Bot) RegisterListener(l ListenerInterface) {
	handlers := l.GetHandlers()
	for _, handler := range handlers {
		regexp := regexp.MustCompile(handler.Regex)
		b.matcher.AddHandler(&Listener{regexp, handler.HandlerFunc})
	}
}

func (b *Bot) ReceiveMessage(message Message) {
	b.matcher.HandleMessage(message)
}

func (b *Bot) Reply(m Message, message string) {

	b.Adapter.Reply(m, message)
}
