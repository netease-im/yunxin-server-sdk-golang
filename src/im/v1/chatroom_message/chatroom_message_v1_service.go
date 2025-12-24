package chatroom_message

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
)

// ChatroomMessageV1Service 聊天室消息服务接口
type ChatroomMessageV1Service interface {
	// SendMsg 发送聊天室消息
	SendMsg(req *SendChatroomMsgRequestV1) (*core.Result[*SendChatroomMsgResponseV1], error)

	// BatchSendMsg 批量发送聊天室消息
	BatchSendMsg(req *BatchSendChatroomMsgRequestV1) (*core.Result[*BatchSendChatroomMsgResponseV1], error)

	// Recall 撤回聊天室消息
	Recall(req *RecallChatroomMsgRequestV1) (*core.Result[*RecallChatroomMsgResponseV1], error)

	// SendMsgToSomeone 发送聊天室定向消息
	SendMsgToSomeone(req *ChatroomTargetMsgRequestV1) (*core.Result[*ChatroomTargetMsgResponseV1], error)

	// BatchSendMsgToSomeone 批量发送聊天室定向消息
	BatchSendMsgToSomeone(req *BatchChatroomTargetMsgRequestV1) (*core.Result[*BatchChatroomTargetMsgResponseV1], error)

	// Broadcast 发送聊天室全服广播消息
	Broadcast(req *ChatroomBroadcastRequestV1) (*core.Result[*SendChatroomMsgResponseV1], error)
}

// ChatroomMessageV1ServiceImpl 聊天室消息服务实现
type ChatroomMessageV1ServiceImpl struct {
	httpClient core.YunxinApiHttpClient
}

// NewChatroomMessageV1Service 创建聊天室消息服务实例
func NewChatroomMessageV1Service(httpClient core.YunxinApiHttpClient) ChatroomMessageV1Service {
	return &ChatroomMessageV1ServiceImpl{
		httpClient: httpClient,
	}
}
