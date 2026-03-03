package conversation_group

// ConversationGroupDetail 会话分组详情
type ConversationGroupDetail struct {
	GroupId         int64  `json:"group_id,omitempty"`         // 分组ID
	Name            string `json:"name,omitempty"`             // 分组名称
	ServerExtension string `json:"server_extension,omitempty"` // 服务端扩展字段
	CreateTime      int64  `json:"create_time,omitempty"`      // 创建时间
	UpdateTime      int64  `json:"update_time,omitempty"`      // 更新时间
}

// FailedConversationItem 失败的会话项
type FailedConversationItem struct {
	ConversationId string `json:"conversation_id,omitempty"` // 会话ID
	ErrorCode      int    `json:"error_code,omitempty"`      // 错误码
	ErrorMsg       string `json:"error_msg,omitempty"`       // 错误信息
}

// CreateConversationGroupResponseV2 创建会话分组响应
type CreateConversationGroupResponseV2 struct {
	ConversationGroup *ConversationGroupDetail `json:"conversation_group,omitempty"` // 会话分组信息
	SuccessList       []string                 `json:"success_list,omitempty"`       // 成功添加的会话ID列表
	FailedList        []FailedConversationItem `json:"failed_list,omitempty"`        // 失败的会话列表
}

// UpdateConversationGroupResponseV2 更新会话分组响应
type UpdateConversationGroupResponseV2 struct {
	ConversationGroup *ConversationGroupDetail `json:"conversation_group,omitempty"` // 会话分组信息
	SuccessList       []string                 `json:"success_list,omitempty"`       // 成功操作的会话ID列表
	FailedList        []FailedConversationItem `json:"failed_list,omitempty"`        // 失败的会话列表
}

// DeleteConversationGroupResponseV2 删除会话分组响应
type DeleteConversationGroupResponseV2 struct {
	// Success response will only contain status code 200
}

// GetConversationGroupResponseV2 获取会话分组响应
type GetConversationGroupResponseV2 struct {
	GroupId         int64    `json:"group_id,omitempty"`         // 分组ID
	AccountId       string   `json:"account_id,omitempty"`       // 账号ID
	Name            string   `json:"name,omitempty"`             // 分组名称
	ServerExtension string   `json:"server_extension,omitempty"` // 服务端扩展字段
	ConversationIds []string `json:"conversation_ids,omitempty"` // 会话ID列表
	CreateTime      int64    `json:"create_time,omitempty"`      // 创建时间
	UpdateTime      int64    `json:"update_time,omitempty"`      // 更新时间
}

// BatchGetConversationGroupsResponseV2 批量获取会话分组响应
type BatchGetConversationGroupsResponseV2 struct {
	SuccessList []ConversationGroupInfo `json:"success_list,omitempty"` // 成功列表
	FailedList  []FailedGroupItem       `json:"failed_list,omitempty"`  // 失败列表
}

// ListAllConversationGroupsResponseV2 列出所有会话分组响应
type ListAllConversationGroupsResponseV2 struct {
	Items []ConversationGroupInfo `json:"conversation_groups,omitempty"` // 分组列表
}

// ConversationGroupInfo 会话分组信息
type ConversationGroupInfo struct {
	GroupId         int64    `json:"group_id,omitempty"`         // 分组ID
	AccountId       string   `json:"account_id,omitempty"`       // 账号ID
	Name            string   `json:"name,omitempty"`             // 分组名称
	ServerExtension string   `json:"server_extension,omitempty"` // 服务端扩展字段
	ConversationIds []string `json:"conversation_ids,omitempty"` // 会话ID列表
	CreateTime      int64    `json:"create_time,omitempty"`      // 创建时间
	UpdateTime      int64    `json:"update_time,omitempty"`      // 更新时间
}

// FailedGroupItem 失败分组项
type FailedGroupItem struct {
	GroupId   int64  `json:"group_id,omitempty"`   // 分组ID
	ErrorCode int    `json:"error_code,omitempty"` // 错误码
	ErrorMsg  string `json:"error_msg,omitempty"`  // 错误信息
}
