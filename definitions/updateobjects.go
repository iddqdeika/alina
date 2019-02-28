package definitions

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
	GetFwdMessages() []interface{}
	GetReplyMessage() interface{}
	GetAction() interface{}
}
