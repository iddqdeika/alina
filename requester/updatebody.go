package requester

import "alina/alina"

type updateBody struct {
	UpdateType string      `json:"type"`
	Object     interface{} `json:"object"`
	GroupId    string      `json:group_id`
}

func (b *updateBody) GetType() alina.UpdateType {
	return alina.UpdateType(b.UpdateType)
}

func (b *updateBody) GetObject() interface{} {
	switch b.GetType() {

	}
	return b.Object
}
