package wechatwork

import (
	"encoding/json"
	"fmt"
	"github.com/lesterhnu/wechatworksdk/request"
	"github.com/lesterhnu/wechatworksdk/response"
	"github.com/lesterhnu/wechatworksdk/util"
)

func (w *wxClient) AddContact(req *request.AddContactReq) (*response.AddContactResp, error) {
	url := fmt.Sprintf("%s?access_token=%s", util.AddContactWayUrl, w.accessToken)
	content := util.Post(url, req, util.ContentTypeJson)
	var resp = new(response.AddContactResp)
	if err := json.Unmarshal([]byte(content), resp); err != nil {
		return nil, err
	}
	return resp, nil
}
func (w *wxClient) GetContact(req *request.GetContact) (*response.GetContactWayResp, error) {
	url := fmt.Sprintf("%s?access_token=%s", util.GetContactWayUrl, w.accessToken)
	content := util.Post(url, req, util.ContentTypeJson)
	var resp = new(response.GetContactWayResp)
	if err := json.Unmarshal([]byte(content), resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *wxClient) DelContact(req *request.DelContactWay) (*response.DelContactWayResp, error) {
	url := fmt.Sprintf("%s?access_token=%s", util.DelContactWayUrl, w.accessToken)
	content := util.Post(url, req, util.ContentTypeJson)
	var resp = new(response.DelContactWayResp)
	if err := json.Unmarshal([]byte(content), resp); err != nil {
		return nil, err
	}
	return resp, nil
}
func (w *wxClient) UpdateContact(req request.UpdateContactWay) (*response.UpdateContactWayResp, error) {
	url := fmt.Sprintf("%s?access_token=%s", util.UpdateContactWayUrl, w.accessToken)
	content := util.Post(url, req, util.ContentTypeJson)
	var resp = new(response.UpdateContactWayResp)
	if err := json.Unmarshal([]byte(content), resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *wxClient) GetFollowUserList() (*response.GetFollowUserListResp, error) {
	url := fmt.Sprintf("%s?access_token=%s", util.GetFollowUserList, w.accessToken)
	content := util.Get(url)
	var resp = new(response.GetFollowUserListResp)
	if err := json.Unmarshal([]byte(content), resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *wxClient) GetGroupChatList(req *request.GetGroupChatListReq) (*response.GetGroupChatListResp, error) {
	url := fmt.Sprintf("%s?access_token=%s", util.GetGroupChatListUrl, w.accessToken)
	content := util.Post(url, req, util.ContentTypeJson)
	var resp = new(response.GetGroupChatListResp)
	if err := json.Unmarshal([]byte(content), resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *wxClient) GetGroupChatDetail(req *request.GetGroupChatDetailReq) (*response.GetGroupChatDetailResp, error) {
	url := fmt.Sprintf("%s?access_token=%s", util.GetGroupChatDetailUrl, w.accessToken)
	content := util.Post(url, req, util.ContentTypeJson)
	var resp = new(response.GetGroupChatDetailResp)
	if err := json.Unmarshal([]byte(content), resp); err != nil {
		return nil, err
	}
	return resp, nil
}
