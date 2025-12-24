package conversation_group

// CreateConversationGroupRequestV2 创建会话分组请求
type CreateConversationGroupRequestV2 struct {
	AccountId       string   `json:"account_id"`                 // 账号ID
	Name            string   `json:"name"`                       // 分组名称
	ServerExtension string   `json:"server_extension,omitempty"` // 服务端扩展字段
	ConversationIds []string `json:"conversation_ids,omitempty"` // 会话ID列表
}

// UpdateConversationGroupRequestV2 更新会话分组请求
type UpdateConversationGroupRequestV2 struct {
	GroupId         int64          `json:"-"`                          // 分组ID，用于构造URL
	AccountId       string         `json:"account_id"`                 // 账号ID
	Name            string         `json:"name,omitempty"`             // 分组名称
	ServerExtension string         `json:"server_extension,omitempty"` // 服务端扩展字段
	Conversations   *Conversations `json:"conversations,omitempty"`    // 会话操作
}

// DeleteConversationGroupRequestV2 删除会话分组请求
type DeleteConversationGroupRequestV2 struct {
	GroupId   int64  `json:"-"`          // 分组ID，用于构造URL
	AccountId string `json:"account_id"` // 账号ID
}

// GetConversationGroupRequestV2 获取会话分组请求
type GetConversationGroupRequestV2 struct {
	GroupId   int64  `json:"-"`          // 分组ID，用于构造URL
	AccountId string `json:"account_id"` // 账号ID
}

// BatchGetConversationGroupsRequestV2 批量获取会话分组请求
type BatchGetConversationGroupsRequestV2 struct {
	AccountId string  `json:"account_id"` // 账号ID
	GroupIds  []int64 `json:"group_ids"`  // 分组ID列表
}

// ListAllConversationGroupsRequestV2 列出所有会话分组请求
type ListAllConversationGroupsRequestV2 struct {
	AccountId string `json:"account_id"` // 账号ID
}

// Conversations 会话操作
type Conversations struct {
	Type            int      `json:"type"`                       // 操作类型
	ConversationIds []string `json:"conversation_ids,omitempty"` // 会话ID列表
}
