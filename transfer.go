package wechatwork

import (
	"encoding/json"
	"fmt"
	"github.com/lesterhnu/wechatworksdk/request"
	"github.com/lesterhnu/wechatworksdk/response"
	"github.com/lesterhnu/wechatworksdk/util"
)

func (w *wxClient) TransferCustomer(req *request.TransferCustomerReq) (*response.TransferCustomerResp, error) {
	url := fmt.Sprintf("%s?access_token=%s", util.TransferCustomerUrl, w.accessToken)
	content := util.Post(url, req, util.ContentTypeJson)
	var resp = new(response.TransferCustomerResp)
	if err := json.Unmarshal([]byte(content), resp); err != nil {
		return nil, err
	}
	return resp, nil
}
func (w *wxClient) TransferResult(req *request.TransferResultReq) (*response.TransferResultResp, error) {
	url := fmt.Sprintf("%s?access_token=%s", util.TransferResultUrl, w.accessToken)
	content := util.Post(url, req, util.ContentTypeJson)
	var resp = new(response.TransferResultResp)
	if err := json.Unmarshal([]byte(content), resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *wxClient) GetUnassignedList(req *request.GetUnassignedListReq) (*response.GetUnassignedListResp, error) {
	url := fmt.Sprintf("%s?access_token=%s", util.GetUnassignedListUrl, w.accessToken)
	content := util.Post(url, req, util.ContentTypeJson)
	var resp = new(response.GetUnassignedListResp)
	if err := json.Unmarshal([]byte(content), resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *wxClient) ResignedTransferCustomer(req *request.ResignedTransferCustomerReq) (*response.ResignedTransferCustomerResp, error) {
	url := fmt.Sprintf("%s?access_token=%s", util.ResignedTransferCustomerUrl, w.accessToken)
	content := util.Post(url, req, util.ContentTypeJson)
	var resp = new(response.ResignedTransferCustomerResp)
	if err := json.Unmarshal([]byte(content), resp); err != nil {
		return nil, err
	}
	return resp, nil
}
func (w *wxClient) GetResignedTransferResult(req *request.GetResignedTransferResultReq) (*response.GetResignedTransferResultResp, error) {
	url := fmt.Sprintf("%s?access_token=%s", util.ResignedTransferResultUrl, w.accessToken)
	content := util.Post(url, req, util.ContentTypeJson)
	var resp = new(response.GetResignedTransferResultResp)
	if err := json.Unmarshal([]byte(content), resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *wxClient) ResignedTransferChatGroup(req *request.ResignedTransferChatGroupReq) (*response.ResignedTransferChatGroupResp, error) {
	url := fmt.Sprintf("%s?access_token=%s", util.ResignedTransferChatGroupUrl, w.accessToken)
	content := util.Post(url, req, util.ContentTypeJson)
	var resp = new(response.ResignedTransferChatGroupResp)
	if err := json.Unmarshal([]byte(content), resp); err != nil {
		return nil, err
	}
	return resp, nil
}
