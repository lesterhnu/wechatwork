package response

type ErrInfo struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}
