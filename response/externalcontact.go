package response

type GetContactListResp struct {
	ErrInfo
	ExternalUserId []string `json:"external_user_id"`
}
type GetExtContactDetail struct {
	ErrInfo
	ExternalContact ExternalContact `json:"external_contact"`
	FollowUser      []FollowUser    `json:"follow_user"`
}

type ExternalContact struct {
	ExternalUserid  string          `json:"external_userid"`
	Name            string          `json:"name"`
	Position        string          `json:"position"`
	Avatar          string          `json:"avatar"`
	CorpName        string          `json:"corp_name"`
	CorpFullName    string          `json:"corp_full_name"`
	Type            int32           `json:"type"`
	Gender          int32           `json:"gender"`
	Unionid         string          `json:"unionid"`
	ExternalProfile ExternalProfile `json:"external_profile"`
}

type FollowUser struct {
	Userid         string          `json:"userid"`
	FollowerAvatar string          `json:"follower_avatar"`
	FollowerName   string          `json:"follower_name"`
	CorpId         string          `json:"corp_id"`
	Remark         string          `json:"remark"`
	Description    string          `json:"description"`
	Createtime     int64           `json:"createtime"`
	Tags           []FollowUserTag `json:"tags"`
	RemarkCorpName string          `json:"remark_corp_name"`
	RemarkMobiles  []string        `json:"remark_mobiles"`
	OperUserid     string          `json:"oper_userid"`
	OperUserName   string          `json:"oper_user_name"`
	AddWay         int32           `json:"add_way"`
	DeleteWay      int32           `json:"delete_way"`
}
type ExternalProfile struct {
	ExternalAttr []ExternalAttr `json:"external_attr"`
}

type ExternalAttr struct {
	Type    int32  `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
}
type FollowUserTag struct {
	GroupName string `json:"group_name"`
	TagName   string `json:"tag_name"`
	TagId     string `json:"tag_id"`
	Type      int32  `json:"type"`
}

type ExtUserRemark struct {
	UserId         string `json:"user_id"`
	ExternalUserId string `json:"external_user_id"`
	Remark         string `json:"remark"`
	Description    string `json:"description"`
}
type BatchGetExtContactDetailResp struct {
	ErrInfo
	ExternalContactList []*ExternalContactItem `json:"external_contact_list"`
}
type ExternalContactItem struct {
	ExternalContact ExternalContact `json:"external_contact"`
	FollowInfo      FollowInfo      `json:"follow_info"`
}
type FollowInfo struct {
	Remark         string   `json:"remark"`
	Description    string   `json:"description"`
	Createtime     int64    `json:"createtime"`
	TagId          []string `json:"tag_id"`
	RemarkCorpName string   `json:"remark_corp_name"`
	RemarkMobiles  []string `json:"remark_mobiles"`
	OperUserId     string   `json:"oper_user_id"`
	AddWay         int      `json:"add_way"`
}

type UpdateExtContactRemarkResp struct {
	ErrInfo
}
