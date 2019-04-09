package main

import (
	"alina/alina"
	"alina/api/messagesapi"
	"alina/config"
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
	al, err := New(cfg.AccessToken, "5.85", cfg.GroupId, logger, cfg.LongPollInt)
	if err != nil {
		logger.Error(fmt.Sprintf("fatal error during Alina initialization: ", err))
		return
	}

	err = al.Init()
	if err != nil {
		logger.Error(fmt.Sprintf("fatal error during Alina initialization: %v", err))
		return
	}

	al.AddMessageHandler(func(message alina.PrivateMessage, e error) {
		if err != nil {
			logger.Error(err)
			return
		}
		if strings.Contains(message.GetText(), "лучшая жена") {
			al.GetMessagesApi().SendSimpleMessage(strconv.Itoa(message.GetPeerId()), "конечно Алина")
		}
	})

	//alinacore.GetMessagesApi().SendSimpleMessage("16729505", "йоу")
	//messages, err := alinacore.GetMessagesApi().GetHistory("16729505", 0, 200, "-1", nil)
	//for _, v := range messages {
	//	println(v.GetText())
	//}
	//if err != nil {
	//
	//}

	al.Run()

}

func New(token string, version string, groupid string, logger alina.Logger, longPollInterval int) (alina.Alina, error) {
	var cfg alina.Config
	var req alina.Requester
	al := &alinacore{}

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
