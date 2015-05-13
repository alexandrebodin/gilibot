package gilibot

type ListenerInterface interface {
	GetHandlers() []*ListenerHandler
}

type ListenerHandler struct {
	Regex       string
	HandlerFunc ListenerFunc
}

type ListenerFunc func(c *Context)
