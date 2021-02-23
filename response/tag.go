package response

type GetCorpTagListResp struct {
	ErrInfo
	TagGroup []*TagGroupItem `json:"tag_group"`
}
type TagGroupItem struct {
	GroupId    string     `json:"group_id"`
	GroupName  string     `json:"group_name"`
	CreateTime int64      `json:"create_time"`
	Order      int        `json:"order"`
	Deleted    bool       `json:"deleted"`
	Tag        []*TagItem `json:"tag"`
}
type TagItem struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	CreateTime int64  `json:"create_time"`
	Order      int    `json:"order"`
	Deleted    bool   `json:"deleted"`
}

type AddCorpTagResp struct {
	ErrInfo
	TagGroup TagGroupItem `json:"tag_group"`
}

type UpdateCorpTagResp struct {
	ErrInfo
}
type DelCorpTagResp struct {
	ErrInfo
}

type SetExtContactTagResp struct {
	ErrInfo
}
