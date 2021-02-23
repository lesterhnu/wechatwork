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
)

const (
	ContentTypeJson = "application/json"
)
