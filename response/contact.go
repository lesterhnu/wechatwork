package response

type AddContactResp struct {
	ErrInfo
	ConfigId string `json:"config_id"`
	QrCode   string `json:"qr_code"`
}

type DelContactWayResp struct {
}
