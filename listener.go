package gilibot

type ListenerInterface interface {
	GetHandlers() []*ListenerHandler
}

type ListenerHandler struct {
	Regex       string
	HandlerFunc ListenerFunc
}

type Context struct {
	Matches []string
	Message string
	Adapter *AdapterInterface
}

type ListenerFunc func(c *Context)
