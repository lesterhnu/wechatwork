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
)

const (
	ContentTypeJson = "application/json"
)
