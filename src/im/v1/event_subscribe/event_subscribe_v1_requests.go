package event_subscribe

// AddEventSubscribeRequestV1 添加事件订阅请求
type AddEventSubscribeRequestV1 struct {
	Accid           string   `json:"accid,omitempty"`           // 事件订阅者账号ID
	EventType       *int     `json:"eventType,omitempty"`       // 事件类型，固定为1表示在线状态事件
	PublisherAccids []string `json:"publisherAccids,omitempty"` // 要订阅的发布者账号ID列表
	Ttl             *int64   `json:"ttl,omitempty"`             // 订阅有效期（秒），60-2592000，即60秒到30天
}

// QueryEventSubscribeRequestV1 查询事件订阅请求
type QueryEventSubscribeRequestV1 struct {
	Accid           string   `json:"accid,omitempty"`           // 事件订阅者账号ID
	EventType       *int     `json:"eventType,omitempty"`       // 事件类型，固定为1表示在线状态事件
	PublisherAccids []string `json:"publisherAccids,omitempty"` // 要查询订阅状态的发布者账号ID列表
}

// DeleteEventSubscribeRequestV1 删除特定事件订阅请求
type DeleteEventSubscribeRequestV1 struct {
	Accid           string   `json:"accid,omitempty"`           // 事件订阅者账号ID
	EventType       *int     `json:"eventType,omitempty"`       // 事件类型，固定为1表示在线状态事件
	PublisherAccids []string `json:"publisherAccids,omitempty"` // 要取消订阅的发布者账号ID列表
}

// BatchDeleteEventSubscribeRequestV1 批量删除所有事件订阅请求
type BatchDeleteEventSubscribeRequestV1 struct {
	Accid     string `json:"accid,omitempty"`     // 事件订阅者账号ID
	EventType *int   `json:"eventType,omitempty"` // 事件类型，固定为1表示在线状态事件
}
