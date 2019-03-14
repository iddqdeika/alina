package messages

import (
	"alina/definitions"
	"alina/objects"
	"encoding/json"
	"fmt"
	"strconv"
)

func New(requester definitions.Requester, logger definitions.Logger) definitions.MessagesApi {
	mapi := &messagesApi{
		requester: requester,
		logger:    logger,
	}

	return mapi
}

type messagesApi struct {
	requester definitions.Requester
	logger    definitions.Logger
}

func (a *messagesApi) SendSimpleMessage(peerId string, message string) {
	params := make(map[string]string)
	params["peer_id"] = peerId
	params["message"] = message
	_, err := a.requester.SendGet("messages.send", params)
	if err != nil {
		a.logger.Error(fmt.Sprintf("error during sending message: %v", err))
	}
}

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
		msg, err = objects.NewPrivateMessageFromInterface(val)
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
