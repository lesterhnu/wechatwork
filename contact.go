package wechatwork

import (
	"encoding/json"
	"fmt"
	"wechatworksdk/request"
	"wechatworksdk/response"
	"wechatworksdk/util"
)

func AddContact(accessToken string, req *request.AddContactReq) (*response.AddContactResp, error) {
	url := fmt.Sprintf("%s?access_token=%s", util.AddContactWayUrl, accessToken)
	content := util.Post(url, req, util.ContentTypeJson)
	var resp = new(response.AddContactResp)
	if err := json.Unmarshal([]byte(content), resp); err != nil {
		return nil, err
	}
	return resp, nil
}
func GetContact(accessToken string, req *request.GetContact) (*response.GetContactWayResp, error) {
	url := fmt.Sprintf("%s?access_token=%s", util.GetContactWayUrl, accessToken)
	content := util.Post(url, req, util.ContentTypeJson)
	var resp = new(response.GetContactWayResp)
	if err := json.Unmarshal([]byte(content), resp); err != nil {
		return nil, err
	}
	return resp, nil
}
