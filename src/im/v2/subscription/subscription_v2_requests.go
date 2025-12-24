package subscription

// SubscribeUserStatusRequestV2 订阅用户状态请求
type SubscribeUserStatusRequestV2 struct {
	SubscriberAccountId string   `json:"-"`           // 订阅者账号ID,路径参数
	Duration            int64    `json:"-"`           // 订阅有效期(秒),查询参数
	AccountIds          []string `json:"account_ids"` // 要订阅的账号ID列表
}

// UnsubscribeUserStatusRequestV2 取消订阅用户状态请求
type UnsubscribeUserStatusRequestV2 struct {
	SubscriberAccountId string   `json:"-"`           // 订阅者账号ID,路径参数
	AccountIds          []string `json:"account_ids"` // 要取消订阅的账号ID列表
}

// QueryUserStatusSubscriptionRequestV2 查询用户状态订阅请求
type QueryUserStatusSubscriptionRequestV2 struct {
	SubscriberAccountId string   `json:"-"`           // 订阅者账号ID,路径参数
	AccountIds          []string `json:"account_ids"` // 要查询订阅的账号ID列表
}
