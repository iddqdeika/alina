package alina

type PrivateMessage interface {
	GetId() int
	GetConversationMessageID() int
	GetDate() int
	GetPeerId() int
	GetFromId() int
	GetText() string
	GetRandomId() int
	GetRef() string
	GetRefSource() string
	GetAttachments() []Attachment
	IsImportant() bool
	GetGeo() interface{}
	GetPayLoad() string
	GetFwdMessages() []FwdMessage
	GetReplyMessage() interface{}
	GetAction() interface{}
}

type FwdMessage interface {
	GetAttachments() []Attachment
	GetDate() int
	GetFromID() int
	GetText() string
}

type Attachment interface {
	GetType() AttachmentType
	IsMedia() bool
	GetAsPhoto() (Photo, error)
}

type Photo interface {
	GetId() int
	GetAlbumId() int
	GetOwnerId() int
	GetUserId() int
	GetText() string
	GetDate() int
	GetSizes() []interface{}
	GetWidth() int
	GetHeight() int
}
