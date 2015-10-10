package gilibot

type AdapterInterface interface {
	Start() error
	Reply(Message, string) error
}
