package wechatwork

import (
	"encoding/json"
	"fmt"
	"wechatworksdk/request"
	"wechatworksdk/response"
	"wechatworksdk/util"
)

func GetAccessToken(req *request.GetAccessTokenReq) (*response.GetAccessTokenResp, error) {
	url := fmt.Sprintf("%s?corpid=%s&corpsecret=%s", util.GetAccessTokenUrl, req.CorpId, req.CorpSecret)
	content := util.Get(url)
	var resp = new(response.GetAccessTokenResp)
	if err := json.Unmarshal([]byte(content), resp); err != nil {
		return nil, err
	}
	return resp, nil
}
