package conversation

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
)

// ConversationV2Service Conversation V2服务接口
type ConversationV2Service interface {
	// CreateConversation 创建会话
	CreateConversation(req *CreateConversationRequestV2) (*core.Result[*CreateConversationResponseV2], error)

	// UpdateConversation 更新会话
	UpdateConversation(req *UpdateConversationRequestV2) (*core.Result[*UpdateConversationResponseV2], error)

	// DeleteConversation 删除会话
	DeleteConversation(req *DeleteConversationRequestV2) (*core.Result[*DeleteConversationResponseV2], error)

	// BatchDeleteConversations 批量删除会话
	BatchDeleteConversations(req *BatchDeleteConversationsRequestV2) (*core.Result[*BatchDeleteConversationsResponseV2], error)

	// GetConversation 获取会话
	GetConversation(req *GetConversationRequestV2) (*core.Result[*GetConversationResponseV2], error)

	// BatchGetConversations 批量获取会话
	BatchGetConversations(req *BatchGetConversationsRequestV2) (*core.Result[*BatchGetConversationsResponseV2], error)

	// ListConversations 列出会话
	ListConversations(req *ListConversationsRequestV2) (*core.Result[*ListConversationsResponseV2], error)

	// StickTopConversation 置顶会话
	StickTopConversation(req *StickTopConversationRequestV2) (*core.Result[*StickTopConversationResponseV2], error)
}

// ConversationV2ServiceImpl Conversation V2服务实现
type ConversationV2ServiceImpl struct {
	httpClient core.YunxinApiHttpClient
}

// NewConversationV2Service 创建Conversation V2服务实例
func NewConversationV2Service(httpClient core.YunxinApiHttpClient) ConversationV2Service {
	return &ConversationV2ServiceImpl{
		httpClient: httpClient,
	}
}
