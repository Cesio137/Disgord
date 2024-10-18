package events

type EventHandler interface {
	OnTeste(string)
	OnInt(int)
}

var Events[]interface{}