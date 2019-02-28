package main

import (
	"alina/definitions"
	"alina/intitializer"
	"alina/logger"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	alina, err := initializer.New(cfg.AccessToken, "5.85", cfg.GroupId, logger, cfg.LongPollInt)
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
		text := message.GetText()
		from := message.GetPeerId()
		logger.Info(fmt.Sprintf("incoming message: %v from %v\r\n", text, from))
	})

	alina.GetMessagesApi().SendMessage("16729505", "йоу")

	alina.Run()

}
