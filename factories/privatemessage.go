package factories

import (
	"alina/alina"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

var privateMessageF = &privateMessageFactory{}

func GetPrivateMessageFactory() alina.PrivateMessagesFactory {
	return privateMessageF
}

type privateMessageFactory struct {
}

func (f *privateMessageFactory) NewPrivateMessageFromUpdate(data alina.UpdateBody) (alina.PrivateMessage, error) {
	return f.NewPrivateMessageFromInterface(data.GetObject())
}

func (f *privateMessageFactory) NewPrivateMessageFromInterface(messageBody interface{}) (alina.PrivateMessage, error) {
	bts, err := json.Marshal(messageBody)
	if err != nil {
		return nil, err
	}
	return f.NewPrivateMessage(bts)
}

func (f *privateMessageFactory) NewPrivateMessage(data []byte) (alina.PrivateMessage, error) {
	pm := &privateMessage{}

	err := json.Unmarshal(data, pm)
	if err != nil {
		return nil, err
	}
	return pm, nil
}

type privateMessage struct {
	Id           int           `json:"id"`                      //
	ConvMsgId    int           `json:"conversation_message_id"` //
	Date         int           `json:"date"`                    //
	PeerId       int           `json:"peer_id"`                 //
	FromId       int           `json:"from_id"`                 //
	Text         string        `json:"text"`                    //
	RandomId     int           `json:"random_id"`               //
	Ref          string        `json:"ref"`                     //
	RefSource    string        `json:"ref_source"`              //
	Attachments  []*attachment `json:"attachments"`             //
	Important    bool          `json:"important"`               //
	Geo          *geo          `json:"geo"`                     //
	PayLoad      string        `json:"payload"`                 //
	FwdMessages  []*fwdMessage `json:"fwd_messages"`            //
	ReplyMessage interface{}   `json:"reply_message"`           //
	Action       interface{}   `json:"action"`                  //
}

func (m *privateMessage) GetId() int {
	return m.Id
}

func (m *privateMessage) GetConversationMessageID() int {
	return m.ConvMsgId
}

func (m *privateMessage) GetDate() int {
	return m.Date
}

func (m *privateMessage) GetPeerId() int {
	return m.PeerId
}

func (m *privateMessage) GetFromId() int {
	return m.FromId
}

func (m *privateMessage) GetText() string {
	return m.Text
}

func (m *privateMessage) GetRandomId() int {
	return m.RandomId
}

func (m *privateMessage) GetRef() string {
	return m.Ref
}

func (m *privateMessage) GetRefSource() string {
	return m.RefSource
}

func (m *privateMessage) GetAttachments() []alina.Attachment {
	result := make([]alina.Attachment, 0)
	for _, v := range m.Attachments {
		result = append(result, v)
	}
	return result
}

func (m *privateMessage) IsImportant() bool {
	return m.Important
}

func (m *privateMessage) GetGeo() interface{} {
	return m.Geo
}

func (m *privateMessage) GetPayLoad() string {
	return m.PayLoad
}

func (m *privateMessage) GetFwdMessages() []alina.FwdMessage {
	result := make([]alina.FwdMessage, 0)
	for _, v := range m.FwdMessages {
		result = append(result, v)
	}
	return result
}

func (m *privateMessage) GetReplyMessage() interface{} {
	return m.ReplyMessage
}

func (m *privateMessage) GetAction() interface{} {
	return m.Action
}

type geo struct {
	Type        string `json:"type"`
	Coordinates string `json:"coordinates"`
	Place       place  `json:"place"`
}

type place struct {
	Id        int     `json:"id"`
	Title     string  `json:"title"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longiture"`
	Created   int     `json:"created"`
	Icon      string  `json:"icon"`
	Country   string  `json:"country"`
	City      string  `json:"city"`
}

type fwdMessage struct {
	Attachments []*attachment `json:"attachments"`
	Date        int           `json:"date"`
	From_id     int           `json:"from_id"`
	Text        string        `json:"text"`
}

func (m *fwdMessage) GetAttachments() []alina.Attachment {
	result := make([]alina.Attachment, 0)
	for _, v := range m.Attachments {
		result = append(result, v)
	}
	return result
}

func (m *fwdMessage) GetDate() int {
	return m.Date
}

func (m *fwdMessage) GetFromID() int {
	return m.From_id
}

func (m *fwdMessage) GetText() string {
	return m.Text
}

type attachment struct {
	Type        alina.AttachmentType `json:"type"`
	Photo       *photo               `json:"photo"`
	Video       *video               `json:"video"`
	Audio       *audio               `json:"audio"`
	Doc         interface{}          `json:"doc"`
	Link        interface{}          `json:"link"`
	Market      interface{}          `json:"market"`
	MarketAlbum interface{}          `json:"market_album"`
	Wall        interface{}          `json:"wall"`
	WallReply   interface{}          `json:"wall_reply"`
	Sticker     interface{}          `json:"sticker"`
	Gift        interface{}          `json:"gift"`
}

func (a *attachment) GetPrivateMessageToken() (string, error) {
	if a.IsMedia() {
		switch a.Type {
		case alina.PhotoAttachment:
			photo, err := a.GetAsPhoto()
			if err != nil {
				return "", fmt.Errorf("error while token generation: %v", err)
			}
			return photo.GetPrivateMessageToken(), nil
		case alina.VideoAttachment:
			video, err := a.GetAsVideo()
			if err != nil {
				return "", fmt.Errorf("error while token generation: %v", err)
			}
			return video.GetPrivateMessageToken(), nil
		case alina.AudioAttachment:
			audio, err := a.GetAsAudio()
			if err != nil {
				return "", fmt.Errorf("error while token generation: %v", err)
			}
			return audio.GetPrivateMessageToken(), nil
		default:
			return "", errors.New("unimplemented")
		}
	}
	return "", errors.New("attachment is not media")
}

func (a *attachment) GetType() alina.AttachmentType {
	return a.Type
}

func (a *attachment) IsMedia() bool {
	return a.Type == alina.PhotoAttachment || a.Type == alina.VideoAttachment || a.Type == alina.AudioAttachment || a.Type == alina.DocAttachment
}

func (a *attachment) GetAsPhoto() (alina.Photo, error) {
	if a.Type == alina.PhotoAttachment {
		return a.Photo, nil
	}
	return nil, errors.New("incorrect type")
}

func (a *attachment) GetAsVideo() (alina.Video, error) {
	if a.Type == alina.VideoAttachment {
		return a.Video, nil
	}
	return nil, errors.New("incorrect type")
}

func (a *attachment) GetAsAudio() (alina.Audio, error) {
	if a.Type == alina.AudioAttachment {
		return a.Audio, nil
	}
	return nil, errors.New("incorrect type")
}

type photo struct {
	Id      int           `json:"id"`
	AlbumId int           `json:"album_id"`
	OwnerId int           `json:"owner_id"`
	UserId  int           `json:"user_id"`
	Text    string        `json:"text"`
	Date    int           `json:"date"`
	Sizes   []interface{} `json:"sizes"`
	Width   int           `json:"width"`
	Height  int           `json:"height"`
}

func (p *photo) GetPrivateMessageToken() string {
	return string(alina.PhotoAttachment) + strconv.Itoa(p.OwnerId) + "_" + strconv.Itoa(p.Id)
}

func (p *photo) GetId() int {
	return p.Id
}

func (p *photo) GetAlbumId() int {
	return p.AlbumId
}

func (p *photo) GetOwnerId() int {
	return p.OwnerId
}

func (p *photo) GetUserId() int {
	return p.UserId
}

func (p *photo) GetText() string {
	return p.Text
}

func (p *photo) GetDate() int {
	return p.Date
}

func (p *photo) GetSizes() []interface{} {
	return p.Sizes
}

func (p *photo) GetWidth() int {
	return p.Width
}

func (p *photo) GetHeight() int {
	return p.Height
}

type video struct {
	Id             int    `json:"id"`
	OwnerId        int    `json:"owner_id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Duration       int    `json:"duration"`
	Photo130       string `json:"photo_130"`
	Photo320       string `json:"photo_320"`
	Photo640       string `json:"photo_640"`
	Photo800       string `json:"photo_800"`
	Photo1280      string `json:"photo_1280"`
	FirstFrame130  string `json:"first_frame_130"`
	FirstFrame320  string `json:"first_frame_320"`
	FirstFrame640  string `json:"first_frame_640"`
	FirstFrame800  string `json:"first_frame_800"`
	FirstFrame1280 string `json:"first_frame_1280"`
	Date           int    `json:"date"`
	AddingDate     int    `json:"adding_date"`
	Views          int    `json:"views"`
	Comments       int    `json:"comments"`
	Player         string `json:"player"`
	Platform       string `json:"platform"`
	CanEdit        int    `json:"can_edit"`
	CanAdd         int    `json:"can_add"`
	IsPrivate      int    `json:"is_private"`
	AccessKey      string `json:"access_key"`
	Processing     int    `json:"processing"`
	Live           int    `json:"live"`
	Upcoming       int    `json:"upcoming"`
	IsFavorite     bool   `json:"is_favorite"`
}

func (v *video) GetPrivateMessageToken() string {
	if len(v.AccessKey) > 0 {
		return string(alina.VideoAttachment) + strconv.Itoa(v.OwnerId) + "_" + strconv.Itoa(v.Id) + "_" + v.AccessKey
	}
	return string(alina.VideoAttachment) + strconv.Itoa(v.OwnerId) + "_" + strconv.Itoa(v.Id)
}

func (v *video) GetId() int {
	return v.Id
}

func (v *video) GetOwnerId() int {
	return v.OwnerId
}

func (v *video) GetTitle() string {
	return v.Title
}

func (v *video) GetDescription() string {
	return v.GetDescription()
}

func (v *video) GetDuration() int {
	return v.Duration
}

func (v *video) GetPhoto130() string {
	return v.Photo130
}

func (v *video) GetPhoto320() string {
	return v.Photo320
}

func (v *video) GetPhoto640() string {
	return v.Photo640
}

func (v *video) GetPhoto800() string {
	return v.Photo800
}

func (v *video) GetPhoto1280() string {
	return v.Photo1280
}

func (v *video) GetFirstFrame130() string {
	return v.FirstFrame130
}

func (v *video) GetFirstFrame320() string {
	return v.FirstFrame320
}

func (v *video) GetFirstFrame640() string {
	return v.FirstFrame640
}

func (v *video) GetFirstFrame800() string {
	return v.FirstFrame800
}

func (v *video) GetFirstFrame1280() string {
	return v.FirstFrame1280
}

func (v *video) GetDate() int {
	return v.Date
}

func (v *video) GetAddingDate() int {
	return v.AddingDate
}

func (v *video) GetViews() int {
	return v.Views
}

func (v *video) GetComments() int {
	return v.Comments
}

func (v *video) GetPlayer() string {
	return v.Player
}

func (v *video) GetPlatform() string {
	return v.Platform
}

func (v *video) GetCanEdit() int {
	return v.CanEdit
}

func (v *video) GetCanAdd() int {
	return v.CanAdd
}

func (v *video) GetIsPrivate() int {
	return v.IsPrivate
}

func (v *video) GetAccessKey() string {
	return v.AccessKey
}

func (v *video) GetProcessing() int {
	return v.Processing
}

func (v *video) GetLive() int {
	return v.Live
}

func (v *video) GetUpcoming() int {
	return v.Upcoming
}

func (v *video) GetIsFavorite() bool {
	return v.IsFavorite
}

type audio struct {
	Id       int    `json:"id"`
	OwnerId  int    `json:"owner_id"`
	Artist   string `json:"artist"`
	Title    string `json:"title"`
	Duration int    `json:"duration"`
	Url      string `json:"url"`
	LyricsId int    `json:"lyrics_id"`
	AlbumId  int    `json:"album_id"`
	GenreId  int    `json:"genre_id"`
	Date     int    `json:"date"`
	NoSearch int    `json:"no_search"`
	IsHq     bool   `json:"is_hq"`
}

func (a *audio) GetPrivateMessageToken() string {
	return string(alina.AudioAttachment) + strconv.Itoa(a.OwnerId) + "_" + strconv.Itoa(a.Id)
}

func (a *audio) GetId() int {
	return a.Id
}

func (a *audio) GetOwnerId() int {
	return a.OwnerId
}

func (a *audio) GetArtist() string {
	return a.Artist
}

func (a *audio) GetTitle() string {
	return a.Title
}

func (a *audio) GetDuration() int {
	return a.Duration
}

func (a *audio) GetUrl() string {
	return a.Url
}

func (a *audio) GetLyricsId() int {
	return a.LyricsId
}

func (a *audio) GetAlbumId() int {
	return a.AlbumId
}

func (a *audio) GetGenreId() int {
	return a.GenreId
}

func (a *audio) GetDate() int {
	return a.Date
}

func (a *audio) GetNoSearch() int {
	return a.NoSearch
}

func (a *audio) GetIsHq() bool {
	return a.IsHq
}
