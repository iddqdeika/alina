package requester

import "alina/definitions"

type updateBody struct {
	UpdateType string      `json:"type"`
	Object     interface{} `json:"object"`
	GroupId    string      `json:group_id`
}

func (b *updateBody) GetType() definitions.UpdateType {
	return definitions.UpdateType(b.UpdateType)
}

func (b *updateBody) GetObject() interface{} {
	switch b.GetType() {

	}
	return b.Object
}
