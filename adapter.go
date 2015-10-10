package gilibot

// AdapterInterface is implemented by bot adapters
type AdapterInterface interface {
	Start() error
	Reply(Message, string) error
}
