package message

// 消息体
type MessageBody struct {
	MsgType int         `json:"msg_type"`         // 消息类型
	Text    string      `json:"text,omitempty"`   // 文本消息内容
	Body    interface{} `json:"body,omitempty"`   // 其他类型消息体
	Attach  interface{} `json:"attach,omitempty"` // 自定义扩展字段
}

// 消息配置
type MessageConfig struct {
	CheckFriend     bool   `json:"check_friend"`               // 是否需要好友关系
	CheckMute       bool   `json:"check_mute"`                 // 是否需要检查消息免打扰
	Offline         bool   `json:"offline"`                    // 消息是否存离线
	Roam            bool   `json:"roam"`                       // 消息是否需要漫游
	UnreadEnabled   bool   `json:"unread_enabled"`             // 消息是否计入未读数
	HistoryEnabled  bool   `json:"history_enabled"`            // 消息是否存云端历史
	PersistEnabled  bool   `json:"persist_enabled"`            // 消息是否需要持久化
	MsgExtension    string `json:"msg_extension,omitempty"`    // 消息扩展字段
	ServerExtension string `json:"server_extension,omitempty"` // 服务端扩展字段
	ClientMsgId     string `json:"client_msg_id,omitempty"`    // 客户端消息ID
	ResendFlag      int    `json:"resend_flag"`                // 重发标识
}

// SendMessageRequestV2 发送消息请求
type SendMessageRequestV2 struct {
	ConversationId string         `json:"-"`                        // 会话ID,路径参数
	Message        MessageBody    `json:"message"`                  // 消息内容
	MessageConfig  *MessageConfig `json:"message_config,omitempty"` // 消息配置
}

// BatchSendP2PMessageRequestV2 批量发送点对点消息请求
type BatchSendP2PMessageRequestV2 struct {
	FromAccountId string         `json:"from_account_id"`          // 发送者账号ID
	ToAccountIds  []string       `json:"to_account_ids"`           // 接收者账号ID列表
	Message       MessageBody    `json:"message"`                  // 消息内容
	MessageConfig *MessageConfig `json:"message_config,omitempty"` // 消息配置
}

// ModifyMessageRequestV2 修改消息请求
type ModifyMessageRequestV2 struct {
	ConversationId  string      `json:"conversation_id"`            // 会话ID
	MessageServerId string      `json:"message_server_id"`          // 消息服务端ID
	Body            interface{} `json:"body"`                       // 修改后的消息体
	ServerExtension string      `json:"server_extension,omitempty"` // 服务端扩展字段
}

// WithdrawMessageRequestV2 撤回消息请求
type WithdrawMessageRequestV2 struct {
	ConversationId  string `json:"-"` // 会话ID,路径参数
	MessageServerId string `json:"-"` // 消息服务端ID,路径参数
	Type            int    `json:"-"` // 撤回类型,查询参数: 1=单向撤回, 2=双向撤回, 3=删除消息
	ServerExtension string `json:"-"` // 服务端扩展字段,查询参数
	Notify          string `json:"-"` // 撤回通知,查询参数
}

// DeleteConversationMessagesRequestV2 删除会话所有消息请求
type DeleteConversationMessagesRequestV2 struct {
	ConversationId string `json:"-"` // 会话ID,路径参数
	DeleteType     int    `json:"-"` // 删除类型,查询参数: 1=删除漫游, 2=删除历史, 3=删除漫游和历史
}

// SendP2PReadReceiptRequestV2 发送点对点消息已读回执请求
type SendP2PReadReceiptRequestV2 struct {
	SendAccountId    string `json:"send_account_id"`    // 发送者账号ID
	ReceiveAccountId string `json:"receive_account_id"` // 接收者账号ID
	MessageServerId  string `json:"message_server_id"`  // 消息服务端ID
}

// SendTeamReadReceiptRequestV2 发送群消息已读回执请求
type SendTeamReadReceiptRequestV2 struct {
	SendAccountId    string   `json:"send_account_id"`    // 发送者账号ID
	TeamId           string   `json:"team_id"`            // 群组ID
	MessageServerIds []string `json:"message_server_ids"` // 消息服务端ID列表
}

// QueryTeamReadReceiptRequestV2 查询群消息已读回执详情请求
type QueryTeamReadReceiptRequestV2 struct {
	SendAccountId   string `json:"-"` // 发送者账号ID,查询参数
	TeamId          string `json:"-"` // 群组ID,查询参数
	MessageServerId string `json:"-"` // 消息服务端ID,查询参数
	IncludeAccount  bool   `json:"-"` // 是否包含账号列表,查询参数
}

