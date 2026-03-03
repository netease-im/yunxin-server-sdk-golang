package message

// MessageBody 消息体
type MessageBody struct {
	MessageType     *int        `json:"message_type,omitempty"`      // 消息类型: 0=文本, 1=图片, 2=音频, 3=视频, 4=地理位置, 6=文件, 10=提示, 100=自定义
	SubType         *int        `json:"sub_type,omitempty"`          // 自定义消息子类型,仅当message_type=100时有效
	Text            string      `json:"text,omitempty"`              // 消息文本内容
	Attachment      interface{} `json:"attachment,omitempty"`        // 消息附件
	MessageClientId string      `json:"message_client_id,omitempty"` // 客户端消息ID
}

// MessageConfig 消息配置
type MessageConfig struct {
	UnreadEnabled             *bool `json:"unread_enabled,omitempty"`              // 是否计入未读数
	MutilSyncEnabled          *bool `json:"mutil_sync_enabled,omitempty"`          // 是否同步到发送者的多端
	OfflineEnabled            *bool `json:"offline_enabled,omitempty"`             // 是否存离线
	HistoryEnabled            *bool `json:"history_enabled,omitempty"`             // 是否存云端历史
	RoamingEnabled            *bool `json:"roaming_enabled,omitempty"`             // 是否漫游
	ConversationUpdateEnabled *bool `json:"conversation_update_enabled,omitempty"` // 是否更新会话列表最后一条消息
}

// TargetOption 群定向消息配置
type TargetOption struct {
	ReceiverAccountIds   []string `json:"receiver_account_ids,omitempty"`    // 接收者账号ID列表
	Inclusive            *bool    `json:"inclusive,omitempty"`               // true=白名单(仅列表中可见), false=黑名单(列表中不可见)
	CheckTeamMemberValid *bool    `json:"check_team_member_valid,omitempty"` // 是否检查列表中的账号是否为有效群成员
	VisibleToNewMember   *bool    `json:"visible_to_new_member,omitempty"`   // 新成员是否可见此消息
}

// RouteConfig 消息路由配置
type RouteConfig struct {
	RouteEnabled     *bool  `json:"route_enabled,omitempty"`     // 是否转发到指定应用服务器
	RouteEnvironment string `json:"route_environment,omitempty"` // 路由环境名称
}

// PushConfig 推送配置
type PushConfig struct {
	PushEnabled          *bool    `json:"push_enabled,omitempty"`           // 是否推送
	PushNickEnabled      *bool    `json:"push_nick_enabled,omitempty"`      // 推送内容是否包含发送者昵称
	PushContent          string   `json:"push_content,omitempty"`           // 自定义推送内容
	PushPayload          string   `json:"push_payload,omitempty"`           // 自定义推送payload,JSON格式
	PushForcepushAll     *bool    `json:"push_forcepush_all,omitempty"`     // 是否强推给所有群成员(@所有人)
	PushForcepushIds     []string `json:"push_forcepush_ids,omitempty"`     // 强推账号列表(@指定人)
	PushForcepushContent string   `json:"push_forcepush_content,omitempty"` // 强推自定义内容
	PushForcepushEnable  *bool    `json:"push_forcepush_enable,omitempty"`  // 是否开启强推(@操作)
}

// AntispamConfig 反垃圾配置
type AntispamConfig struct {
	AntispamEnabled              *bool  `json:"antispam_enabled,omitempty"`                // 是否开启反垃圾
	AntispamBusinessId           string `json:"antispam_business_id,omitempty"`            // 反垃圾业务ID
	AntispamExtension            string `json:"antispam_extension,omitempty"`              // 反垃圾扩展参数
	AntispamCheating             string `json:"antispam_cheating,omitempty"`               // 反垃圾作弊参数
	AntispamCustomMessageEnabled *bool  `json:"antispam_custom_message_enabled,omitempty"` // 是否检测自定义消息
	AntispamCustomMessage        string `json:"antispam_custom_message,omitempty"`         // 待检测的自定义消息内容
}

// P2POption 点对点消息配置
type P2POption struct {
	CheckFriend *bool `json:"check_friend,omitempty"` // 是否检查好友关系
}

// TeamOption 高级群消息配置
type TeamOption struct {
	MarkAsRead             *bool `json:"mark_as_read,omitempty"`              // 是否需要已读回执
	IgnoreChatBanned       *bool `json:"ignore_chat_banned,omitempty"`        // 是否忽略群禁言
	IgnoreMemberChatBanned *bool `json:"ignore_member_chat_banned,omitempty"` // 是否忽略成员禁言
	CheckTeamMemberValid   *bool `json:"check_team_member_valid,omitempty"`   // 是否检查发送者是否为有效群成员
}

