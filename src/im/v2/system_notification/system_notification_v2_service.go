package system_notification

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
)

// CustomNotificationV2Service 自定义系统通知服务接口
type CustomNotificationV2Service interface {
	// SendCustomNotification 发送自定义系统通知
	SendCustomNotification(req *SendCustomNotificationRequestV2) (*core.Result[*SendCustomNotificationResponseV2], error)

	// SendBatchCustomNotification 批量发送自定义系统通知
	SendBatchCustomNotification(req *SendBatchCustomNotificationRequestV2) (*core.Result[*SendBatchCustomNotificationResponseV2], error)
}

// CustomNotificationV2ServiceImpl 自定义系统通知服务实现
type CustomNotificationV2ServiceImpl struct {
	httpClient core.YunxinApiHttpClient
}

// NewCustomNotificationV2Service 创建自定义系统通知服务实例
func NewCustomNotificationV2Service(httpClient core.YunxinApiHttpClient) CustomNotificationV2Service {
	return &CustomNotificationV2ServiceImpl{
		httpClient: httpClient,
	}
}
