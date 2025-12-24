package broadcast

// SendBroadcastNotificationResponseV2 发送广播通知响应
type SendBroadcastNotificationResponseV2 struct {
	BroadcastId    string   `json:"broadcast_id"`              // 广播通知ID
	FromAccountId  string   `json:"from_account_id,omitempty"` // 发送方账号ID
	Content        string   `json:"content,omitempty"`         // 通知内容
	OfflineEnabled bool     `json:"offline_enabled,omitempty"` // 是否启用离线推送
	TargetOs       []string `json:"target_os,omitempty"`       // 目标操作系统列表
	CreateTime     int64    `json:"create_time,omitempty"`     // 创建时间
	ExpireTime     int64    `json:"expire_time,omitempty"`     // 过期时间
}

// DeleteBroadcastNotificationResponseV2 删除广播通知响应
type DeleteBroadcastNotificationResponseV2 struct {
	// Success response will only contain status code 200
	// No additional fields are needed for this response class
}

// QueryBroadcastNotificationResponseV2 查询广播通知响应
type QueryBroadcastNotificationResponseV2 struct {
	BroadcastId    string   `json:"broadcast_id"`              // 广播通知ID
	FromAccountId  string   `json:"from_account_id,omitempty"` // 发送方账号ID
	Content        string   `json:"content,omitempty"`         // 通知内容
	OfflineEnabled bool     `json:"offline_enabled,omitempty"` // 是否启用离线推送
	TargetOs       []string `json:"target_os,omitempty"`       // 目标操作系统列表
	CreateTime     int64    `json:"create_time,omitempty"`     // 创建时间
	ExpireTime     int64    `json:"expire_time,omitempty"`     // 过期时间
}

// QueryBroadcastNotificationListResponseV2 查询广播通知列表响应
type QueryBroadcastNotificationListResponseV2 struct {
	HasMore   bool                        `json:"has_more,omitempty"`   // 是否有更多数据
	NextToken string                      `json:"next_token,omitempty"` // 下一页标记
	Items     []BroadcastNotificationItem `json:"items,omitempty"`      // 广播通知列表
}

// BroadcastNotificationItem 广播通知项
type BroadcastNotificationItem struct {
	BroadcastId    string   `json:"broadcast_id"`              // 广播通知ID
	FromAccountId  string   `json:"from_account_id,omitempty"` // 发送方账号ID
	Content        string   `json:"content,omitempty"`         // 通知内容
	OfflineEnabled bool     `json:"offline_enabled,omitempty"` // 是否启用离线推送
	TargetOs       []string `json:"target_os,omitempty"`       // 目标操作系统列表
	CreateTime     int64    `json:"create_time,omitempty"`     // 创建时间
	ExpireTime     int64    `json:"expire_time,omitempty"`     // 过期时间
}

// SendChatroomBroadcastNotificationResponseV2 发送聊天室广播通知响应
type SendChatroomBroadcastNotificationResponseV2 struct {
	CreateTime        int64                  `json:"create_time,omitempty"`         // 创建时间
	ClientBroadcastId string                 `json:"client_broadcast_id,omitempty"` // 客户端广播ID
	SenderId          string                 `json:"sender_id,omitempty"`           // 发送方ID
	FromNick          string                 `json:"from_nick,omitempty"`           // 发送方昵称
	FromAvatar        string                 `json:"from_avatar,omitempty"`         // 发送方头像
	MessageType       int                    `json:"message_type,omitempty"`        // 消息类型
	Text              string                 `json:"text,omitempty"`                // 文本内容
	Attachment        map[string]interface{} `json:"attachment,omitempty"`          // 消息附件
	SenderClientType  int                    `json:"sender_client_type,omitempty"`  // 发送方客户端类型
	HighPriority      bool                   `json:"high_priority,omitempty"`       // 是否高优先级
	SubType           int                    `json:"sub_type,omitempty"`            // 子类型
}
