package conversation

import "fmt"

// CreateConversationRequestV2 创建会话请求
type CreateConversationRequestV2 struct {
	ConversationId string `json:"conversation_id"` // 会话ID
}

// CreateP2PConversationId 创建点对点会话ID
func CreateP2PConversationId(senderId, receiverId string) string {
	return fmt.Sprintf("%s|1|%s", senderId, receiverId)
}

// CreateAdvancedTeamConversationId 创建高级群会话ID
func CreateAdvancedTeamConversationId(senderId, teamId string) string {
	return fmt.Sprintf("%s|2|%s", senderId, teamId)
}

// CreateSuperTeamConversationId 创建超大群会话ID
func CreateSuperTeamConversationId(senderId, teamId string) string {
	return fmt.Sprintf("%s|3|%s", senderId, teamId)
}

// UpdateConversationRequestV2 更新会话请求
type UpdateConversationRequestV2 struct {
	ConversationId  string `json:"-"`                // 会话ID，用于构造URL
	ServerExtension string `json:"server_extension"` // 服务端扩展字段
}

// DeleteConversationRequestV2 删除会话请求
type DeleteConversationRequestV2 struct {
	ConversationId string `json:"-"` // 会话ID，用于构造URL
	ClearMessage   bool   `json:"-"` // 是否清除消息，用于构造查询参数
}

// BatchDeleteConversationsRequestV2 批量删除会话请求
type BatchDeleteConversationsRequestV2 struct {
	ConversationIds []string `json:"-"` // 会话ID列表，用于构造查询参数
	ClearMessage    bool     `json:"-"` // 是否清除消息，用于构造查询参数
}

// GetConversationRequestV2 获取会话请求
type GetConversationRequestV2 struct {
	ConversationId string `json:"-"` // 会话ID，用于构造URL
}

// BatchGetConversationsRequestV2 批量获取会话请求
type BatchGetConversationsRequestV2 struct {
	ConversationIds []string `json:"conversation_ids"` // 会话ID列表
}

// ListConversationsRequestV2 列出会话请求
type ListConversationsRequestV2 struct {
	AccountId string `json:"-"` // 账号ID，用于构造查询参数
	PageToken string `json:"-"` // 分页token，用于构造查询参数
	Limit     int    `json:"-"` // 限制数量，用于构造查询参数
}

// StickTopConversationRequestV2 置顶会话请求
type StickTopConversationRequestV2 struct {
	ConversationId string `json:"-"`        // 会话ID，用于构造URL
	TopType        int    `json:"top_type"` // 置顶类型
}
