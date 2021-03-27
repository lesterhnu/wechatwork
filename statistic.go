package wechatwork

import (
	"encoding/json"
	"fmt"
	"github.com/lesterhnu/wechatwork/request"
	"github.com/lesterhnu/wechatwork/response"
	"github.com/lesterhnu/wechatwork/util"
)

func (w *wxClient) GetUserBehaviorData(req *request.GetUserBehaviorDataReq) (*response.GetUserBehaviorDataResp, error) {
	url := fmt.Sprintf("%s?access_token=%s", util.GetUserBehaviorDataUrl, w.accessToken)
	content := util.Post(url, req, util.ContentTypeJson)
	var resp = new(response.GetUserBehaviorDataResp)
	if err := json.Unmarshal([]byte(content), resp); err != nil {
		return nil, err
	}
	return resp, nil
}
