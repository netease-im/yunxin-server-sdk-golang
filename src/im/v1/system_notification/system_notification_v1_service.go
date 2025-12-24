package system_notification

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
)

// System NotificationV1Service system_notification服务接口
type SystemNotificationV1Service interface {
	// SendAttachMsg SendAttachMsg
	SendAttachMsg(req *SendAttachMsgRequestV1) (*core.Result[*SendAttachMsgResponseV1], error)

	// SendBatchAttachMsg SendBatchAttachMsg
	SendBatchAttachMsg(req *SendBatchAttachMsgRequestV1) (*core.Result[*SendBatchAttachMsgResponseV1], error)
}

// SystemNotificationV1ServiceImpl system_notification服务实现
type SystemNotificationV1ServiceImpl struct {
	httpClient core.YunxinApiHttpClient
}

// NewSystemNotificationV1Service 创建system_notification服务实例
func NewSystemNotificationV1Service(httpClient core.YunxinApiHttpClient) SystemNotificationV1Service {
	return &SystemNotificationV1ServiceImpl{
		httpClient: httpClient,
	}
}
