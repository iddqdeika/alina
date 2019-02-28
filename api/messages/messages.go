package messages

import (
	"alina/definitions"
	"alina/definitions/api"
	"fmt"
)

func New(requester definitions.Requester, logger definitions.Logger) api.MessagesApi {
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

func (a *messagesApi) SendMessage(peerId string, message string) {
	params := make(map[string]string)
	params["peer_id"] = peerId
	params["message"] = message
	_, err := a.requester.SendGet("messages.send", params)
	if err != nil {
		a.logger.Error(fmt.Sprintf("error during sending message: %v", err))
	}

}
