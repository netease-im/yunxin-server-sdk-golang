package broadcast

// SendBroadcastNotificationRequestV2 发送广播通知请求
type SendBroadcastNotificationRequestV2 struct {
	Content        string   `json:"content"`                   // 通知内容
	FromAccountId  string   `json:"from_account_id,omitempty"` // 发送方账号ID
	OfflineEnabled bool     `json:"offline_enabled,omitempty"` // 是否启用离线推送
	Ttl            int      `json:"ttl,omitempty"`             // 消息存活时间（秒）
	TargetOs       []string `json:"target_os,omitempty"`       // 目标操作系统列表
}

// DeleteBroadcastNotificationRequestV2 删除广播通知请求
type DeleteBroadcastNotificationRequestV2 struct {
	BroadcastId string `json:"-"` // 广播通知ID，用于构造URL
}

// QueryBroadcastNotificationRequestV2 查询广播通知请求
type QueryBroadcastNotificationRequestV2 struct {
	BroadcastId string `json:"-"` // 广播通知ID，用于构造URL
}

// QueryBroadcastNotificationListRequestV2 查询广播通知列表请求
type QueryBroadcastNotificationListRequestV2 struct {
	PageToken string `json:"page_token,omitempty"` // 分页标记
	Limit     int    `json:"limit,omitempty"`      // 分页大小，默认100
	Type      int    `json:"type,omitempty"`       // 类型：1-未推送，2-已推送，3-全部，默认1
}

// SendChatroomBroadcastNotificationRequestV2 发送聊天室广播通知请求
type SendChatroomBroadcastNotificationRequestV2 struct {
	ClientBroadcastId string          `json:"client_broadcast_id"`       // 客户端生成的广播唯一标识
	SenderId          string          `json:"sender_id"`                 // 发送方账号ID
	Extension         string          `json:"extension,omitempty"`       // 自定义扩展字段
	ResendFlag        int             `json:"resend_flag,omitempty"`     // 重发标记（0-否，1-是），默认0
	Message           *Message        `json:"message"`                   // 消息内容
	MessageConfig     *MessageConfig  `json:"message_config,omitempty"`  // 消息配置
	RouteConfig       *RouteConfig    `json:"route_config,omitempty"`    // 路由配置
	AntispamConfig    *AntispamConfig `json:"antispam_config,omitempty"` // 反垃圾配置
}

// Message 广播消息内容
type Message struct {
	MessageType int                    `json:"message_type"`         // 消息类型（如：0-文本，1-图片，100-自定义）
	SubType     int                    `json:"sub_type,omitempty"`   // 自定义消息子类型
	Text        string                 `json:"text,omitempty"`       // 文本消息内容
	Attachment  map[string]interface{} `json:"attachment,omitempty"` // 消息附件，用于复杂消息类型
}

// MessageConfig 广播消息配置
type MessageConfig struct {
	HighPriority     bool   `json:"high_priority,omitempty"`      // 是否高优先级
	NotifyTargetTags string `json:"notify_target_tags,omitempty"` // 目标通知标签
}

// RouteConfig 路由配置
type RouteConfig struct {
	RouteEnabled     bool   `json:"route_enabled,omitempty"`     // 是否启用路由
	RouteEnvironment string `json:"route_environment,omitempty"` // 路由环境
}

// AntispamConfig 反垃圾配置
type AntispamConfig struct {
	AntispamEnabled              bool   `json:"antispam_enabled,omitempty"`                // 是否启用反垃圾
	AntispamBusinessId           string `json:"antispam_business_id,omitempty"`            // 反垃圾业务ID
	AntispamExtension            string `json:"antispam_extension,omitempty"`              // 反垃圾自定义扩展
	AntispamCheating             string `json:"antispam_cheating,omitempty"`               // 反垃圾作弊检测字段
	AntispamCustomMessageEnabled bool   `json:"antispam_custom_message_enabled,omitempty"` // 是否使用自定义消息用于反垃圾结果
	AntispamCustomMessage        string `json:"antispam_custom_message,omitempty"`         // 反垃圾结果自定义消息
}
