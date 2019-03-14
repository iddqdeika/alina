package dispatcher

import (
	"alina/definitions"
)

func New(privateMessageFactory definitions.PrivateMessagesFactory) definitions.Dispatcher {
	return &dispatcher{privateMessageFactory: privateMessageFactory}
}

type dispatcher struct {
	privateMessageFactory definitions.PrivateMessagesFactory
	messageHandlers       []func(message definitions.PrivateMessage, err error)
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
	msg, err := d.privateMessageFactory.NewPrivateMessageFromUpdate(update)
	for _, v := range d.messageHandlers {
		v(msg, err)
	}
}
