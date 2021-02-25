package response

type TransferCustomerResponse struct {
	ErrInfo
	Customer []*CustomerItem `json:"customer"`
}

type CustomerItem struct {
	ExternalUserid string `json:"external_userid"`
	Errcode        int    `json:"errcode"`
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
