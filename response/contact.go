package response

type AddContactResp struct {
	ErrInfo
	ConfigId string `json:"config_id"`
	QrCode   string `json:"qr_code"`
}

type DelContactWayResp struct {
	ErrInfo
}
type UpdateContactWayResp struct {
	ErrInfo
}

type GetContactWayResp struct {
	ErrInfo
	ContactWay ContactWay `json:"contact_way"`
}
type ContactWay struct {
	Type          int32       `json:"type"`
	Scene         int32       `json:"scene"`
	Style         int32       `json:"style"`
	Remark        string      `json:"remark"`
	SkipVerify    bool        `json:"skip_verify"`
	State         string      `json:"state"`
	User          []string    `json:"user"`
	Party         []int32     `json:"party"`
	IsTemp        bool        `json:"is_temp"`
	ExpiresIn     int32       `json:"expires_in"`
	ChatExpiresIn int32       `json:"chat_expires_in"`
	UnionId       string      `json:"union_id"`
	Conclusions   Conclusions `json:"conclusions"`
}
type Conclusions struct {
	Text        *Text        `json:"text"`
	Image       *Image       `json:"image"`
	Link        *Link        `json:"link"`
	Miniprogram *Miniprogram `json:"miniprogram"`
}
type Text struct {
	Content string `json:"content"`
}
type Image struct {
	MediaId string `json:"media_id"`
}
type Link struct {
	Title  string `json:"title"`
	Picurl string `json:"picurl"`
	Desc   string `json:"desc"`
	Url    string `json:"url"`
}
type Miniprogram struct {
	Title      string `json:"title"`
	PicMediaId string `json:"pic_media_id"`
	Appid      string `json:"appid"`
	Page       string `json:"page"`
}

type GetFollowUserListResp struct {
	ErrInfo
	FollowUser []string `json:"follow_user"`
}

type GetGroupChatListResp struct {
	ErrInfo
	GroupChatList []*GroupChatInfo `json:"group_chat_list"`
	NextCursor    string           `json:"next_cursor"`
}
type GroupChatInfo struct {
	ChatId string `json:"chat_id"`
	Status int    `json:"status"`
}
