package messagesapi

import (
	"alina/alina"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func New(requester alina.Requester, logger alina.Logger, privateMessageFactory alina.PrivateMessagesFactory) alina.MessagesApi {
	mapi := &messagesApi{
		requester:             requester,
		logger:                logger,
		privateMessageFactory: privateMessageFactory,
	}

	return mapi
}

type messagesApi struct {
	requester             alina.Requester
	logger                alina.Logger
	privateMessageFactory alina.PrivateMessagesFactory
}

func (a *messagesApi) SendMessageWithAttachment(peerId string, message string, attachment string) {
	params := make(map[string]string)
	params["peer_id"] = peerId
	params["message"] = message
	params["attachment"] = attachment
	res, err := a.requester.SendGet("messages.send", params)
	if err != nil {
		a.logger.Error(fmt.Sprintf("error during sending message: %v, %v", err, string(res)))
	}
}

func (a *messagesApi) GetConversationMessageId(peerId string, convMessagesId string) (string, error) {
	params := make(map[string]string)
	params["peer_id"] = peerId
	params["conversation_message_ids"] = convMessagesId
	res, err := a.requester.SendGet("messages.getByConversationMessageId", params)
	if err != nil {
		a.logger.Error(fmt.Sprintf("error during gettinf message id for conversation message: %v, %v", err, string(res)))
	}

	str := string(res)

	return str, nil
}

func (a *messagesApi) SendSimpleMessage(peerId string, message string) {
	params := make(map[string]string)
	params["peer_id"] = peerId
	params["message"] = message
	res, err := a.requester.SendGet("messages.send", params)
	if err != nil {
		a.logger.Error(fmt.Sprintf("error during sending message: %v, %v", err, string(res)))
	}
}

func (a *messagesApi) SendMessageWithForward(peerId string, message string, forward_messages []string) {
	params := make(map[string]string)
	params["peer_id"] = peerId
	params["message"] = message
	params["forward_messages"] = strings.Join(forward_messages, ",")
	res, err := a.requester.SendGet("messages.send", params)
	if err != nil {
		a.logger.Error(fmt.Sprintf("error during sending message with forwarding: %v, %v", err, string(res)))
	}
}

//peerid is id of message destination
//offset defines offset of messages got (due to max messages count is 200)
//startMessageId is id of start message if known. to retrieve last messages use -1
//fields are names of message object data additional fields you need to retrieve
func (a *messagesApi) GetHistory(peerId string, offset int, count int, startMessageId string, fields []string) ([]alina.PrivateMessage, error) {
	params := make(map[string]string)
	params["peer_id"] = peerId
	params["offset"] = strconv.Itoa(offset)
	params["count"] = strconv.Itoa(count)
	if len(startMessageId) > 0 {
		params["start_message_id"] = startMessageId
	}
	if fields != nil || len(fields) != 0 {
		params["fields"] = toList(fields)
	}
	data, err := a.requester.SendGet("messages.getHistory", params)

	if err != nil {
		return nil, err
	}

	responseBody := &messagesHistoryResponseBody{}

	err = json.Unmarshal(data, responseBody)
	if err != nil {
		return nil, err
	}

	messages := make([]alina.PrivateMessage, 0)

	for _, val := range responseBody.Response.Items {
		var msg alina.PrivateMessage
		msg, err = a.privateMessageFactory.NewPrivateMessageFromInterface(val)
		if err != nil {
			return nil, fmt.Errorf("error during parsing message from history: %v", err)
		}
		messages = append(messages, msg)
	}

	return messages, nil
}

func toList(list []string) string {
	if list == nil || len(list) == 0 {
		return ""
	}
	result := list[0]
	if len(list) > 1 {
		for _, val := range list {
			result = result + "," + val
		}
	}
	return result
}
