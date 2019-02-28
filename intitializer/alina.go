package initializer

import (
	"alina/api/messages"
	"alina/config"
	"alina/definitions"
	"alina/definitions/api"
	"alina/dispatcher"
	"alina/requester"
)

func New(token string, version string, groupid string, logger definitions.Logger, longPollInterval int) (definitions.Alina, error) {
	var cfg definitions.Config
	var req definitions.Requester
	al := &alina{}

	cfg = config.New(token, version, groupid, longPollInterval)

	dispatcher := dispatcher.New()
	al.dispatcher = dispatcher

	req, err := requester.New(cfg, logger, dispatcher)
	if err != nil {
		return nil, err
	}
	al.requester = req

	msg := messages.New(req, logger)
	al.messagesApi = msg

	return al, nil
}

type alina struct {
	requester   definitions.Requester
	dispatcher  definitions.Dispatcher
	messagesApi api.MessagesApi
}

func (a *alina) GetMessagesApi() api.MessagesApi {
	return a.messagesApi
}

func (a *alina) AddMessageHandler(handler func(definitions.PrivateMessage, error)) {
	a.dispatcher.AddMessageHandler(handler)
}

func (a *alina) Init() error {
	err := a.requester.Init()
	if err != nil {
		return err
	}
	return nil
}

func (a *alina) Run() {
	a.requester.Run()
}
