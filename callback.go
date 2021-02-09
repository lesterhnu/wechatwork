package wechatwork

import (
	"errors"
	"fmt"
	"wechatworksdk/request"
	"wechatworksdk/util/wxbizmsgcrypt"
)

func VerifyURL(config *request.CommonConfig, req *request.VerifyURLReq) (string, error) {
	wxcpt := wxbizmsgcrypt.NewWXBizMsgCrypt(config.Token, config.EncodingAESKey, config.ReceiverId, wxbizmsgcrypt.XmlType)
	timestamp := fmt.Sprintf("%d", req.Timestamp)
	echoStr, err := wxcpt.VerifyURL(req.MsgSignature, timestamp, req.Nonce, req.Echostr)
	if err != nil {
		return "", errors.New(err.ErrMsg)
	}
	return string(echoStr), nil
}
