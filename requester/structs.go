package requester

type getLongPollServerStruct struct {
	Response getLongPollServerResponse `json:"response"`
}

type getLongPollServerResponse struct {
	Key    string `json:"key"`
	Server string `json:"server"`
	Ts     string `json:"ts"`
}

type updateResponse struct {
	data []byte
	err  error
}

type updatesResponseBody struct {
	Ts      string        `json:"ts"`
	Updates []*updateBody `json:"updates"`
}
