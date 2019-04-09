package alina

type PrivateMessagesFactory interface {
	NewPrivateMessageFromUpdate(data UpdateBody) (PrivateMessage, error)
	NewPrivateMessageFromInterface(messageBody interface{}) (PrivateMessage, error)
	NewPrivateMessage(data []byte) (PrivateMessage, error)
}
