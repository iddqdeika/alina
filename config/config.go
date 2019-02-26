package config

import "alina"

func New(key string) alina.Config {
	return &config{
		key: key,
	}
}

type config struct {
	key string
}

func (c *config) GetKey() string {
	return c.key
}
