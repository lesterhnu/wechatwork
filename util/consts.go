package util

const (
	// 获取access_token
	GetAccessTokenUrl = "https://qyapi.weixin.qq.com/cgi-bin/gettoken"

	// 添加联系方式
	AddContactWayUrl = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_contact_way"
	// 获取联系方式
	GetContactWayUrl = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_contact_way"
	// 更新联系方式
	UpdateContactWayUrl = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/update_contact_way"
	// 删除联系方式
	DelContactWayUrl = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/del_contact_way"
	// 获取配置了客户联系功能的成员列表
	GetFollowUserList = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_follow_user_list"

	// 获取客户列表
	GetExtContactList = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/list"
	// 获取客户详情
	GetExtContactDetail = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get"
	// 批量获取客户详情
	BatchGetExtContactDetail = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/batch/get_by_user"
	// 修改客户备注信息
	UpdateExtContactRemark = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/remark"

	// 客户标签管理
	// 管理企业客户标签
	// 获取企业客户标签
	GetCorpTagListUrl = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_corp_tag_list"
	// 添加企业客户标签
	AddCorpTagUrl = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_corp_tag"
	// 编辑企业客户标签
	UpdateCorpTagUrl = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/edit_corp_tag"
	// 删除企业客户标签
	DelCorpTagUrl = "https://qyapi.weixin.qq.com/cgi-bin/externalcUrlontact/del_corp_tag"
	// 编辑客户企业标签
	SetExtContactTagUrl = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/mark_tag"

	// 在职继承
	// 分配在职成员的客户
	TransferCustomerUrl = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/transfer_customer"
	// 查询客户接替状态
	TransferResultUrl = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/transfer_result"

	// 离职继承
	// 获取待分配的离职成员列表
	GetUnassignedListUrl = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_unassigned_list"
	// 分配离职成员的客户
	ResignedTransferCustomerUrl = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/resigned/transfer_customer"
	// 查询客户接替状态
	ResignedTransferResultUrl = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/resigned/transfer_result"
	// 分配离职成员的客户群
	ResignedTransferChatGroupUrl = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/transfer"

	// 客户群管理
	// 获取客户群列表
	GetGroupChatListUrl = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/list"
	// 获取客户群详情
	GetGroupChatDetailUrl = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/get"

	// 客户朋友圈
	// 获取朋友圈发表记录
	// GetMomentListUrl = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_moment_list"

	// 消息推送
	// 创建企业群发
	AddMsgTemplateUrl = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_msg_template"
	// 获取群发记录列表
	GetGroupmsgListUrl = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_groupmsg_list"
	// 获取群发成员发送任务列表
	GetGroupmsgTaskUrl = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_groupmsg_task"
	// 获取企业群发成员执行结果
	GetGroupmsgSendResultUrl = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_groupmsg_send_result"
	// 发送新客户欢迎语
	SendWelcomeMsgUrl = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/send_welcome_msg"
	// 添加入群欢迎语素材
	AddGroupWelcomeTemplateUrl = ":https://qyapi.weixin.qq.com/cgi-bin/externalcontact/group_welcome_template/add"
	// 编辑入群欢迎语素材
	EditGroupWelcomeTemplateUrl = ":https://qyapi.weixin.qq.com/cgi-bin/externalcontact/group_welcome_template/edit"
	// 获取入群欢迎语素材
	GetGroupWelcomeTemplateUrl = ":https://qyapi.weixin.qq.com/cgi-bin/externalcontact/group_welcome_template/get"
	// 删除入群欢迎语素材
	DelGroupWelcomeTemplateUrl = ":https://qyapi.weixin.qq.com/cgi-bin/externalcontact/group_welcome_template/del"

	// 统计管理
	// 获取【联系客户统计】数据
	GetUserBehaviorDataUrl = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_user_behavior_data"
	// 获取【群聊数据统计】数据 按群主聚合的方式
	GetGroupChatStatisticByGroupOwnerUrl = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/statistic"
	// 获取【群聊数据统计】数据 按自然日聚合的方式
	GetGroupChatStatisticByDayUrl = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/statistic_group_by_day"
)

const (
	ContentTypeJson = "application/json"
)
