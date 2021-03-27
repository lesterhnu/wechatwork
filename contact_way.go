package wechatwork

import (
	"encoding/json"
	"fmt"
	"github.com/lesterhnu/wechatwork/request"
	"github.com/lesterhnu/wechatwork/response"
	"github.com/lesterhnu/wechatwork/util"
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

// 创建企业群发
func (w *wxClient) AddMsgTemplate(req *request.AddMsgTemplateReq) (*response.AddMsgTemplateResp, error) {
	url := fmt.Sprintf("%s?access_token=%s", util.AddMsgTemplateUrl, w.accessToken)
	content := util.Post(url, req, util.ContentTypeJson)
	var resp = new(response.AddMsgTemplateResp)
	if err := json.Unmarshal([]byte(content), resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// 获取群发记录
func (w *wxClient) GetGroupmsgList(req *request.GetGroupmsgListReq) (*response.GetGroupmsgListResp, error) {
	url := fmt.Sprintf("%s?access_token=%s", util.GetGroupmsgListUrl, w.accessToken)
	content := util.Post(url, req, util.ContentTypeJson)
	var resp = new(response.GetGroupmsgListResp)
	if err := json.Unmarshal([]byte(content), resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *wxClient) GetGroupmsgTask(req *request.GetGroupmsgTaskReq) (*response.GetGroupmsgTaskResp, error) {
	url := fmt.Sprintf("%s?access_token=%s", util.GetGroupmsgTaskUrl, w.accessToken)
	content := util.Post(url, req, util.ContentTypeJson)
	var resp = new(response.GetGroupmsgTaskResp)
	if err := json.Unmarshal([]byte(content), resp); err != nil {
		return nil, err
	}
	return resp, nil
}
func (w *wxClient) GetGroupmsgSendResult(req *request.GetGroupmsgSendResultReq) (*response.GetGroupmsgSendResultResp, error) {
	url := fmt.Sprintf("%s?access_token=%s", util.GetGroupmsgSendResultUrl, w.accessToken)
	content := util.Post(url, req, util.ContentTypeJson)
	var resp = new(response.GetGroupmsgSendResultResp)
	if err := json.Unmarshal([]byte(content), resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *wxClient) SendWelcomeMsg(req *request.SendWelcomeMsgReq) (*response.SendWelcomeMsgResp, error) {
	url := fmt.Sprintf("%s?access_token=%s", util.SendWelcomeMsgUrl, w.accessToken)
	content := util.Post(url, req, util.ContentTypeJson)
	var resp = new(response.SendWelcomeMsgResp)
	if err := json.Unmarshal([]byte(content), resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *wxClient) AddGroupWelcomeTemplate(req *request.AddGroupWelcomeTemplateReq) (*response.AddGroupWelcomeTemplateResp, error) {
	url := fmt.Sprintf("%s?access_token=%s", util.AddGroupWelcomeTemplateUrl, w.accessToken)
	content := util.Post(url, req, util.ContentTypeJson)
	var resp = new(response.AddGroupWelcomeTemplateResp)
	if err := json.Unmarshal([]byte(content), resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *wxClient) EditGroupWelcomeTemplate(req *request.EditGroupWelcomeTemplateReq) (*response.EditGroupWelcomeTemplateResp, error) {
	url := fmt.Sprintf("%s?access_token=%s", util.EditGroupWelcomeTemplateUrl, w.accessToken)
	content := util.Post(url, req, util.ContentTypeJson)
	var resp = new(response.EditGroupWelcomeTemplateResp)
	if err := json.Unmarshal([]byte(content), resp); err != nil {
		return nil, err
	}
	return resp, nil
}
func (w *wxClient) GetGroupWelcomeTemplate(req *request.GetGroupWelcomeTemplateReq) (*response.GetGroupWelcomeTemplateResp, error) {
	url := fmt.Sprintf("%s?access_token=%s", util.GetGroupWelcomeTemplateUrl, w.accessToken)
	content := util.Post(url, req, util.ContentTypeJson)
	var resp = new(response.GetGroupWelcomeTemplateResp)
	if err := json.Unmarshal([]byte(content), resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *wxClient) DelGroupWelcomeTemplate(req *request.DelGroupWelcomeTemplateReq) (*response.DelGroupWelcomeTemplateResp, error) {
	url := fmt.Sprintf("%s?access_token=%s", util.DelGroupWelcomeTemplateUrl, w.accessToken)
	content := util.Post(url, req, util.ContentTypeJson)
	var resp = new(response.DelGroupWelcomeTemplateResp)
	if err := json.Unmarshal([]byte(content), resp); err != nil {
		return nil, err
	}
	return resp, nil
}
