package response

type TransferCustomerResp struct {
	ErrInfo
	Customer []*CustomerItem `json:"customer"`
}

type CustomerItem struct {
	ExternalUserid string `json:"external_userid"`
	Errcode        int    `json:"errcode"`
}

type TransferResultResp struct {
}

type GetUnassignedListResp struct {
	ErrInfo
	Info       []*InfoItem `json:"info"`
	IsLast     bool        `json:"is_last"`
	NextCursor string      `json:"next_cursor"`
}
type InfoItem struct {
	HandoverUserid string `json:"handover_userid"`
	ExternalUserid string `json:"external_userid"`
	DimissionTime  int64  `json:"dimission_time"`
}

type ResignedTransferCustomerResp struct {
	ErrInfo
	Customer []*CustomerItem `json:"customer"`
}

type GetResignedTransferResultResp struct {
	ErrInfo
	Customer []*ResignedTransferResultItem `json:"customer"`
}
type ResignedTransferResultItem struct {
	ExternalUserid string `json:"external_userid"`
	Status         int    `json:"status"`
	TakeoverTime   int64  `json:"takeover_time"`
}

type ResignedTransferChatGroupResp struct {
	ErrInfo
	FailedChatList []*FailedChat `json:"failed_chat_list"`
}
type FailedChat struct {
	ChatId string `json:"chat_id"`
	ErrInfo
}
