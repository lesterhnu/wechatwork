package request

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

type GetCorpTagListReq struct {
	TagId   []string `json:"tag_id"`
	GroupId []string `json:"group_id"`
}

type AddCorpTagReq struct {
	GroupId   string     `json:"group_id"`
	GroupName string     `json:"group_name"`
	Order     int        `json:"order"`
	Tag       []*TagItem `json:"tag"`
	AgentId   int        `json:"agent_id"`
}
type TagItem struct {
	Name  string `json:"name" validate:"required"`
	Order int    `json:"order"`
}

func (a *AddCorpTagReq) Validate() error {
	return validator.New().Struct(a)
}

type UpdateCorpTagReq struct {
	Id    string `json:"id" validate:"required"`
	Name  string `json:"name"`
	Order int    `json:"order"`
	Agent int    `json:"agent"`
}

func (c *UpdateCorpTagReq) Validate() error {
	return validator.New().Struct(c)
}

type DelCorpTagReq struct {
	TagId   []string `json:"tag_id"`
	GroupId []string `json:"group_id"`
	Agent   int      `json:"agent"`
}

func (c *DelCorpTagReq) Validate() error {
	if len(c.TagId) == 0 && len(c.GroupId) == 0 {
		return errors.New("tag_id和group_id不可同时为空")
	}
	return nil
}

type SetExtContactTagReq struct {
	UserId         string   `json:"user_id" validate:"required"`
	ExternalUserId string   `json:"external_user_id" validate:"required"`
	AddTag         []string `json:"add_tag"`
	RemoveTag      []string `json:"remove_tag"`
}

func (c *SetExtContactTagReq) Validate() error {
	if len(c.AddTag) == 0 && len(c.RemoveTag) == 0 {
		return errors.New("tag_id和group_id不可同时为空")
	}
	return validator.New().Struct(c)
}
