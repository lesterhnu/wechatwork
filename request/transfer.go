package request

type TransferCustomerReq struct {
	HandoverUserid     string   `json:"handover_userid"`
	TakeoverUserid     string   `json:"takeover_userid"`
	ExternalUserid     []string `json:"external_userid"`
	TransferSuccessMsg string   `json:"transfer_success_msg"`
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
