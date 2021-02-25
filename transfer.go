package wechatwork

import (
	"encoding/json"
	"fmt"
	"github.com/lesterhnu/wechatworksdk/request"
	"github.com/lesterhnu/wechatworksdk/response"
	"github.com/lesterhnu/wechatworksdk/util"
)

func (w *wxClient) TransferCustomer(req *request.TransferCustomerReq) (*response.TransferCustomerResponse, error) {
	url := fmt.Sprintf("%s?access_token=%s", util.TransferCustomerUrl, w.accessToken)
	content := util.Post(url, req, util.ContentTypeJson)
	var resp = new(response.TransferCustomerResponse)
	if err := json.Unmarshal([]byte(content), resp); err != nil {
		return nil, err
	}
	return resp, nil
}
