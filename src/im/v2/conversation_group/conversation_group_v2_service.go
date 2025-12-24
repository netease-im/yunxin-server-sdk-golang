package conversation_group

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
)

// ConversationGroupV2Service Conversation Group V2服务接口
type ConversationGroupV2Service interface {
	// CreateConversationGroup 创建会话分组
	CreateConversationGroup(req *CreateConversationGroupRequestV2) (*core.Result[*CreateConversationGroupResponseV2], error)

	// UpdateConversationGroup 更新会话分组
	UpdateConversationGroup(req *UpdateConversationGroupRequestV2) (*core.Result[*UpdateConversationGroupResponseV2], error)

	// DeleteConversationGroup 删除会话分组
	DeleteConversationGroup(req *DeleteConversationGroupRequestV2) (*core.Result[*DeleteConversationGroupResponseV2], error)

	// GetConversationGroup 获取会话分组
	GetConversationGroup(req *GetConversationGroupRequestV2) (*core.Result[*GetConversationGroupResponseV2], error)

	// BatchGetConversationGroups 批量获取会话分组
	BatchGetConversationGroups(req *BatchGetConversationGroupsRequestV2) (*core.Result[*BatchGetConversationGroupsResponseV2], error)

	// ListAllConversationGroups 列出所有会话分组
	ListAllConversationGroups(req *ListAllConversationGroupsRequestV2) (*core.Result[*ListAllConversationGroupsResponseV2], error)
}

// ConversationGroupV2ServiceImpl Conversation Group V2服务实现
type ConversationGroupV2ServiceImpl struct {
	httpClient core.YunxinApiHttpClient
}

// NewConversationGroupV2Service 创建Conversation Group V2服务实例
func NewConversationGroupV2Service(httpClient core.YunxinApiHttpClient) ConversationGroupV2Service {
	return &ConversationGroupV2ServiceImpl{
		httpClient: httpClient,
	}
}
