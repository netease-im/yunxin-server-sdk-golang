package chatroom_queue

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
)

// ChatroomQueueV2Service Chatroom Queue V2服务接口
type ChatroomQueueV2Service interface {
	// InitializeChatroomQueue 初始化聊天室队列
	InitializeChatroomQueue(req *InitializeChatroomQueueRequestV2) (*core.Result[*InitializeChatroomQueueResponseV2], error)

	// QueryChatroomQueueElements 查询聊天室队列元素
	QueryChatroomQueueElements(req *QueryChatroomQueueElementsRequestV2) (*core.Result[*QueryChatroomQueueElementsResponseV2], error)

	// UpdateChatroomQueue 更新聊天室队列
	UpdateChatroomQueue(req *UpdateChatroomQueueRequestV2) (*core.Result[*UpdateChatroomQueueResponseV2], error)

	// DeleteChatroomQueue 删除聊天室队列
	DeleteChatroomQueue(req *DeleteChatroomQueueRequestV2) (*core.Result[*DeleteChatroomQueueResponseV2], error)

	// PollChatroomQueueElement 轮询聊天室队列元素
	PollChatroomQueueElement(req *PollChatroomQueueElementRequestV2) (*core.Result[*PollChatroomQueueElementResponseV2], error)
}

// ChatroomQueueV2ServiceImpl Chatroom Queue V2服务实现
type ChatroomQueueV2ServiceImpl struct {
	httpClient core.YunxinApiHttpClient
}

// NewChatroomQueueV2Service 创建Chatroom Queue V2服务实例
func NewChatroomQueueV2Service(httpClient core.YunxinApiHttpClient) ChatroomQueueV2Service {
	return &ChatroomQueueV2ServiceImpl{
		httpClient: httpClient,
	}
}
