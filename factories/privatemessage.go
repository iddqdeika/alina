package factories

import (
	"alina/alina"
	"encoding/json"
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
	Attachments  []interface{} `json:"attachments"`             //
	Important    bool          `json:"important"`               //
	Geo          *geo          `json:"geo"`                     //
	PayLoad      string        `json:"payload"`                 //
	FwdMessages  []*fwdMessage `json:"fwd_messages"`            //
	ReplyMessage interface{}   `json:"reply_message"`           //
	Action       interface{}   `json:"action"`                  //
}

func (m *privateMessage) GetId() int {
	if m.ConvMsgId != 0 {
		return m.ConvMsgId
	}
	return m.Id
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

func (m *privateMessage) GetAttachments() []interface{} {
	return m.Attachments
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
	Attachments []interface{} `json:"attachments"`
	Date        int           `json:"date"`
	From_id     int           `json:"from_id"`
	Text        string        `json:"text"`
}

func (m *fwdMessage) GetAttachments() []interface{} {
	return m.Attachments
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
	Type        string      `json:"type"`
	Photo       interface{} `json:"photo"`
	Video       interface{} `json:"video"`
	Audio       interface{} `json:"audio"`
	Doc         interface{} `json:"doc"`
	Link        interface{} `json:"link"`
	Market      interface{} `json:"market"`
	MarketAlbum interface{} `json:"market_album"`
	Wall        interface{} `json:"wall"`
	WallReply   interface{} `json:"wall_reply"`
	Sticker     interface{} `json:"sticker"`
	Gift        interface{} `json:"gift"`
}
