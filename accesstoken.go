package wechatwork

import (
	"encoding/json"
	"fmt"
	"github.com/muesli/cache2go"
	"time"
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

func (w *wxClient) GetAccessToken() (*wxClient, error) {
	url := fmt.Sprintf("%s?corpid=%s&corpsecret=%s", util.GetAccessTokenUrl, w.corpId, w.secret)
	content := util.Get(url)
	var resp = new(response.GetAccessTokenResp)
	if err := json.Unmarshal([]byte(content), resp); err != nil {
		return nil, err
	}
	w.accessToken = resp.AccessToken
	return w, nil
}
func (w *wxClient) getAccessTokenFromCache() (*wxClient, error) {
	cache := cache2go.Cache("accessToken")
	cacheItem, err := cache.Value(w.secret)
	if err != nil {
		return nil, err
	}
	w.accessToken = cacheItem.Data().(string)
	return w, nil
}
func (w *wxClient) setAccessTokenToCache() {
	cache := cache2go.Cache("accessToken")
	cache.Add(w.secret, time.Second*7200, w.accessToken)
}
