package request

type TransferCustomerReq struct {
	HandoverUserid     string   `json:"handover_userid"`
	TakeoverUserid     string   `json:"takeover_userid"`
	ExternalUserid     []string `json:"external_userid"`
	TransferSuccessMsg string   `json:"transfer_success_msg"`
}

type TransferResultReq struct {
	HandoverUserid string `json:"handover_userid"`
	TakeoverUserid string `json:"takeover_userid"`
	Cursor         string `json:"cursor"`
}

type GetUnassignedListReq struct {
	PageId   int    `json:"page_id"`
	Cursor   string `json:"cursor"`
	PageSize int    `json:"page_size"`
}

type ResignedTransferCustomerReq struct {
	HandoverUserid string   `json:"handover_userid"`
	TakeoverUserid string   `json:"takeover_userid"`
	ExternalUserid []string `json:"external_userid"`
}

type GetResignedTransferResultReq struct {
	HandoverUserid string `json:"handover_userid"`
	TakeoverUserid string `json:"takeover_userid"`
	Cursor         string `json:"cursor"`
}

type ResignedTransferChatGroupReq struct {
	ChatIdList []string `json:"chat_id_list"`
	NewOwner   string   `json:"new_owner"`
}

type GetGroupChatListReq struct {
	StatusFilter int `json:"status_filter"`
	OwnerFilter  struct {
		UseridList []string `json:"userid_list"`
	} `json:"owner_filter"`
	Cursor string `json:"cursor"`
	Limit  int    `json:"limit"`
}

type GetGroupChatDetailReq struct {
	ChatId string
}
