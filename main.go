package wechatwork

import (
	"github.com/go-playground/validator/v10"
	"log"
)

type WxCfg struct {
	CorpId                string `validate:"required"`
	ContactToken          string
	ContactEncodingAesKey string
	Secret                string `validate:"required"`
	AccessToken           string
}

type wxClient struct {
	corpId                string
	secret                string
	contactToken          string
	contactEncodingAesKey string
	accessToken           string
}
type WxClientInterface interface {
	Token()
}

// 参数校验
func (w *WxCfg) Validate() error {
	return validator.New().Struct(w)
}
func NewWxClient(cfg *WxCfg) (*wxClient, error) {
	err := cfg.Validate()
	if err != nil {
		return nil, err
	}
	client := &wxClient{
		corpId:                cfg.CorpId,
		secret:                cfg.Secret,
		contactToken:          cfg.ContactToken,
		contactEncodingAesKey: cfg.ContactEncodingAesKey,
		accessToken:           cfg.AccessToken,
	}
	if client.accessToken != "" {
		return client, nil
	}
	err = client.getAccessToken()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (w *wxClient) Test() {
	log.Println("accessToken:", w.accessToken)
	log.Println("corpId", w.corpId)
	log.Println("secret", w.secret)
	log.Println("contactToken", w.contactToken)
}
func (w *wxClient) Token() {
	log.Println("token:", w.accessToken)
}
