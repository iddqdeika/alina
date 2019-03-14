package messagesapi

type messagesHistoryResponseBody struct {
	Response messagesHistoryResponse `json:"response"`
}

type messagesHistoryResponse struct {
	Count int           `json:"count"`
	Items []interface{} `json:"items"`
}
