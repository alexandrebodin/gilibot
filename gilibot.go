// Package gilibot implements a simple library to receive messages from adapters and match them to custom listeners
package gilibot

import (
	"errors"
	"regexp"
)

// Bot controls adapter and listeners
type Bot struct {
	Name      string
	Adapter   AdapterInterface
	listeners []ListenerInterface
	matcher   *Matcher
}

var (
	errNoAdapterSet = errors.New("You must add at least one adapter")
)

const (
	defaultBotName = "GiliBot"
)

// New returns a bot instance
func New(arguments ...string) *Bot {

	b := new(Bot)
	b.listeners = []ListenerInterface{}
	b.matcher = &Matcher{Bot: b}

	if len(arguments) > 0 {
		b.Name = arguments[0]
	} else {
		b.Name = defaultBotName
	}

	return b
}

// AddAdapter adds an adapter to the bot
func (b *Bot) AddAdapter(a AdapterInterface) {
	b.Adapter = a
}

// Start starts the bot
func (b *Bot) Start() error {

	if b.Adapter == nil {
		return errNoAdapterSet
	}

	return b.Adapter.Start()
}

// ListenFunc adds a custom ListenerFunc to the bot
func (b *Bot) ListenFunc(regex string, handler ListenerFunc) {

	regexp := regexp.MustCompile(regex)
	b.matcher.AddHandler(&Listener{regexp, handler})
}

// RegisterListener adds a list of ListenFunc to the bot
func (b *Bot) RegisterListener(l ListenerInterface) {
	handlers := l.GetHandlers()
	for _, handler := range handlers {
		regexp := regexp.MustCompile(handler.Regex)
		b.matcher.AddHandler(&Listener{regexp, handler.HandlerFunc})
	}
}

// ReceiveMessage transfers a message received by an adapter to the listeners
func (b *Bot) ReceiveMessage(message Message) {
	b.matcher.HandleMessage(message)
}

// Reply send and answer through an adapter
func (b *Bot) Reply(m Message, message string) {
	b.Adapter.Reply(m, message)
}
