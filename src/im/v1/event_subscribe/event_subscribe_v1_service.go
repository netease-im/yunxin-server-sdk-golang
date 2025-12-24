package event_subscribe

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
)

// EventSubscribeV1Service 事件订阅服务接口
type EventSubscribeV1Service interface {
	// AddEventSubscribe 添加事件订阅
	AddEventSubscribe(req *AddEventSubscribeRequestV1) (*core.Result[*AddEventSubscribeResponseV1], error)

	// QueryEventSubscribe 查询事件订阅
	QueryEventSubscribe(req *QueryEventSubscribeRequestV1) (*core.Result[*QueryEventSubscribeResponseV1], error)

	// DeleteEventSubscribe 删除特定事件订阅
	DeleteEventSubscribe(req *DeleteEventSubscribeRequestV1) (*core.Result[*DeleteEventSubscribeResponseV1], error)

	// BatchDeleteEventSubscribe 批量删除所有事件订阅
	BatchDeleteEventSubscribe(req *BatchDeleteEventSubscribeRequestV1) (*core.Result[*BatchDeleteEventSubscribeResponseV1], error)
}

// EventSubscribeV1ServiceImpl 事件订阅服务实现
type EventSubscribeV1ServiceImpl struct {
	httpClient core.YunxinApiHttpClient
}

// NewEventSubscribeV1Service 创建事件订阅服务实例
func NewEventSubscribeV1Service(httpClient core.YunxinApiHttpClient) EventSubscribeV1Service {
	return &EventSubscribeV1ServiceImpl{
		httpClient: httpClient,
	}
}
