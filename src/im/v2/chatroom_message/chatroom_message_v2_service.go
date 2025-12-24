package chatroom_message

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
)

// ChatroomMessageV2Service Chatroom Message V2服务接口
type ChatroomMessageV2Service interface {
	// SendChatroomMessage 发送聊天室消息
	SendChatroomMessage(req *SendChatroomMessageRequestV2) (*core.Result[*SendChatroomMessageResponseV2], error)

	// BatchSendChatroomMessages 批量发送聊天室消息
	BatchSendChatroomMessages(req *BatchSendChatroomMessagesRequestV2) (*core.Result[*BatchSendChatroomMessagesResponseV2], error)

	// RecallChatroomMessage 撤回聊天室消息
	RecallChatroomMessage(req *RecallChatroomMessageRequestV2) (*core.Result[*RecallChatroomMessageResponseV2], error)

	// QueryChatroomHistoryMessages 查询聊天室历史消息
	QueryChatroomHistoryMessages(req *QueryChatroomHistoryMessagesRequestV2) (*core.Result[*QueryChatroomHistoryMessagesResponseV2], error)
}

// ChatroomMessageV2ServiceImpl Chatroom Message V2服务实现
type ChatroomMessageV2ServiceImpl struct {
	httpClient core.YunxinApiHttpClient
}

// NewChatroomMessageV2Service 创建Chatroom Message V2服务实例
func NewChatroomMessageV2Service(httpClient core.YunxinApiHttpClient) ChatroomMessageV2Service {
	return &ChatroomMessageV2ServiceImpl{
		httpClient: httpClient,
	}
}
