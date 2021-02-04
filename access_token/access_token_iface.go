package access_token

import (
	"wechatworksdk/request"
	"wechatworksdk/response"
)

type AccessTokenIFace interface {
	GetAccessToken(req *request.GetAccessTokenReq) (*response.GetAccessTokenResp, error)
}
