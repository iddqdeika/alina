package main

import (
	"alina/api/messagesapi"
	"alina/config"
	"alina/definitions"
	"alina/dispatcher"
	"alina/factories"
	"alina/logger"
	"alina/requester"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	cfgName = "config.cfg"
)

func main() {

	cfg := &struct {
		AccessToken string `json:"access_token"`
		GroupId     string `json:"group_id"`
		LongPollInt int    `json:"long_poll_interval"`
	}{}

	cfgData, err := ioutil.ReadFile(cfgName)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(cfgData, cfg)
	if err != nil || len(cfg.AccessToken) == 0 {
		panic(err)
	}

	logger.InitDefaultLogger()
	logger := logger.DefaultLogger
	alina, err := New(cfg.AccessToken, "5.85", cfg.GroupId, logger, cfg.LongPollInt)
	if err != nil {
		logger.Error(fmt.Sprintf("fatal error during Alina initialization: ", err))
		return
	}

	err = alina.Init()
	if err != nil {
		logger.Error(fmt.Sprintf("fatal error during Alina initialization: %v", err))
		return
	}

	alina.AddMessageHandler(func(message definitions.PrivateMessage, e error) {
		if err != nil {
			logger.Error(err)
			return
		}
		if strings.Contains(message.GetText(), "лучшая жена") {
			alina.GetMessagesApi().SendSimpleMessage(strconv.Itoa(message.GetPeerId()), "конечно Алина")
		}
	})

	//alina.GetMessagesApi().SendSimpleMessage("16729505", "йоу")
	//messages, err := alina.GetMessagesApi().GetHistory("16729505", 0, 200, "-1", nil)
	//for _, v := range messages {
	//	println(v.GetText())
	//}
	//if err != nil {
	//
	//}

	alina.Run()

}

func New(token string, version string, groupid string, logger definitions.Logger, longPollInterval int) (definitions.Alina, error) {
	var cfg definitions.Config
	var req definitions.Requester
	al := &alina{}

	cfg = config.NewConfig(token, version, groupid, longPollInterval)

	dispatcher := dispatcher.New(factories.GetPrivateMessageFactory())
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

type alina struct {
	requester   definitions.Requester
	dispatcher  definitions.Dispatcher
	messagesApi definitions.MessagesApi
}

func (a *alina) GetMessagesApi() definitions.MessagesApi {
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
