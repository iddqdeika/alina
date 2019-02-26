package alina

import "alina/config"

type Alina interface {
	Init(config Config) error
	Auth() error
	GetMessages() error
}

type Tokenizer interface {
	Init(config Config) error
	StartSession() error
	GetServer() (string, error)
	GetKey() (string, error)
}

type Config interface {
	GetKey() string
}

func New(key string) Alina {
	cfg := config.New(key)
	var a Alina
	a.Init(cfg)
	return a
}
