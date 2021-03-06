package request

type AddContactReq struct {
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

type GetContact struct {
	ConfigId string `json:"config_id"`
}

type DelContactWay struct {
	ConfigId string `json:"config_id"`
}
type UpdateContactWay struct {
	ConfigId      string      `json:"config_id"`
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

type GetFollowUserListReq struct {
}