// SuperTeamOption 超大群消息配置
type SuperTeamOption struct {
	IgnoreChatBanned       *bool `json:"ignore_chat_banned,omitempty"`        // 是否忽略群禁言
	IgnoreMemberChatBanned *bool `json:"ignore_member_chat_banned,omitempty"` // 是否忽略成员禁言
	CheckTeamMemberValid   *bool `json:"check_team_member_valid,omitempty"`   // 是否检查发送者是否为有效群成员
}

// RobotConfig 机器人配置
type RobotConfig struct {
	RobotAccountId     string `json:"robot_account_id,omitempty"`     // 机器人账号ID
	RobotTopic         string `json:"robot_topic,omitempty"`          // 机器人消息主题
	RobotFunction      string `json:"robot_function,omitempty"`       // 机器人功能
	RobotCustomContent string `json:"robot_custom_content,omitempty"` // 机器人自定义内容
}

// ThreadConfig Thread消息配置
type ThreadConfig struct {
	ThreadRoot  *ThreadMessage `json:"thread_root,omitempty"`  // Thread根消息信息
	ThreadReply *ThreadMessage `json:"thread_reply,omitempty"` // 回复消息信息
}

// ThreadMessage Thread消息信息
type ThreadMessage struct {
	SenderId        string `json:"sender_id,omitempty"`         // 发送者ID
	ReceiverId      string `json:"receiver_id,omitempty"`       // 接收者ID
	CreateTime      int64  `json:"create_time,omitempty"`       // 消息创建时间
	MessageServerId int64  `json:"message_server_id,omitempty"` // 服务端消息ID
	MessageClientId string `json:"message_client_id,omitempty"` // 客户端消息ID
}

// AIParams AI数字人配置参数
type AIParams struct {
	Account         string `json:"account,omitempty"`          // 账号
	Content         string `json:"content,omitempty"`          // 内容
	Messages        string `json:"messages,omitempty"`         // 消息
	PromptVariables string `json:"prompt_variables,omitempty"` // 提示变量
	Config          string `json:"config,omitempty"`           // 配置
	SendOriginal    *bool  `json:"send_original,omitempty"`    // 是否发送原始消息
}

// SendMessageRequestV2 发送消息请求
type SendMessageRequestV2 struct {
	ConversationId  string           `json:"-"`                           // 会话ID,路径参数
	Message         *MessageBody     `json:"message"`                     // 消息内容
	MessageConfig   *MessageConfig   `json:"message_config,omitempty"`    // 消息配置
	TargetOption    *TargetOption    `json:"target_option,omitempty"`     // 群定向消息配置
	RouteConfig     *RouteConfig     `json:"route_config,omitempty"`      // 消息路由配置
	PushConfig      *PushConfig      `json:"push_config,omitempty"`       // 推送配置
	AntispamConfig  *AntispamConfig  `json:"antispam_config,omitempty"`   // 反垃圾配置
	P2pOption       *P2POption       `json:"p2p_option,omitempty"`        // 点对点消息配置
	TeamOption      *TeamOption      `json:"team_option,omitempty"`       // 高级群消息配置
	SuperteamOption *SuperTeamOption `json:"superteam_option,omitempty"`  // 超大群消息配置
	RobotConfig     *RobotConfig     `json:"robot_config,omitempty"`      // 机器人配置
	ThreadConfig    *ThreadConfig    `json:"thread_config,omitempty"`     // Thread消息配置
	SenderNoSense   *bool            `json:"sender_no_sense,omitempty"`   // 发送者是否无感知
	ReceiverNoSense *bool            `json:"receiver_no_sense,omitempty"` // 接收者是否无感知
	Extension       string           `json:"extension,omitempty"`         // 扩展信息,JSON格式
	AiParams        *AIParams        `json:"ai_params,omitempty"`         // AI数字人配置参数
}

// StreamAttachment 流式消息附件
type StreamAttachment struct {
	Text   string `json:"text,omitempty"`   // 文本内容
	Index  *int   `json:"index,omitempty"`  // 索引
	Finish *int   `json:"finish,omitempty"` // 是否结束: 0=未结束, 1=结束
}

// StreamMessage 流式消息
type StreamMessage struct {
	MessageClientId string            `json:"message_client_id,omitempty"` // 客户端消息ID
	MessageType     *int              `json:"message_type,omitempty"`      // 消息类型
	Attachment      *StreamAttachment `json:"attachment,omitempty"`        // 附件
}

// RagInfo RAG信息
type RagInfo struct {
	Name        string `json:"name,omitempty"`        // 名称
	Icon        string `json:"icon,omitempty"`        // 图标URL
	Title       string `json:"title,omitempty"`       // 标题
	Time        int64  `json:"time,omitempty"`        // 时间戳
	Url         string `json:"url,omitempty"`         // URL
	Description string `json:"description,omitempty"` // 描述
}

