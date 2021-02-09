package request

type VerfiyUrlReq struct {
	MsgSignature string `json:"msg_signature"`
	Timestamp    int64  `json:"timestamp"`
	Nonce        string `json:"nonce"`
	Echostr      string `json:"echostr"`
}

type DecryptReq struct {
	MsgSignature string `json:"msg_signature"`
	Timestamp    int64  `json:"timestamp"`
	Nonce        string `json:"nonce"`
	Echostr      string `json:"echostr"`
}
type ReceiveMsg struct {
	ToUserName string `xml:"ToUserName"`
	AgentId    string `xml:"AgentId"`
	Encrypt    string `xml:"Encrypt"`
}
