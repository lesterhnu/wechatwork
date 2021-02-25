package request

type TransferCustomerReq struct {
	HandoverUserid     string   `json:"handover_userid"`
	TakeoverUserid     string   `json:"takeover_userid"`
	ExternalUserid     []string `json:"external_userid"`
	TransferSuccessMsg string   `json:"transfer_success_msg"`
}
