package alinafactory

import (
	"alina/alina"
	"alina/api/messagesapi"
	"alina/config"
	"alina/dispatcher"
	"alina/factories"
	"alina/requester"
)

func New(token string, version string, groupid string, logger alina.Logger, longPollInterval int) (alina.Alina, error) {
	var cfg alina.Config
	var req alina.Requester
	al := &alinacore{}

	cfg = config.NewConfig(token, version, groupid, longPollInterval)

	dispatcher := dispatcher.New(factories.GetPrivateMessageFactory(), logger)
	al.dispatcher = dispatcher

	req, err := requester.New(cfg, logger, dispatcher)
	if err != nil {
		return nil, err
	}
	al.requester = req

	msg := messagesapi.New(req, logger, factories.GetPrivateMessageFactory())
	al.messagesApi = msg

	return al, nil
}

type alinacore struct {
	requester   alina.Requester
	dispatcher  alina.Dispatcher
	messagesApi alina.MessagesApi
}

func (a *alinacore) GetMessagesApi() alina.MessagesApi {
	return a.messagesApi
}

func (a *alinacore) AddMessageHandler(handler func(alina.PrivateMessage, error)) {
	a.dispatcher.AddMessageHandler(handler)
}

func (a *alinacore) Init() error {
	err := a.requester.Init()
	if err != nil {
		return err
	}
	return nil
}

func (a *alinacore) Run() {
	a.requester.Run()
}
