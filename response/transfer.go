package response

type TransferCustomerResponse struct {
	ErrInfo
}
type Customer struct {
}
type CustomerItem struct {
	ExternalUserid string `json:"external_userid"`
	Errcode        int    `json:"errcode"`
}
