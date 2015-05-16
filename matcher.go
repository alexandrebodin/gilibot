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
		matches := h.Regex.FindStringSubmatch(message.getText())

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
	getText() string
	getChannel() string
	getUser() string
}

type Message struct {
	Channel string
	User    string
	Text    string
}

func (m *Message) getText() string {
	return m.Text
}

func (m *Message) getUser() string {
	return m.User
}

func (m *Message) getChannel() string {
	return m.Channel
}

type TestMessage struct {
	Message
}

func (m *TestMessage) getText() string {
	return m.Text
}

func (m *TestMessage) getUser() string {
	return m.User
}

func (m *TestMessage) getChannel() string {
	return m.Channel
}
