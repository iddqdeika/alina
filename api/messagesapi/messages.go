package messagesapi

import (
	"alina/definitions"
	"encoding/json"
	"fmt"
	"strconv"
)

func New(requester definitions.Requester, logger definitions.Logger, privateMessageFactory definitions.PrivateMessagesFactory) definitions.MessagesApi {
	mapi := &messagesApi{
		requester:             requester,
		logger:                logger,
		privateMessageFactory: privateMessageFactory,
	}

	return mapi
}

type messagesApi struct {
	requester             definitions.Requester
	logger                definitions.Logger
	privateMessageFactory definitions.PrivateMessagesFactory
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

//peerid is id of message destination
//offset defines offset of messages got (due to max messages count is 200)
//startMessageId is id of start message if known. to retrieve last messages use -1
//fields are names of message object data additional fields you need to retrieve
func (a *messagesApi) GetHistory(peerId string, offset int, count int, startMessageId string, fields []string) ([]definitions.PrivateMessage, error) {
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

	messages := make([]definitions.PrivateMessage, 0)

	for _, val := range responseBody.Response.Items {
		var msg definitions.PrivateMessage
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
