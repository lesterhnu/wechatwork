package request

import (
	"github.com/go-playground/validator/v10"
)

type BatchGetExtContactDetailReq struct {
	UserId string `json:"user_id" validate:"required"`
	Cursor string `json:"cursor"`
	Limit  int    `json:"limit"`
}

func (b *BatchGetExtContactDetailReq) Validate() error {
	return validator.New().Struct(b)
}

type UpdateExtContactRemarkReq struct {
	UserId           string   `json:"user_id" validate:"required"`
	ExternalUserId   string   `json:"external_user_id" validate:"required"`
	Remark           string   `json:"remark"`
	Description      string   `json:"description"`
	RemarkCompany    string   `json:"remark_company"`
	RemarkMobiles    []string `json:"remark_mobiles"`
	RemarkPicMediaid string   `json:"remark_pic_mediaid"`
}

func (b *UpdateExtContactRemarkReq) Validate() error {
	return validator.New().Struct(b)
}
