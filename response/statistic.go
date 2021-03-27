package response

type GetUserBehaviorDataResp struct {
	ErrInfo
	BehaviorData []*struct {
		StatTime            int64   `json:"stat_time"`
		ChatCnt             int     `json:"chat_cnt"`
		MessageCnt          int     `json:"message_cnt"`
		ReplyPercentage     float32 `json:"reply_percentage"`
		AvgReplyTime        int     `json:"avg_reply_time"`
		NegativeFeedbackCnt int     `json:"negative_feedback_cnt"`
		NewApplyCnt         int     `json:"new_apply_cnt"`
		NewContactCnt       int     `json:"new_contact_cnt"`
	} `json:"behavior_data"`
}

type GetGroupChatStatisticByGroupOwnerResp struct {
	ErrInfo
	Total      int `json:"total"`
	NextOffset int `json:"next_offset"`
	Items      []*struct {
		Owner string `json:"owner"`
		Data  *struct {
			NewChatCnt   int `json:"new_chat_cnt"`
			ChatTotal    int `json:"chat_total"`
			ChatHasMsg   int `json:"chat_has_msg"`
			NewMemberCnt int `json:"new_member_cnt"`
			MemberTotal  int `json:"member_total"`
			MemberHasMsg int `json:"member_has_msg"`
			MsgTotal     int `json:"msg_total"`
		} `json:"data"`
	} `json:"items"`
}
