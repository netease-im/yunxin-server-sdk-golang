package message

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
)

// MessageV1Service message服务接口
type MessageV1Service interface {
	// SendMessage SendMessage
	SendMessage(req *SendMessageRequestV1) (*core.Result[*SendMessageResponseV1], error)

	// SendBatchMessage SendBatchMessage
	SendBatchMessage(req *SendBatchMessageRequestV1) (*core.Result[*SendBatchMessageResponseV1], error)

	// UploadFile UploadFile
	UploadFile(req *UploadFileRequestV1) (*core.Result[*UploadFileResponseV1], error)

	// RecallMessage RecallMessage
	RecallMessage(req *RecallMessageRequestV1) (*core.Result[*RecallMessageResponseV1], error)

	// BroadcastMessage BroadcastMessage
	BroadcastMessage(req *BroadcastMessageRequestV1) (*core.Result[*BroadcastMessageResponseV1], error)

	// DeleteBroadcastMessageById DeleteBroadcastMessageById
	DeleteBroadcastMessageById(req *DeleteBroadcastMessageByIdRequestV1) (*core.Result[*DeleteBroadcastMessageByIdResponseV1], error)

	// DeleteMessage DeleteMessage
	DeleteMessage(req *DeleteMessageRequestV1) (*core.Result[*DeleteMessageResponseV1], error)

	// DeleteMessageOneWay DeleteMessageOneWay
	DeleteMessageOneWay(req *DeleteMessageOneWayRequestV1) (*core.Result[*DeleteMessageOneWayResponseV1], error)

	// DeleteFile DeleteFile
	DeleteFile(req *DeleteFileRequestV1) (*core.Result[*DeleteFileResponseV1], error)

	// DeleteRoamSession DeleteRoamSession
	DeleteRoamSession(req *DeleteRoamSessionRequestV1) (*core.Result[*DeleteRoamSessionResponseV1], error)

	// MarkReadMessage MarkReadMessage
	MarkReadMessage(req *MarkReadMessageRequestV1) (*core.Result[*MarkReadMessageResponseV1], error)

	// MarkReadTeamMessage MarkReadTeamMessage
	MarkReadTeamMessage(req *MarkReadTeamMessageRequestV1) (*core.Result[*MarkReadTeamMessageResponseV1], error)
}

// MessageV1ServiceImpl message服务实现
type MessageV1ServiceImpl struct {
	httpClient core.YunxinApiHttpClient
}

// NewMessageV1Service 创建message服务实例
func NewMessageV1Service(httpClient core.YunxinApiHttpClient) MessageV1Service {
	return &MessageV1ServiceImpl{
		httpClient: httpClient,
	}
}
