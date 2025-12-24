package conversation_unread

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
)

// ConversationUnreadV2Service Conversation Unread V2服务接口
type ConversationUnreadV2Service interface {
	// OverViewsConversation 获取会话概览
	OverViewsConversation(req *OverViewsConversationRequestV2) (*core.Result[*OverViewsConversationResponseV2], error)

	// ClearConversationUnread 清除会话未读数
	ClearConversationUnread(req *ClearConversationUnreadRequestV2) (*core.Result[*ClearConversationUnreadResponseV2], error)
}

// ConversationUnreadV2ServiceImpl Conversation Unread V2服务实现
type ConversationUnreadV2ServiceImpl struct {
	httpClient core.YunxinApiHttpClient
}

// NewConversationUnreadV2Service 创建Conversation Unread V2服务实例
func NewConversationUnreadV2Service(httpClient core.YunxinApiHttpClient) ConversationUnreadV2Service {
	return &ConversationUnreadV2ServiceImpl{
		httpClient: httpClient,
	}
}
