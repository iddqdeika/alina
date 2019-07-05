package main

import (
	"alina/alina"
	"alina/alinafactory"
	"alina/logger"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

const (
	cfgName      = "config.cfg"
	vkApiVersion = "5.85"
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
	al, err := alinafactory.New(cfg.AccessToken, vkApiVersion, cfg.GroupId, logger, cfg.LongPollInt)
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

		test(message)
		if message.GetConversationMessageID() > 0 {

			msgs, err := al.GetMessagesApi().GetHistory("2000000002", -1, 10, "", []string{})
			if err != nil {

			}
			if msgs != nil {

			}
			al.GetMessagesApi().GetConversationMessageId(strconv.Itoa(message.GetPeerId()), strconv.Itoa(message.GetConversationMessageID()))
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

func test(message alina.PrivateMessage) {
	println("str")

}
