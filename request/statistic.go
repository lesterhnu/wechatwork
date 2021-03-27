package request

type GetUserBehaviorDataReq struct {
	UserId    []string `json:"user_id"`
	Partyid   []int32  `json:"partyid"`
	StartTime int64    `json:"start_time"`
	EndTime   int64    `json:"end_time"`
}

type GetGroupChatStatisticByGroupOwnerReq struct {
	DayBeginTime int64 `json:"day_begin_time"`
	DayEndTime   int64 `json:"day_end_time"`
	OwnerFilter  *struct {
		UseridList []string `json:"userid_list"`
	} `json:"owner_filter"`
	OrderBy  int `json:"order_by"`
	OrderAsc int `json:"order_asc"`
	Offset   int `json:"offset"`
	Limit    int `json:"limit"`
}
