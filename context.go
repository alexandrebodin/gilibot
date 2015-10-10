package gilibot

// Context passed to handlers
type Context struct {
	Matches []string
	Message Message
	Bot     *Bot
}

// Reply through the bot to the adapter
func (c *Context) Reply(message string) {
	c.Bot.Reply(c.Message, message)
}
