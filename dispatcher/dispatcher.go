package dispatcher

import (
	"alina/definitions"
	"alina/objects"
)

func New() definitions.Dispatcher {
	return &dispatcher{}
}

type dispatcher struct {
	messageHandlers []func(message definitions.PrivateMessage, err error)
}

func (d *dispatcher) Handle(update definitions.UpdateBody) {
	switch update.GetType() {
	case definitions.MessageNew:
		d.handleMessage(update)
	}
}

func (d *dispatcher) AddMessageHandler(handler func(message definitions.PrivateMessage, err error)) {
	if d.messageHandlers == nil {
		d.messageHandlers = make([]func(message definitions.PrivateMessage, err error), 0)
	}
	d.messageHandlers = append(d.messageHandlers, handler)
}

func (d *dispatcher) handleMessage(update definitions.UpdateBody) {
	msg, err := objects.NewPrivateMessageFromUpdate(update)
	for _, v := range d.messageHandlers {
		v(msg, err)
	}
}
