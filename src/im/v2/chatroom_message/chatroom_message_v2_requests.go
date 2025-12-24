package chatroom_message

// SendChatroomMessageRequestV2 发送聊天室消息请求
type SendChatroomMessageRequestV2 struct {
	RoomId         int64           `json:"-"`                         // 聊天室ID，用于构造URL
	SenderId       string          `json:"sender_id"`                 // 发送者ID
	ReceiverIds    []string        `json:"receiver_ids,omitempty"`    // 接收者ID列表
	ResendFlag     int             `json:"resend_flag,omitempty"`     // 重发标识
	Extension      string          `json:"extension,omitempty"`       // 扩展字段
	Message        *MessageBody    `json:"message"`                   // 消息体
	MessageConfig  *MessageConfig  `json:"message_config,omitempty"`  // 消息配置
	RouteConfig    *RouteConfig    `json:"route_config,omitempty"`    // 路由配置
	AntispamConfig *AntispamConfig `json:"antispam_config,omitempty"` // 反垃圾配置
	AiParams       *AiParams       `json:"ai_params,omitempty"`       // AI参数
}

// BatchSendChatroomMessagesRequestV2 批量发送聊天室消息请求
type BatchSendChatroomMessagesRequestV2 struct {
	RoomId         int64           `json:"room_id"`                   // 聊天室ID
	SenderId       string          `json:"sender_id"`                 // 发送者ID
	ReceiverIds    []string        `json:"receiver_ids,omitempty"`    // 接收者ID列表
	ResendFlag     int             `json:"resend_flag,omitempty"`     // 重发标识
	Extension      string          `json:"extension,omitempty"`       // 扩展字段
	Messages       []MessageBody   `json:"messages"`                  // 消息列表
	MessageConfig  *MessageConfig  `json:"message_config,omitempty"`  // 消息配置
	RouteConfig    *RouteConfig    `json:"route_config,omitempty"`    // 路由配置
	AntispamConfig *AntispamConfig `json:"antispam_config,omitempty"` // 反垃圾配置
}

// RecallChatroomMessageRequestV2 撤回聊天室消息请求
type RecallChatroomMessageRequestV2 struct {
	RoomId                int64  `json:"-"`                                // 聊天室ID，用于构造URL
	MessageClientId       string `json:"-"`                                // 消息客户端ID，用于构造URL
	OperatorId            string `json:"operator_id,omitempty"`            // 操作者ID
	SenderId              string `json:"sender_id,omitempty"`              // 发送者ID
	CreateTime            int64  `json:"create_time,omitempty"`            // 创建时间
	NotificationEnabled   bool   `json:"notification_enabled,omitempty"`   // 是否启用通知
	NotificationExtension string `json:"notification_extension,omitempty"` // 通知扩展字段
}

// QueryChatroomHistoryMessagesRequestV2 查询聊天室历史消息请求
type QueryChatroomHistoryMessagesRequestV2 struct {
	RoomId       int64  `json:"-"` // 聊天室ID，用于构造URL
	SenderId     string `json:"-"` // 发送者ID，用于构造查询参数
	PageToken    string `json:"-"` // 分页token，用于构造查询参数
	Limit        int    `json:"-"` // 限制数量，用于构造查询参数
	Descending   bool   `json:"-"` // 是否降序，用于构造查询参数
	MessageTypes string `json:"-"` // 消息类型，用于构造查询参数
}

// MessageBody 消息体
type MessageBody struct {
	MessageClientId string                 `json:"message_client_id"`    // 消息客户端ID
	MessageType     int                    `json:"message_type"`         // 消息类型
	SubType         int                    `json:"sub_type,omitempty"`   // 子类型
	Text            string                 `json:"text,omitempty"`       // 文本内容
	Attachment      map[string]interface{} `json:"attachment,omitempty"` // 附件
	LocationX       float64                `json:"location_x,omitempty"` // 位置X坐标
	LocationY       float64                `json:"location_y,omitempty"` // 位置Y坐标
	LocationZ       float64                `json:"location_z,omitempty"` // 位置Z坐标
}

// MessageConfig 消息配置
type MessageConfig struct {
	IgnoreChatBanned               bool   `json:"ignore_chat_banned,omitempty"`                  // 是否忽略禁言
	HistoryEnabled                 bool   `json:"history_enabled,omitempty"`                     // 是否启用历史记录
	HighPriority                   bool   `json:"high_priority,omitempty"`                       // 是否高优先级
	NeedHighPriorityMsgResend      bool   `json:"need_high_priority_msg_resend,omitempty"`       // 是否需要高优先级消息重发
	AbandonRatio                   int    `json:"abandon_ratio,omitempty"`                       // 丢弃比例
	NotifyTargetTags               string `json:"notify_target_tags,omitempty"`                  // 通知目标标签
	ChatMsgPriority                int    `json:"chat_msg_priority,omitempty"`                   // 聊天消息优先级
	ForbiddenIfHighPriorityMsgFreq int    `json:"forbidden_if_high_priority_msg_freq,omitempty"` // 高优先级消息频率禁止阈值
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
	AntispamExtension            string `json:"antispam_extension,omitempty"`              // 反垃圾扩展字段
	AntispamCheating             string `json:"antispam_cheating,omitempty"`               // 反垃圾作弊信息
	AntispamCustomMessageEnabled bool   `json:"antispam_custom_message_enabled,omitempty"` // 是否启用自定义反垃圾消息
	AntispamCustomMessage        string `json:"antispam_custom_message,omitempty"`         // 自定义反垃圾消息
}

// AiParams AI参数
type AiParams struct {
	Account         string `json:"account,omitempty"`          // 账号
	Content         string `json:"content,omitempty"`          // 内容
	Messages        string `json:"messages,omitempty"`         // 消息
	PromptVariables string `json:"prompt_variables,omitempty"` // 提示变量
	Config          string `json:"config,omitempty"`           // 配置
	SendOriginal    bool   `json:"send_original,omitempty"`    // 是否发送原始消息
}
