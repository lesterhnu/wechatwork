package wechatwork

import (
	"encoding/json"
	"fmt"
	"wechatworksdk/request"
	"wechatworksdk/response"
	"wechatworksdk/util"
)

func (w *wxClient) GetContactList(userId string) (*response.GetContactListResp, error) {
	url := fmt.Sprintf("%s?access_token=%s&userid=%s", util.GetExtContactList, w.accessToken, userId)
	content := util.Get(url)
	var resp = new(response.GetContactListResp)
	if err := json.Unmarshal([]byte(content), resp); err != nil {
		return nil, err
	}
	return resp, nil
}
func (w *wxClient) GetExtContactDetail(extUserId string) (*response.GetExtContactDetail, error) {
	url := fmt.Sprintf("%s?access_token=%s&userid=%s", util.GetExtContactDetail, w.accessToken, extUserId)
	content := util.Get(url)
	var resp = new(response.GetExtContactDetail)
	if err := json.Unmarshal([]byte(content), resp); err != nil {
		return nil, err
	}
	return resp, nil
}
func (w *wxClient) BatchGetExtContactDetail(req *request.BatchGetExtContactDetailReq) (*response.BatchGetExtContactDetailResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s?access_token=%s", util.BatchGetExtContactDetail, w.accessToken)
	content := util.Post(url, req, util.ContentTypeJson)
	var resp = new(response.BatchGetExtContactDetailResp)
	if err := json.Unmarshal([]byte(content), resp); err != nil {
		return nil, err
	}
	return resp, nil
}
