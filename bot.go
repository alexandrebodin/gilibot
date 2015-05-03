package gilibot

import (
	"fmt"
)

type Bot struct {
	Name      string
	adapter   AdapterInterface
	listeners []ListenerInterface
	matcher   *Matcher
}

const (
	defaultBotName     = "GiliBot"
	defaultAdapterName = "shell"
)

func New() *Bot {
	return &Bot{
		Name:      defaultBotName,
		listeners: []ListenerInterface{},
		matcher:   &Matcher{},
	}
}

func (b *Bot) SetAdapter(a AdapterInterface) {
	b.adapter = a
}

func (b *Bot) Start() error {

	if b.adapter == nil {
		return fmt.Errorf("no adapter set")
	}

	return b.adapter.Start()
}

func (b *Bot) ListenFunc(regex string, handler ListenerFunc) {
	b.matcher.AddHandler(&ListenerHandler{regex, handler})
}

func (b *Bot) RegisterListener(l ListenerInterface) {
	handlers := l.GetHandlers()
	for _, handler := range handlers {
		b.matcher.AddHandler(handler)
	}
}

func (b *Bot) ReceiveMessage(message string) {
	b.matcher.HandleMessage(message)
}
