package gilibot

import "regexp"

type Matcher struct {
	handlers []*ListenerHandler
}

func (m *Matcher) AddHandler(l *ListenerHandler) {
	m.handlers = append(m.handlers, l)
}

func (m *Matcher) HandleMessage(message string) {
	for _, h := range m.handlers {

		re := regexp.MustCompile(h.Regex)
		matches := re.FindStringSubmatch(message)

		if len(matches) > 0 {
			c := &Context{matches}
			h.HandlerFunc(c)
		}
	}
}