// StreamMessageRequestV2 流式消息请求
type StreamMessageRequestV2 struct {
	ConversationId string         `json:"-"`                       // 会话ID,路径参数
	Message        *StreamMessage `json:"message"`                 // 消息内容
	RagInfoList    []RagInfo      `json:"rag_info_list,omitempty"` // RAG信息列表
	TeamOption     *TeamOption    `json:"team_option,omitempty"`   // 高级群消息配置
	TargetOption   *TargetOption  `json:"target_option,omitempty"` // 群定向消息配置
}

// BatchMessage 批量发送消息体
type BatchMessage struct {
	MessageType     *int        `json:"message_type,omitempty"`      // 消息类型
	SubType         *int        `json:"sub_type,omitempty"`          // 自定义消息子类型
	Text            string      `json:"text,omitempty"`              // 消息文本内容
	Attachment      interface{} `json:"attachment,omitempty"`        // 消息附件
	MessageClientId string      `json:"message_client_id,omitempty"` // 客户端消息ID
}

// BatchSendP2PMessageRequestV2 批量发送点对点消息请求
type BatchSendP2PMessageRequestV2 struct {
	SenderId          string          `json:"sender_id"`                     // 发送者账号ID
	ReceiverIds       []string        `json:"receiver_ids"`                  // 接收者账号ID列表
	Message           *BatchMessage   `json:"message"`                       // 消息内容
	MessageConfig     *MessageConfig  `json:"message_config,omitempty"`      // 消息配置
	PushConfig        *PushConfig     `json:"push_config,omitempty"`         // 推送配置
	AntispamConfig    *AntispamConfig `json:"antispam_config,omitempty"`     // 反垃圾配置
	NeedMessageDetail *bool           `json:"need_message_detail,omitempty"` // 是否需要消息详情
}

// ModifyMessageRequestV2 修改消息请求
type ModifyMessageRequestV2 struct {
	Operator  string      `json:"operator"`            // 操作者账号ID (required)
	Type      *int        `json:"type"`                // 会话类型: 1=P2P, 2=高级群, 3=超大群 (required)
	Extension string      `json:"extension,omitempty"` // 开发者扩展字段
	Message   interface{} `json:"message"`             // 消息内容 (required)
}

// WithdrawMessageRequestV2 撤回消息请求
type WithdrawMessageRequestV2 struct {
	ConversationId  string `json:"-"` // 会话ID,路径参数
	MessageServerId string `json:"-"` // 消息服务端ID,路径参数
	Type            *int   `json:"-"` // 撤回类型,查询参数: 1=单向撤回, 2=双向撤回, 3=删除消息
	ServerExtension string `json:"-"` // 服务端扩展字段,查询参数
	Notify          string `json:"-"` // 撤回通知,查询参数
}

// DeleteConversationMessagesRequestV2 删除会话所有消息请求
type DeleteConversationMessagesRequestV2 struct {
	ConversationId string `json:"-"` // 会话ID,路径参数
	DeleteType     *int   `json:"-"` // 删除类型,查询参数: 1=删除漫游, 2=删除历史, 3=删除漫游和历史
}

// SendP2PReadReceiptRequestV2 发送点对点消息已读回执请求
type SendP2PReadReceiptRequestV2 struct {
	FromAccountId   string `json:"from_account_id"`   // 发送者账号ID(回执发送者/消息接收者)
	ToAccountId     string `json:"to_account_id"`     // 接收者账号ID(回执接收者/消息发送者)
	MessageServerId int64  `json:"message_server_id"` // 消息服务端ID
}

// SendTeamReadReceiptRequestV2 发送群消息已读回执请求
type SendTeamReadReceiptRequestV2 struct {
	FromAccountId    string  `json:"from_account_id"`    // 发送者账号ID
	TeamId           int64   `json:"team_id"`            // 群组ID
	MessageServerIds []int64 `json:"message_server_ids"` // 消息服务端ID列表
}

// QueryTeamReadReceiptRequestV2 查询群消息已读回执详情请求
type QueryTeamReadReceiptRequestV2 struct {
	TeamId          int64 `json:"-"` // 群组ID,查询参数
	MessageServerId int64 `json:"-"` // 消息服务端ID,查询参数
	IncludeAccount  *bool `json:"-"` // 是否包含账号列表,查询参数
}

// QueryMessageRequestV2 查询单条消息请求
type QueryMessageRequestV2 struct {
	ConversationId  string `json:"-"` // 会话ID,路径参数
	MessageServerId int64  `json:"-"` // 消息服务端ID,路径参数
}

