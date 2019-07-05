package alina

type MessagesApi interface {
	SendSimpleMessage(peerId string, message string)
	GetHistory(peerId string, offset int, count int, startMessageId string, fields []string) ([]PrivateMessage, error)
	SendMessageWithForward(peerId string, message string, forward_messages []string)
	GetConversationMessageId(peerId string, convMessagesId string) (string, error)
}
