package wechatwork

import (
	"encoding/json"
	"fmt"
	"github.com/lesterhnu/wechatwork/response"
	"github.com/lesterhnu/wechatwork/util"
	"github.com/muesli/cache2go"
	"log"
	"time"
)

func (w *wxClient) getAccessToken() error {
	err := w.getAccessTokenFromCache()
	if err == nil {
		log.Println(err)
		return nil
	}
	url := fmt.Sprintf("%s?corpid=%s&corpsecret=%s", util.GetAccessTokenUrl, w.corpId, w.secret)
	content := util.Get(url)
	var resp = new(response.GetAccessTokenResp)
	if err := json.Unmarshal([]byte(content), resp); err != nil {
		return err
	}
	w.accessToken = resp.AccessToken
	w.setAccessTokenToCache()
	return nil
}
func (w *wxClient) getAccessTokenFromCache() error {
	cache := cache2go.Cache("accessToken")
	cacheItem, err := cache.Value(w.secret)
	if err != nil {
		return err
	}
	w.accessToken = cacheItem.Data().(string)
	return nil
}
func (w *wxClient) setAccessTokenToCache() {
	cache := cache2go.Cache("accessToken")
	cache.Add(w.secret, time.Second*7200, w.accessToken)
}
