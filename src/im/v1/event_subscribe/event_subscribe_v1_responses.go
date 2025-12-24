package event_subscribe

// Subscription 订阅信息
type Subscription struct {
	Accid         string `json:"accid"`         // 发布者账号ID
	EventType     int    `json:"eventType"`     // 事件类型
	ExpireTime    int64  `json:"expireTime"`    // 过期时间（毫秒时间戳）
	SubscribeTime int64  `json:"subscribeTime"` // 订阅时间（毫秒时间戳）
}

// AddEventSubscribeResponseV1 添加事件订阅响应
type AddEventSubscribeResponseV1 struct {
	FailedAccid []string `json:"failedAccid"` // 订阅失败的账号ID列表
}

// QueryEventSubscribeResponseV1 查询事件订阅响应
type QueryEventSubscribeResponseV1 struct {
	Subscribes []Subscription `json:"subscribes"` // 订阅信息列表
}

// DeleteEventSubscribeResponseV1 删除事件订阅响应
type DeleteEventSubscribeResponseV1 struct {
	FailedAccid []string `json:"failedAccid"` // 取消订阅失败的账号ID列表
}

// BatchDeleteEventSubscribeResponseV1 批量删除所有事件订阅响应
type BatchDeleteEventSubscribeResponseV1 struct {
	// API仅返回成功码，无额外数据
}
