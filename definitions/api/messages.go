package api

type MessagesApi interface {
	SendMessage(peerId string, message string)
}
