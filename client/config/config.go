package config

func New(key string) *Config {
	return &Config{
		key: key,
	}
}

type Config struct {
	key string
}

func (c *Config) GetKey() string {
	return c.key
}