// StreamMessageRequestV2 流式消息请求
type StreamMessageRequestV2 struct {
	ConversationId string      `json:"-"`       // 会话ID,路径参数
	Message        MessageBody `json:"message"` // 消息内容
}

// QueryMessageRequestV2 查询单条消息请求
type QueryMessageRequestV2 struct {
	ConversationId  string `json:"-"` // 会话ID,路径参数
	MessageServerId string `json:"-"` // 消息服务端ID,路径参数
}

// SearchMessagesRequestV2 搜索历史消息请求
type SearchMessagesRequestV2 struct {
	SendAccountId  string `json:"-"` // 发送者账号ID,查询参数
	ConversationId string `json:"-"` // 会话ID,查询参数
	MsgType        int    `json:"-"` // 消息类型,查询参数
	Keyword        string `json:"-"` // 关键词,查询参数
	BeginTime      int64  `json:"-"` // 开始时间,查询参数
	EndTime        int64  `json:"-"` // 结束时间,查询参数
	Limit          int    `json:"-"` // 返回条数限制,查询参数
	Reverse        bool   `json:"-"` // 是否倒序,查询参数
}

// QueryMessagesByPageRequestV2 分页查询会话消息请求
type QueryMessagesByPageRequestV2 struct {
	ConversationId string `json:"-"` // 会话ID,路径参数
	BeginTime      int64  `json:"-"` // 开始时间,查询参数
	EndTime        int64  `json:"-"` // 结束时间,查询参数
	Limit          int    `json:"-"` // 返回条数限制,查询参数
	Reverse        bool   `json:"-"` // 是否倒序,查询参数
	MsgTypes       string `json:"-"` // 消息类型列表,查询参数
}

// BatchQueryMessagesByIdRequestV2 批量查询消息请求
type BatchQueryMessagesByIdRequestV2 struct {
	ConversationId   string   `json:"-"`                  // 会话ID,路径参数
	MessageServerIds []string `json:"message_server_ids"` // 消息服务端ID列表
}

// QueryThreadMessagesRequestV2 查询Thread消息请求
type QueryThreadMessagesRequestV2 struct {
	ThreadRoot     string `json:"-"` // Thread根消息,查询参数
	ConversationId string `json:"-"` // 会话ID,查询参数
	BeginTime      int64  `json:"-"` // 开始时间,查询参数
	EndTime        int64  `json:"-"` // 结束时间,查询参数
	Limit          int    `json:"-"` // 返回条数限制,查询参数
	Reverse        bool   `json:"-"` // 是否倒序,查询参数
	ExcludeRoot    bool   `json:"-"` // 是否排除根消息,查询参数
}

// AddQuickCommentRequestV2 添加快捷评论请求
type AddQuickCommentRequestV2 struct {
	OperatorAccountId string `json:"operator_account_id"`        // 操作者账号ID
	ConversationId    string `json:"conversation_id"`            // 会话ID
	MessageServerId   string `json:"message_server_id"`          // 消息服务端ID
	CommentIdx        int    `json:"comment_idx"`                // 评论索引
	ServerExtension   string `json:"server_extension,omitempty"` // 服务端扩展字段
	PushEnabled       bool   `json:"push_enabled"`               // 是否推送
	NeedBadge         bool   `json:"need_badge"`                 // 是否需要角标
}

// DeleteQuickCommentRequestV2 删除快捷评论请求
type DeleteQuickCommentRequestV2 struct {
	OperatorAccountId string `json:"-"` // 操作者账号ID,查询参数
	ConversationId    string `json:"-"` // 会话ID,查询参数
	MessageServerId   string `json:"-"` // 消息服务端ID,查询参数
	CommentIdx        int    `json:"-"` // 评论索引,查询参数
}

// BatchQueryQuickCommentsRequestV2 批量查询快捷评论请求
type BatchQueryQuickCommentsRequestV2 struct {
	Messages []QuickCommentMessage `json:"messages"` // 消息列表
}

// QuickCommentMessage 快捷评论消息
type QuickCommentMessage struct {
	ConversationId  string `json:"conversation_id"`   // 会话ID
	MessageServerId string `json:"message_server_id"` // 消息服务端ID
}
