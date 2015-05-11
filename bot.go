package gilibot

import (
	"errors"
	"log"
)

type Bot struct {
	Name      string
	Adapter   AdapterInterface
	listeners []ListenerInterface
	matcher   *Matcher
}

var (
	errNoAdapterSet   = errors.New("No adapter set")
	errInvalidAdapter = errors.New("Invalid adapter name")
)

const (
	defaultBotName     = "GiliBot"
	defaultAdapterName = "shell"
)

func New(arguments ...string) *Bot {

	var adapterName string
	if len(arguments) > 0 {
		adapterName = arguments[0]
	} else {
		adapterName = defaultAdapterName
	}

	b := &Bot{
		Name:      defaultBotName,
		listeners: []ListenerInterface{},
		matcher:   &Matcher{},
	}

	err := b.initAdapter(adapterName)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func (b *Bot) initAdapter(adapterName string) error {
	switch adapterName {
	case "shell":
		b.Adapter = NewShellAdapter(b)
	case "slack":
		b.Adapter = NewSlackAdapter(b)
	default:
		return errInvalidAdapter
	}
	return nil
}

func (b *Bot) Start() error {

	if b.Adapter == nil {
		return errNoAdapterSet
	}

	return b.Adapter.Start()
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
