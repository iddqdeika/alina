package alina

type AttachmentType string

const (
	PhotoAttachment       AttachmentType = "photo"
	VideoAttachment       AttachmentType = "video"
	AudioAttachment       AttachmentType = "audio"
	DocAttachment         AttachmentType = "doc"
	LinkAttachment        AttachmentType = "link"
	MarketAttachment      AttachmentType = "market"
	MarketAlbumAttachment AttachmentType = "market_album"
	WallAttachment        AttachmentType = "wall"
	WallReplyAttachment   AttachmentType = "wall_reply"
	StickerAttachment     AttachmentType = "sticker"
	GiftAttachment        AttachmentType = "gift"
)
