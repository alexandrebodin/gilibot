package gilibot

type Response struct {
	Bot      *Bot
	Envelope *Envelope
}

type Envelope struct {
	Room string
	User string
}

func (r *Response) Reply(messages []string) {
	r.Bot.Reply(r.Envelope, messages)
}
