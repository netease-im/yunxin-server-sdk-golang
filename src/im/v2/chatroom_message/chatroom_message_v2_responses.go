package chatroom_message

// SendChatroomMessageResponseV2 发送聊天室消息响应
type SendChatroomMessageResponseV2 struct {
	Message    *MessageResponseBody `json:"message,omitempty"`     // 消息体
	FailedList []FailedInfo         `json:"failed_list,omitempty"` // 失败列表
}

// BatchSendChatroomMessagesResponseV2 批量发送聊天室消息响应
type BatchSendChatroomMessagesResponseV2 struct {
	SuccessList []SuccessMessage `json:"success_list,omitempty"` // 成功列表
	FailedList  []FailMessage    `json:"failed_list,omitempty"`  // 失败列表
}

// RecallChatroomMessageResponseV2 撤回聊天室消息响应
type RecallChatroomMessageResponseV2 struct {
	// Success response will only contain status code 200
}

// QueryChatroomHistoryMessagesResponseV2 查询聊天室历史消息响应
type QueryChatroomHistoryMessagesResponseV2 struct {
	HasMore   bool          `json:"has_more,omitempty"`   // 是否有更多数据
	NextToken string        `json:"next_token,omitempty"` // 下一页token
	Items     []MessageItem `json:"items,omitempty"`      // 消息列表
}

// MessageResponseBody 消息响应体
type MessageResponseBody struct {
	Text            string                 `json:"text,omitempty"`              // 文本内容
	MessageClientId string                 `json:"message_client_id,omitempty"` // 消息客户端ID
	CreateTime      int64                  `json:"create_time,omitempty"`       // 创建时间
	MessageType     int                    `json:"message_type,omitempty"`      // 消息类型
	SenderId        string                 `json:"sender_id,omitempty"`         // 发送者ID
	SenderNick      string                 `json:"sender_nick,omitempty"`       // 发送者昵称
	SenderAvatar    string                 `json:"sender_avatar,omitempty"`     // 发送者头像
	RoomId          int64                  `json:"room_id,omitempty"`           // 聊天室ID
	Attachment      map[string]interface{} `json:"attachment,omitempty"`        // 附件
	Extension       string                 `json:"extension,omitempty"`         // 扩展字段
	MsgAbandon      bool                   `json:"msg_abandon,omitempty"`       // 消息是否被丢弃
}

// FailedInfo 失败信息
type FailedInfo struct {
	ReceiverId string `json:"receiver_id,omitempty"` // 接收者ID
	ErrorCode  int    `json:"error_code,omitempty"`  // 错误码
	ErrorMsg   string `json:"error_msg,omitempty"`   // 错误信息
}

// SuccessMessage 成功消息
type SuccessMessage struct {
	MessageClientId string                 `json:"message_client_id,omitempty"` // 消息客户端ID
	CreateTime      int64                  `json:"create_time,omitempty"`       // 创建时间
	MessageType     int                    `json:"message_type,omitempty"`      // 消息类型
	SenderId        string                 `json:"sender_id,omitempty"`         // 发送者ID
	SenderNick      string                 `json:"sender_nick,omitempty"`       // 发送者昵称
	SenderAvatar    string                 `json:"sender_avatar,omitempty"`     // 发送者头像
	RoomId          int64                  `json:"room_id,omitempty"`           // 聊天室ID
	MsgAbandon      bool                   `json:"msg_abandon,omitempty"`       // 消息是否被丢弃
	Text            string                 `json:"text,omitempty"`              // 文本内容
	Attachment      map[string]interface{} `json:"attachment,omitempty"`        // 附件
	Extension       string                 `json:"extension,omitempty"`         // 扩展字段
}

// FailMessage 失败消息
type FailMessage struct {
	MessageClientId string `json:"message_client_id,omitempty"` // 消息客户端ID
	ErrorCode       int    `json:"error_code,omitempty"`        // 错误码
	ErrorMsg        string `json:"error_msg,omitempty"`         // 错误信息
}

// MessageItem 消息项
type MessageItem struct {
	SenderId         string                 `json:"sender_id,omitempty"`          // 发送者ID
	MessageType      int                    `json:"message_type,omitempty"`       // 消息类型
	CreateTime       int64                  `json:"create_time,omitempty"`        // 创建时间
	MessageClientId  string                 `json:"message_client_id,omitempty"`  // 消息客户端ID
	SenderClientType int                    `json:"sender_client_type,omitempty"` // 发送者客户端类型
	Text             string                 `json:"text,omitempty"`               // 文本内容
	Attachment       map[string]interface{} `json:"attachment,omitempty"`         // 附件
	Extension        string                 `json:"extension,omitempty"`          // 扩展字段
	SubType          int                    `json:"sub_type,omitempty"`           // 子类型
}
