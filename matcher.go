package gilibot

type Matcher struct {
	Bot      *Bot
	handlers []*Listener
}

func (m *Matcher) AddHandler(l *Listener) {
	m.handlers = append(m.handlers, l)
}

func (m *Matcher) HandleMessage(message Message) {

	for _, h := range m.handlers {
		matches := h.Regex.FindStringSubmatch(message.Text)

		if len(matches) > 0 {
			c := &Context{
				Bot:     m.Bot,
				Matches: matches,
				Message: message,
			}
			h.HandlerFunc(c)
		}
	}
}
