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

type Video interface {
	GetId() int
	GetOwnerId() int
	GetTitle() string
	GetDescription() string
	GetDuration() int
	GetPhoto130() string
	GetPhoto320() string
	GetPhoto640() string
	GetPhoto800() string
	GetPhoto1280() string
	GetFirstFrame130() string
	GetFirstFrame320() string
	GetFirstFrame640() string
	GetFirstFrame800() string
	GetFirstFrame1280() string
	GetDate() int
	GetAddingDate() int
	GetViews() int
	GetComments() int
	GetPlayer() string
	GetPlatform() string
	GetCanEdit() int
	GetCanAdd() int
	GetIsPrivate() int
	GetAccessKey() string
	GetProcessing() int
	GetLive() int
	GetUpcoming() int
	GetIsFavorite() bool
}

type Audio interface {
	GetId() int
	GetOwnerId() int
	GetArtist() string
	GetTitle() string
	GetDuration() int
	GetUrl() string
	GetLyricsId() int
	GetAlbumId() int
	GetGenreId() int
	GetDate() int
	GetNoSearch() int
	GetIsHq() int
}
