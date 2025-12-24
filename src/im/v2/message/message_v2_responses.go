package message

// Message 消息信息
type Message struct {
	MessageServerId string      `json:"message_server_id"` // 消息服务端ID
	ConversationId  string      `json:"conversation_id"`   // 会话ID
	FromAccountId   string      `json:"from_account_id"`   // 发送者账号ID
	MsgType         int         `json:"msg_type"`          // 消息类型
	Body            interface{} `json:"body"`              // 消息体
	Timestamp       int64       `json:"timestamp"`         // 时间戳
	ServerExtension string      `json:"server_extension"`  // 服务端扩展字段
}

// SendMessageResponseV2 发送消息响应
type SendMessageResponseV2 struct {
	Message
}

// BatchSendP2PMessageResponseV2 批量发送点对点消息响应
type BatchSendP2PMessageResponseV2 struct {
	SuccessList []Message    `json:"success_list"` // 成功列表
	FailedList  []FailedSend `json:"failed_list"`  // 失败列表
}

// FailedSend 发送失败信息
type FailedSend struct {
	AccountId string `json:"account_id"` // 账号ID
	ErrorCode int    `json:"error_code"` // 错误码
	ErrorMsg  string `json:"error_msg"`  // 错误信息
}

// ModifyMessageResponseV2 修改消息响应
type ModifyMessageResponseV2 struct {
	Message
}

// WithdrawMessageResponseV2 撤回消息响应
type WithdrawMessageResponseV2 struct {
	// API返回的data为空对象 {}，无需任何字段
}

// DeleteConversationMessagesResponseV2 删除会话所有消息响应
type DeleteConversationMessagesResponseV2 struct {
	// API返回的data为空对象 {}，无需任何字段
}

// SendP2PReadReceiptResponseV2 发送点对点消息已读回执响应
type SendP2PReadReceiptResponseV2 struct {
	Timestamp int64 `json:"timestamp"` // 时间戳
}

// SendTeamReadReceiptResponseV2 发送群消息已读回执响应
type SendTeamReadReceiptResponseV2 struct {
	SuccessList []string        `json:"success_list"` // 成功的消息ID列表
	FailedList  []FailedReceipt `json:"failed_list"`  // 失败列表
}

// FailedReceipt 回执失败信息
type FailedReceipt struct {
	MessageServerId string `json:"message_server_id"` // 消息服务端ID
	ErrorCode       int    `json:"error_code"`        // 错误码
	ErrorMsg        string `json:"error_msg"`         // 错误信息
}

// QueryTeamReadReceiptResponseV2 查询群消息已读回执详情响应
type QueryTeamReadReceiptResponseV2 struct {
	ReadCount      int      `json:"read_count"`      // 已读人数
	UnreadCount    int      `json:"unread_count"`    // 未读人数
	ReadAccounts   []string `json:"read_accounts"`   // 已读账号列表
	UnreadAccounts []string `json:"unread_accounts"` // 未读账号列表
}

// StreamMessageResponseV2 流式消息响应
type StreamMessageResponseV2 struct {
	Message
}

// QueryMessageResponseV2 查询单条消息响应
type QueryMessageResponseV2 struct {
	Message
}

// SearchMessagesResponseV2 搜索历史消息响应
type SearchMessagesResponseV2 struct {
	Messages []Message `json:"messages"` // 消息列表
}

// QueryMessagesByPageResponseV2 分页查询会话消息响应
type QueryMessagesByPageResponseV2 struct {
	Messages []Message `json:"messages"` // 消息列表
}

// BatchQueryMessagesByIdResponseV2 批量查询消息响应
type BatchQueryMessagesByIdResponseV2 struct {
	SuccessList []Message     `json:"success_list"` // 成功列表
	FailedList  []FailedQuery `json:"failed_list"`  // 失败列表
}

// FailedQuery 查询失败信息
type FailedQuery struct {
	MessageServerId string `json:"message_server_id"` // 消息服务端ID
	ErrorCode       int    `json:"error_code"`        // 错误码
	ErrorMsg        string `json:"error_msg"`         // 错误信息
}

// QueryThreadMessagesResponseV2 查询Thread消息响应
type QueryThreadMessagesResponseV2 struct {
	Messages []Message `json:"messages"` // 消息列表
}

// AddQuickCommentResponseV2 添加快捷评论响应
type AddQuickCommentResponseV2 struct {
	CommentId string `json:"comment_id"` // 评论ID
	Timestamp int64  `json:"timestamp"`  // 时间戳
}

// DeleteQuickCommentResponseV2 删除快捷评论响应
type DeleteQuickCommentResponseV2 struct {
	Timestamp int64 `json:"timestamp"` // 时间戳
}

// BatchQueryQuickCommentsResponseV2 批量查询快捷评论响应
type BatchQueryQuickCommentsResponseV2 struct {
	SuccessList []QuickCommentDetail `json:"success_list"` // 成功列表
	FailedList  []FailedCommentQuery `json:"failed_list"`  // 失败列表
}

// QuickCommentDetail 快捷评论详情
type QuickCommentDetail struct {
	ConversationId  string    `json:"conversation_id"`   // 会话ID
	MessageServerId string    `json:"message_server_id"` // 消息服务端ID
	Comments        []Comment `json:"comments"`          // 评论列表
}

// Comment 评论信息
type Comment struct {
	CommentIdx      int    `json:"comment_idx"`      // 评论索引
	AccountId       string `json:"account_id"`       // 账号ID
	Timestamp       int64  `json:"timestamp"`        // 时间戳
	ServerExtension string `json:"server_extension"` // 服务端扩展字段
}

// FailedCommentQuery 评论查询失败信息
type FailedCommentQuery struct {
	ConversationId  string `json:"conversation_id"`   // 会话ID
	MessageServerId string `json:"message_server_id"` // 消息服务端ID
	ErrorCode       int    `json:"error_code"`        // 错误码
	ErrorMsg        string `json:"error_msg"`         // 错误信息
}
