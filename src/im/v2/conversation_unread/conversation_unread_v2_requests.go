package conversation_unread

// OverViewsConversationRequestV2 获取会话概览请求
type OverViewsConversationRequestV2 struct {
	AccountId string `json:"-"` // 账号ID，用于构造URL
}

// ClearConversationUnreadRequestV2 清除会话未读数请求
type ClearConversationUnreadRequestV2 struct {
	ConversationId string `json:"-"` // 会话ID，用于构造URL
}
