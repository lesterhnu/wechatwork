package wechatwork

import (
	"errors"
	"fmt"
	"github.com/lesterhnu/wechatwork/request"
	"github.com/lesterhnu/wechatwork/util/wxbizmsgcrypt"
)

func (w *wxClient) VerifyURL(req *request.VerifyURLReq) (string, error) {
	wxcpt := wxbizmsgcrypt.NewWXBizMsgCrypt(w.contactToken, w.contactEncodingAesKey, w.corpId, wxbizmsgcrypt.XmlType)
	timestamp := fmt.Sprintf("%d", req.Timestamp)
	echoStr, err := wxcpt.VerifyURL(req.MsgSignature, timestamp, req.Nonce, req.Echostr)
	if err != nil {
		return "", errors.New(err.ErrMsg)
	}
	return string(echoStr), nil
}
func (w *wxClient) Decrypt(req *request.DecryptReq) ([]byte, error) {
	wxcpt := wxbizmsgcrypt.NewWXBizMsgCrypt(w.contactToken, w.contactEncodingAesKey, w.corpId, wxbizmsgcrypt.XmlType)
	timestamp := fmt.Sprintf("%d", req.Timestamp)
	msg, err := wxcpt.DecryptMsg(req.MsgSignature, timestamp, req.Nonce, []byte(req.EchoStr))
	if err != nil {
		return nil, errors.New(err.ErrMsg)
	}
	return msg, nil
}
