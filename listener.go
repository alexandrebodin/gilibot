package gilibot

import "regexp"

type ListenerInterface interface {
	GetHandlers() []*ListenerHandler
}

type ListenerHandler struct {
	Regex       string
	HandlerFunc ListenerFunc
}

type Listener struct {
	Regex       *regexp.Regexp
	HandlerFunc ListenerFunc
}

type ListenerFunc func(c *Context)
