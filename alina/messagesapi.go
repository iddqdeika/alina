package alina

type MessagesApi interface {
	SendSimpleMessage(peerId string, message string) ([]byte, error)
	SendMessageWithAttachment(peerId string, message string, attachment string) ([]byte, error)
	SendMessageWithForward(peerId string, message string, forward_messages []string) ([]byte, error)
	GetHistory(peerId string, offset int, count int, startMessageId string, fields []string) ([]PrivateMessage, error)
	GetConversationMessageId(peerId string, convMessagesId string) (string, error)
}
