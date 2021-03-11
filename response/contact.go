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

type GetGroupChatDetailResp struct {
	ErrInfo
	GroupChat []*struct {
		ChatId     string `json:"chat_id"`
		Name       string `json:"name"`
		Owner      string `json:"owner"`
		CreateTime int64  `json:"create_time"`
		Notice     string `json:"notice"`
		MemberList []*struct {
			Userid    string `json:"userid"`
			Type      int    `json:"type"`
			JoinTime  int64  `json:"join_time"`
			JoinScene int    `json:"join_scene"`
			Unionid   string `json:"unionid"`
		} `json:"member_list"`
	} `json:"group_chat"`
}

type AddMsgTemplateResp struct {
	ErrInfo
	FailList []string `json:"fail_list"`
	Msgid    string   `json:"msgid"`
}

type GetGroupmsgListResp struct {
	ErrInfo
	NextCursor   string `json:"next_cursor"`
	GroupMsgList []*struct {
		Msgid      string `json:"msgid"`
		Creator    string `json:"creator"`
		CreateTime string `json:"create_time"`
		CreateType int    `json:"create_type"`
		Text       *struct {
			Content string `json:"content"`
		} `json:"text"`
		Image *struct {
			MediaId string `json:"media_id"`
		} `json:"image"`
		Link *struct {
			Title  string `json:"title"`
			Picurl string `json:"picurl"`
			Desc   string `json:"desc"`
			Url    string `json:"url"`
		} `json:"link"`
		Miniprogram *struct {
			Title string `json:"title"`
			Appid string `json:"appid"`
			Page  string `json:"page"`
		}
		Video *struct {
			MediaId string `json:"media_id"`
		} `json:"video"`
	}
}
