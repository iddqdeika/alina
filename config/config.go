package config

import (
	"alina/alina"
	"time"
)

func NewConfig(token string, version string, groupid string, longPollIntervalInMillis int) alina.Config {
	return &config{
		token:                    token,
		version:                  version,
		groupid:                  groupid,
		longPollIntervalInMillis: time.Duration(longPollIntervalInMillis) * time.Millisecond,
	}
}

type config struct {
	token                    string
	version                  string
	groupid                  string
	longPollIntervalInMillis time.Duration
}

func (c *config) GetAccessToken() string {
	return c.token
}

func (c *config) GetVersion() string {
	return c.version
}

func (c *config) GetGroupId() string {
	return c.groupid
}

func (c *config) GetLongPollInterval() time.Duration {
	return c.longPollIntervalInMillis
}
