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

type Context struct {
	Matches []string
	Message Message
	Bot     *Bot
}

func (c *Context) Reply(message string) {
	c.Bot.Reply(c.Message, message)
}

// type MessageInterface interface {
// 	Text() string
// 	Channel() string
// 	User() string
// }

type Message struct {
	Channel string
	User    string
	Text    string
}
//
// func (m *Message) Text() string {
// 	return m.text
// }
//
// func (m *Message) User() string {
// 	return m.user
// }
//
// func (m *Message) Channel() string {
// 	return m.channel
// }