// SearchMessagesRequestV2 搜索历史消息请求
type SearchMessagesRequestV2 struct {
	OperatorId       string   `json:"-"` // 操作者账号ID,查询参数 (required)
	ConversationId   string   `json:"-"` // 会话ID,查询参数
	KeywordList      []string `json:"-"` // 关键词列表,查询参数
	SenderAccountIds string   `json:"-"` // 发送者账号ID,查询参数
	MessageTypes     string   `json:"-"` // 消息类型列表,查询参数
	MessageSubTypes  string   `json:"-"` // 消息子类型列表,查询参数
	KeywordMatchType *int     `json:"-"` // 关键词匹配类型,查询参数: 0=OR, 1=AND
	StartTime        int64    `json:"-"` // 开始时间,查询参数
	TimePeriod       int64    `json:"-"` // 时间段,查询参数
	Limit            *int     `json:"-"` // 返回条数限制,查询参数
	PageToken        string   `json:"-"` // 分页token,查询参数
}

// QueryMessagesByPageRequestV2 分页查询会话消息请求
type QueryMessagesByPageRequestV2 struct {
	ConversationId string `json:"-"` // 会话ID,路径参数
	BeginTime      int64  `json:"-"` // 开始时间,查询参数
	EndTime        int64  `json:"-"` // 结束时间,查询参数
	Limit          int    `json:"-"` // 返回条数限制,查询参数
	Descending     *bool  `json:"-"` // 是否降序,查询参数
	MessageType    string `json:"-"` // 消息类型,查询参数
	PageToken      string `json:"-"` // 分页token,查询参数
}

// MessageQuery 消息查询条件
type MessageQuery struct {
	MessageServerId int64 `json:"message_server_id"` // 消息服务端ID
	CreateTime      int64 `json:"create_time"`       // 创建时间
}

// BatchQueryMessagesByIdRequestV2 批量查询消息请求
type BatchQueryMessagesByIdRequestV2 struct {
	ConversationId string         `json:"-"`        // 会话ID,路径参数
	Messages       []MessageQuery `json:"messages"` // 消息列表
}

// QueryThreadMessagesRequestV2 查询Thread消息请求
type QueryThreadMessagesRequestV2 struct {
	BeginTime           int64  `json:"-"` // 开始时间,查询参数
	EndTime             int64  `json:"-"` // 结束时间,查询参数
	Limit               int    `json:"-"` // 返回条数限制,查询参数
	ConversationType    int    `json:"-"` // 会话类型,查询参数: 1=P2P, 2=高级群
	SenderId            string `json:"-"` // 发送者ID,查询参数
	ReceiverId          string `json:"-"` // 接收者ID,查询参数
	RootMessageId       int64  `json:"-"` // 根消息ID,查询参数
	RootMessageClientId string `json:"-"` // 根消息客户端ID,查询参数
	RootMessageTime     int64  `json:"-"` // 根消息时间,查询参数
	Descending          *bool  `json:"-"` // 是否降序,查询参数
	ExcludeRoot         *bool  `json:"-"` // 是否排除根消息,查询参数
	PageToken           string `json:"-"` // 分页token,查询参数
}

// QuickCommentMessage 快捷评论消息信息
type QuickCommentMessage struct {
	ConversationType int    `json:"conversation_type"` // 会话类型: 1=P2P, 2=高级群, 3=超大群
	SenderId         string `json:"sender_id"`         // 发送者ID
	ReceiverId       string `json:"receiver_id"`       // 接收者ID
	MessageServerId  int64  `json:"message_server_id"` // 消息服务端ID
	MessageClientId  string `json:"message_client_id"` // 消息客户端ID
	CreateTime       int64  `json:"create_time"`       // 创建时间
}

// AddQuickCommentRequestV2 添加快捷评论请求
type AddQuickCommentRequestV2 struct {
	OperatorId string               `json:"operator_id"`           // 操作者账号ID (required)
	Index      int64                `json:"index"`                 // 快捷评论类型,必须>0 (required)
	Extension  string               `json:"extension,omitempty"`   // 扩展字段,JSON格式
	Message    *QuickCommentMessage `json:"message"`               // 消息对象 (required)
	PushConfig *PushConfig          `json:"push_config,omitempty"` // 推送配置
}

// DeleteQuickCommentRequestV2 删除快捷评论请求
type DeleteQuickCommentRequestV2 struct {
	OperatorId string               `json:"operator_id"`           // 操作者账号ID (required)
	Index      int64                `json:"index"`                 // 快捷评论类型,必须>0 (required)
	Extension  string               `json:"extension,omitempty"`   // 扩展字段,JSON格式
	Message    *QuickCommentMessage `json:"message"`               // 消息对象 (required)
	PushConfig *PushConfig          `json:"push_config,omitempty"` // 推送配置
}

// BatchQueryQuickCommentsRequestV2 批量查询快捷评论请求
type BatchQueryQuickCommentsRequestV2 struct {
	Messages []QuickCommentMessage `json:"messages"` // 消息列表 (required)
}
