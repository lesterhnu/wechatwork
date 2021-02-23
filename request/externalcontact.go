package request

import "github.com/go-playground/validator/v10"

type BatchGetExtContactDetailReq struct {
	UserId string `json:"user_id" validate:"required"`
	Cursor string `json:"cursor"`
	Limit  int    `json:"limit"`
}

func (b *BatchGetExtContactDetailReq) Validate() error {
	return validator.New().Struct(b)
}
