package gilibot

type Matcher struct {
	Bot      *Bot
	handlers []*Listener
}

func (m *Matcher) AddHandler(l *Listener) {
	m.handlers = append(m.handlers, l)
}

func (m *Matcher) HandleMessage(message MessageInterface) {

	for _, h := range m.handlers {
		matches := h.Regex.FindStringSubmatch(message.Text())

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
	Message MessageInterface
	Bot     *Bot
}

func (c *Context) Reply(messages []string) {
	c.Bot.Reply(c.Message, messages)
}

type MessageInterface interface {
	Text() string
	Channel() string
	User() string
}

type Message struct {
	channel string
	user    string
	text    string
}

func (m *Message) Text() string {
	return m.text
}

func (m *Message) User() string {
	return m.user
}

func (m *Message) Channel() string {
	return m.channel
}
