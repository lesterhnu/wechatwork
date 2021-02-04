package response

type GetAccessTokenResp struct {
	ErrInfo
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
