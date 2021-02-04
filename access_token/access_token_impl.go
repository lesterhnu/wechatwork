package access_token

import (
	"encoding/json"
	"fmt"
	"wechatworksdk/request"
	"wechatworksdk/response"
	"wechatworksdk/util"
)

type accessToken struct {
}

func NewAccessToken() *accessToken {
	return &accessToken{}
}
func (c *accessToken) GetAccessToken(req *request.GetAccessTokenReq) (*response.GetAccessTokenResp, error) {
	url := fmt.Sprintf("%s?corpid=%s&corpsecret=%s", util.GET_ACCESS_TOKEN_URL, req.CorpId, req.CorpSecret)
	content := util.Get(url)
	var resp = new(response.GetAccessTokenResp)
	if err := json.Unmarshal([]byte(content), resp); err != nil {
		return nil, err
	}
	return resp, nil
}
