package definitions

import (
	"alina/definitions/api"
	"time"
)

type Alina interface {
	Init() error
	Run()
	AddMessageHandler(func(PrivateMessage, error))
	GetMessagesApi() api.MessagesApi
}

type Requester interface {
	Init() error
	Run()
	GetServer() (string, error)
	GetKey() (string, error)
	SendGet(methodName string, paramMap map[string]string) ([]byte, error)
}

type Config interface {
	GetAccessToken() string
	GetVersion() string
	GetGroupId() string
	GetLongPollInterval() time.Duration
}

type Logger interface {
	Info(args ...interface{})
	Error(args ...interface{})
}

type Dispatcher interface {
	AddMessageHandler(func(PrivateMessage, error))
	Handle(update UpdateBody)
}

type UpdateBody interface {
	GetType() UpdateType
	GetObject() interface{}
}
