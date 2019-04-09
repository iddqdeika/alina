package dispatcher

import (
	"alina/alina"
)

func New(privateMessageFactory alina.PrivateMessagesFactory) alina.Dispatcher {
	return &dispatcher{privateMessageFactory: privateMessageFactory}
}

type dispatcher struct {
	privateMessageFactory alina.PrivateMessagesFactory
	messageHandlers       []func(message alina.PrivateMessage, err error)
}

func (d *dispatcher) Handle(update alina.UpdateBody) {
	switch update.GetType() {
	case alina.MessageNew:
		d.handleMessage(update)
	}
}

func (d *dispatcher) AddMessageHandler(handler func(message alina.PrivateMessage, err error)) {
	if d.messageHandlers == nil {
		d.messageHandlers = make([]func(message alina.PrivateMessage, err error), 0)
	}
	d.messageHandlers = append(d.messageHandlers, handler)
}

func (d *dispatcher) handleMessage(update alina.UpdateBody) {
	msg, err := d.privateMessageFactory.NewPrivateMessageFromUpdate(update)
	for _, v := range d.messageHandlers {
		v(msg, err)
	}
}
