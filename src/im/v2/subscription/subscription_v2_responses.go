package subscription

// FailedSubscription 订阅失败详情
type FailedSubscription struct {
	AccountId string `json:"account_id"` // 账号ID
	ErrorCode int    `json:"error_code"` // 错误码
	ErrorMsg  string `json:"error_msg"`  // 错误信息
}

// SubscribeUserStatusResponseV2 订阅用户状态响应
type SubscribeUserStatusResponseV2 struct {
	FailedList []FailedSubscription `json:"failed_list"` // 订阅失败列表
}

// FailedUnsubscription 取消订阅失败详情
type FailedUnsubscription struct {
	AccountId string `json:"account_id"` // 账号ID
	ErrorCode int    `json:"error_code"` // 错误码
	ErrorMsg  string `json:"error_msg"`  // 错误信息
}

// UnsubscribeUserStatusResponseV2 取消订阅用户状态响应
type UnsubscribeUserStatusResponseV2 struct {
	FailedList []FailedUnsubscription `json:"failed_list"` // 取消订阅失败列表
}

// FailedQuery 查询失败详情
type FailedQuery struct {
	AccountId string `json:"account_id"` // 账号ID
	ErrorCode int    `json:"error_code"` // 错误码
	ErrorMsg  string `json:"error_msg"`  // 错误信息
}

// SubscriptionDetail 订阅详情
type SubscriptionDetail struct {
	AccountId     string `json:"account_id"`     // 账号ID
	EventType     int    `json:"event_type"`     // 事件类型
	SubscribeTime int64  `json:"subscribe_time"` // 订阅时间
	ExpireTime    int64  `json:"expire_time"`    // 过期时间
}

// QueryUserStatusSubscriptionResponseV2 查询用户状态订阅响应
type QueryUserStatusSubscriptionResponseV2 struct {
	FailedList []FailedQuery        `json:"failed_list"` // 查询失败列表
	Subscribes []SubscriptionDetail `json:"subscribes"`  // 订阅详情列表
}
