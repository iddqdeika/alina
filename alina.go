package main

import "alina/client/config"

type Alina interface {
	Init(config Config) error
	Auth() error
	GetMessages() error
}

type Token interface {
}

type Config interface {
	GetKey() string
}

func New(key string) Alina {
	cfg := config.New(key)

}
