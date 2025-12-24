package conversation_unread

// OverViewsConversationResponseV2 获取会话概览响应
type OverViewsConversationResponseV2 struct {
	AccountId   string `json:"account_id,omitempty"`   // 账号ID
	UnreadCount int    `json:"unread_count,omitempty"` // 未读数
}

// ClearConversationUnreadResponseV2 清除会话未读数响应
type ClearConversationUnreadResponseV2 struct {
	// Success response will only contain status code 200
}
