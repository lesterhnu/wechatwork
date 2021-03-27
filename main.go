package wechatwork

import (
	"github.com/go-playground/validator/v10"
	"github.com/lesterhnu/wechatwork/request"
	"github.com/lesterhnu/wechatwork/response"
	"log"
)

type WxCfg struct {
	CorpId                string `validate:"required"`
	ContactToken          string
	ContactEncodingAesKey string
	Secret                string `validate:"required"`
	AccessToken           string
}

type wxClient struct {
	corpId                string
	secret                string
	contactToken          string
	contactEncodingAesKey string
	accessToken           string
}

// 参数校验
func (w *WxCfg) Validate() error {
	return validator.New().Struct(w)
}
func NewWxClient(cfg *WxCfg) (*wxClient, error) {
	err := cfg.Validate()
	if err != nil {
		return nil, err
	}
	client := &wxClient{
		corpId:                cfg.CorpId,
		secret:                cfg.Secret,
		contactToken:          cfg.ContactToken,
		contactEncodingAesKey: cfg.ContactEncodingAesKey,
		accessToken:           cfg.AccessToken,
	}
	if client.accessToken != "" {
		return client, nil
	}
	err = client.getAccessToken()
	if err != nil {
		return nil, err
	}
	return client, nil
}
func (w *wxClient) Test() {
	log.Println("accessToken:", w.accessToken)
	log.Println("corpId", w.corpId)
	log.Println("secret", w.secret)
	log.Println("contactToken", w.contactToken)
}
func (w *wxClient) Token() {
	log.Println("token:", w.accessToken)
}

type WxClientInterface interface {
	// 回调
	// 验证回调url
	VerifyURL(req *request.VerifyURLReq) (string, error)
	// 解密回调信息
	Decrypt(req *request.DecryptReq) ([]byte, error)

	// 企业服务人员管理
	// 获取配置了客户联系功能的成员列表
	GetFollowUserList() (*response.GetFollowUserListResp, error)
	// 配置客户联系「联系我」方式
	AddContact(req *request.AddContactReq) (*response.AddContactResp, error)
	// 获取企业已配置的「联系我」方式
	GetContact(req *request.GetContact) (*response.GetContactWayResp, error)
	// 更新企业已配置的「联系我」方式
	UpdateContact(req request.UpdateContactWay) (*response.UpdateContactWayResp, error)
	// 删除企业已配置的「联系我」方式
	DelContact(req *request.DelContactWay) (*response.DelContactWayResp, error)

	// 客户管理
	// 获取客户列表
	GetContactList(userId string) (*response.GetContactListResp, error)
	// 获取客户详情
	GetExtContactDetail(extUserId string) (*response.GetExtContactDetail, error)
	// 批量获取客户详情
	BatchGetExtContactDetail(req *request.BatchGetExtContactDetailReq) (*response.BatchGetExtContactDetailResp, error)
	// 修改客户备注信息
	UpdateExtContactRemark(req *request.UpdateExtContactRemarkReq) (*response.UpdateExtContactRemarkResp, error)

	// 客户标签管理
	// 获取企业标签库
	GetCorpTagList(req *request.GetCorpTagListReq) (*response.GetCorpTagListResp, error)
	// 添加企业客户标签
	AddCorpTag(req *request.AddCorpTagReq) (*response.AddCorpTagResp, error)
	// 编辑企业客户标签
	UpdateCorpTag(req *request.UpdateCorpTagReq) (*response.UpdateCorpTagResp, error)
	// 删除企业客户标签
	DelCorpTag(req *request.DelCorpTagReq) (*response.DelCorpTagResp, error)
	// 编辑客户企业标签
	SetExtContactTag(req *request.SetExtContactTagReq) (*response.SetExtContactTagResp, error)

	// 在职继承
	// 分配在职成员的客户
	TransferCustomer(req *request.TransferCustomerReq) (*response.TransferCustomerResp, error)
	// 查询客户接替状态
	TransferResult(req *request.TransferResultReq) (*response.TransferResultResp, error)

	// 离职继承
	// 获取待分配的离职成员列表
	GetUnassignedList(req *request.GetUnassignedListReq) (*response.GetUnassignedListResp, error)

	// 分配离职成员的客户
	ResignedTransferCustomer(req *request.ResignedTransferCustomerReq) (*response.ResignedTransferCustomerResp, error)
	// 查询客户接替状态
	GetResignedTransferResult(req *request.GetResignedTransferResultReq) (*response.GetResignedTransferResultResp, error)
	// 分配离职成员的客户群
	ResignedTransferChatGroup(req *request.ResignedTransferChatGroupReq) (*response.ResignedTransferChatGroupResp, error)

	// 获取客户群列表
	GetGroupChatList(req *request.GetGroupChatListReq) (*response.GetGroupChatListResp, error)
	// 获取客户群详情
	GetGroupChatDetail(req *request.GetGroupChatDetailReq) (*response.GetGroupChatDetailResp, error)

	// 消息推送
	// 创建企业群发
	AddMsgTemplate(req *request.AddMsgTemplateReq) (*response.AddMsgTemplateResp, error)
	// 获取群发记录
	GetGroupmsgList(req *request.GetGroupmsgListReq) (*response.GetGroupmsgListResp, error)
	// 获取群发成员发送任务列表
	GetGroupmsgTask(req *request.GetGroupmsgTaskReq) (*response.GetGroupmsgTaskResp, error)

	GetGroupmsgSendResult(req *request.GetGroupmsgSendResultReq) (*response.GetGroupmsgSendResultResp, error)

	SendWelcomeMsg(req *request.SendWelcomeMsgReq) (*response.SendWelcomeMsgResp, error)

	AddGroupWelcomeTemplate(req *request.AddGroupWelcomeTemplateReq) (*response.AddGroupWelcomeTemplateResp, error)

	EditGroupWelcomeTemplate(req *request.EditGroupWelcomeTemplateReq) (*response.EditGroupWelcomeTemplateResp, error)

	GetGroupWelcomeTemplate(req *request.GetGroupWelcomeTemplateReq) (*response.GetGroupWelcomeTemplateResp, error)

	DelGroupWelcomeTemplate(req *request.DelGroupWelcomeTemplateReq) (*response.DelGroupWelcomeTemplateResp, error)

	// 数据统计
	// 获取「联系客户统计」数据
	GetUserBehaviorData(req *request.GetUserBehaviorDataReq) (*response.GetUserBehaviorDataResp, error)
}
