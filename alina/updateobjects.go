package alina

type PrivateMessage interface {
	GetId() int
	GetDate() int
	GetPeerId() int
	GetFromId() int
	GetText() string
	GetRandomId() int
	GetRef() string
	GetRefSource() string
	GetAttachments() []interface{}
	IsImportant() bool
	GetGeo() interface{}
	GetPayLoad() string
	GetFwdMessages() []FwdMessage
	GetReplyMessage() interface{}
	GetAction() interface{}
}

type FwdMessage interface {
	GetAttachments() []interface{}
	GetDate() int
	GetFromID() int
	GetText() string
}
