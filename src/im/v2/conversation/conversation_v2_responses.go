package conversation

// CreateConversationResponseV2 创建会话响应
type CreateConversationResponseV2 struct {
	ConversationId  string  `json:"conversation_id,omitempty"`  // 会话ID
	SenderId        string  `json:"sender_id,omitempty"`        // 发送者ID
	ReceiverId      string  `json:"receiver_id,omitempty"`      // 接收者ID
	Type            int     `json:"type,omitempty"`             // 会话类型
	StickTop        bool    `json:"stick_top,omitempty"`        // 是否置顶
	GroupIds        []int64 `json:"group_ids,omitempty"`        // 分组ID列表
	ServerExtension string  `json:"server_extension,omitempty"` // 服务端扩展字段
	UpdateTime      int64   `json:"update_time,omitempty"`      // 更新时间
}

// UpdateConversationResponseV2 更新会话响应
type UpdateConversationResponseV2 struct {
	ConversationId  string `json:"conversation_id,omitempty"`  // 会话ID
	SenderId        string `json:"sender_id,omitempty"`        // 发送者ID
	ReceiverId      string `json:"receiver_id,omitempty"`      // 接收者ID
	Type            int    `json:"type,omitempty"`             // 会话类型
	ServerExtension string `json:"server_extension,omitempty"` // 服务端扩展字段
	UpdateTime      int64  `json:"update_time,omitempty"`      // 更新时间
}

// DeleteConversationResponseV2 删除会话响应
type DeleteConversationResponseV2 struct {
	// Success response will only contain status code 200
}

// BatchDeleteConversationsResponseV2 批量删除会话响应
type BatchDeleteConversationsResponseV2 struct {
	SuccessList []string     `json:"success_list,omitempty"` // 成功列表
	FailedList  []FailedItem `json:"failed_list,omitempty"`  // 失败列表
}

// GetConversationResponseV2 获取会话响应
type GetConversationResponseV2 struct {
	ConversationId  string  `json:"conversation_id,omitempty"`  // 会话ID
	SenderId        string  `json:"sender_id,omitempty"`        // 发送者ID
	ReceiverId      string  `json:"receiver_id,omitempty"`      // 接收者ID
	Type            int     `json:"type,omitempty"`             // 会话类型
	StickTop        bool    `json:"stick_top,omitempty"`        // 是否置顶
	GroupIds        []int64 `json:"group_ids,omitempty"`        // 分组ID列表
	ServerExtension string  `json:"server_extension,omitempty"` // 服务端扩展字段
	UpdateTime      int64   `json:"update_time,omitempty"`      // 更新时间
}

// BatchGetConversationsResponseV2 批量获取会话响应
type BatchGetConversationsResponseV2 struct {
	SuccessList []ConversationInfo `json:"success_list,omitempty"` // 成功列表
	FailedList  []FailedItem       `json:"failed_list,omitempty"`  // 失败列表
}

// ListConversationsResponseV2 列出会话响应
type ListConversationsResponseV2 struct {
	HasMore   bool               `json:"has_more,omitempty"`   // 是否有更多数据
	NextToken string             `json:"next_token,omitempty"` // 下一页token
	Items     []ConversationItem `json:"items,omitempty"`      // 会话列表
}

// StickTopConversationResponseV2 置顶会话响应
type StickTopConversationResponseV2 struct {
	Type           int    `json:"type,omitempty"`            // 会话类型
	ConversationId string `json:"conversation_id,omitempty"` // 会话ID
	SenderId       string `json:"sender_id,omitempty"`       // 发送者ID
	ReceiverId     string `json:"receiver_id,omitempty"`     // 接收者ID
	StickTop       bool   `json:"stick_top,omitempty"`       // 是否置顶
	SortOrder      int64  `json:"sort_order,omitempty"`      // 排序顺序
	UpdateTime     int64  `json:"update_time,omitempty"`     // 更新时间
}

// ConversationInfo 会话信息
type ConversationInfo struct {
	ConversationId  string  `json:"conversation_id,omitempty"`  // 会话ID
	SenderId        string  `json:"sender_id,omitempty"`        // 发送者ID
	ReceiverId      string  `json:"receiver_id,omitempty"`      // 接收者ID
	Type            int     `json:"type,omitempty"`             // 会话类型
	StickTop        bool    `json:"stick_top,omitempty"`        // 是否置顶
	GroupIds        []int64 `json:"group_ids,omitempty"`        // 分组ID列表
	ServerExtension string  `json:"server_extension,omitempty"` // 服务端扩展字段
	UpdateTime      int64   `json:"update_time,omitempty"`      // 更新时间
}

// ConversationItem 会话项
type ConversationItem struct {
	ConversationId  string  `json:"conversation_id,omitempty"`  // 会话ID
	SenderId        string  `json:"sender_id,omitempty"`        // 发送者ID
	ReceiverId      string  `json:"receiver_id,omitempty"`      // 接收者ID
	Type            int     `json:"type,omitempty"`             // 会话类型
	StickTop        bool    `json:"stick_top,omitempty"`        // 是否置顶
	GroupIds        []int64 `json:"group_ids,omitempty"`        // 分组ID列表
	ServerExtension string  `json:"server_extension,omitempty"` // 服务端扩展字段
	UpdateTime      int64   `json:"update_time,omitempty"`      // 更新时间
}

// FailedItem 失败项
type FailedItem struct {
	ConversationId string `json:"conversation_id,omitempty"` // 会话ID
	ErrorCode      int    `json:"error_code,omitempty"`      // 错误码
	ErrorMsg       string `json:"error_msg,omitempty"`       // 错误信息
}
