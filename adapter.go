package gilibot

type AdapterInterface interface {
	Start() error
	Reply(*Envelope, []string)
}
