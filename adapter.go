package gilibot

type AdapterInterface interface {
	Start() error
	Reply(MessageInterface, string) error
}
