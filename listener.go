package gilibot

import "regexp"

// ListenerInterface allows a struct to returns a list of handlers to add to the bot
type ListenerInterface interface {
	GetHandlers() []*ListenerHandler
}

// ListenerHandler links matching command to a HandlerFunc
type ListenerHandler struct {
	Regex       string
	HandlerFunc ListenerFunc
}

// Listener is a Listener with a compiled regexp
type Listener struct {
	Regex       *regexp.Regexp
	HandlerFunc ListenerFunc
}

// ListenerFunc is the function prototype for listener's handler function
type ListenerFunc func(c *Context)
