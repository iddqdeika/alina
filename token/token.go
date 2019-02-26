package token

import "alina"

func New() alina.Tokenizer {
	return &tokenizer{}
}

type tokenizer struct {
}

func (*tokenizer) Init(config alina.Config) error {
	panic("implement me")
}

func (*tokenizer) StartSession() error {
	panic("implement me")
}

func (*tokenizer) GetServer() (string, error) {
	panic("implement me")
}

func (*tokenizer) GetKey() (string, error) {
	panic("implement me")
}
