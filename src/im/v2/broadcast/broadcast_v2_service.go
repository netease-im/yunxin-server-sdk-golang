package broadcast

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
)

// BroadcastV2Service Broadcast V2服务接口
type BroadcastV2Service interface {
	// SendBroadcastNotification 发送广播通知
	SendBroadcastNotification(req *SendBroadcastNotificationRequestV2) (*core.Result[*SendBroadcastNotificationResponseV2], error)

	// DeleteBroadcastNotification 删除广播通知
	DeleteBroadcastNotification(req *DeleteBroadcastNotificationRequestV2) (*core.Result[*DeleteBroadcastNotificationResponseV2], error)

	// QueryBroadcastNotification 查询广播通知
	QueryBroadcastNotification(req *QueryBroadcastNotificationRequestV2) (*core.Result[*QueryBroadcastNotificationResponseV2], error)

	// QueryBroadcastNotificationList 查询广播通知列表
	QueryBroadcastNotificationList(req *QueryBroadcastNotificationListRequestV2) (*core.Result[*QueryBroadcastNotificationListResponseV2], error)

	// SendChatroomBroadcastNotification 发送聊天室广播通知
	SendChatroomBroadcastNotification(req *SendChatroomBroadcastNotificationRequestV2) (*core.Result[*SendChatroomBroadcastNotificationResponseV2], error)
}

// BroadcastV2ServiceImpl Broadcast V2服务实现
type BroadcastV2ServiceImpl struct {
	httpClient core.YunxinApiHttpClient
}

// NewBroadcastV2Service 创建Broadcast V2服务实例
func NewBroadcastV2Service(httpClient core.YunxinApiHttpClient) BroadcastV2Service {
	return &BroadcastV2ServiceImpl{
		httpClient: httpClient,
	}
}
