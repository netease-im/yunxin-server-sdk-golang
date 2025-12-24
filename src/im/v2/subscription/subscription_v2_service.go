package subscription

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
)

// SubscriptionV2Service Subscription V2服务接口
type SubscriptionV2Service interface {
	// SubscribeUserStatus 订阅用户状态
	SubscribeUserStatus(req *SubscribeUserStatusRequestV2) (*core.Result[*SubscribeUserStatusResponseV2], error)

	// UnsubscribeUserStatus 取消订阅用户状态
	UnsubscribeUserStatus(req *UnsubscribeUserStatusRequestV2) (*core.Result[*UnsubscribeUserStatusResponseV2], error)

	// QueryUserStatusSubscription 查询用户状态订阅
	QueryUserStatusSubscription(req *QueryUserStatusSubscriptionRequestV2) (*core.Result[*QueryUserStatusSubscriptionResponseV2], error)
}

// SubscriptionV2ServiceImpl Subscription V2服务实现
type SubscriptionV2ServiceImpl struct {
	httpClient core.YunxinApiHttpClient
}

// NewSubscriptionV2Service 创建Subscription V2服务实例
func NewSubscriptionV2Service(httpClient core.YunxinApiHttpClient) SubscriptionV2Service {
	return &SubscriptionV2ServiceImpl{
		httpClient: httpClient,
	}
}
