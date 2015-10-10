package gilibot

// Message is passed through the adapter to the listener
type Message struct {
	Channel string
	User    string
	Text    string
}
