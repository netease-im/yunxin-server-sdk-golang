package subscription

import (
	"encoding/json"
	"fmt"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/util"
)

// SubscribeUserStatus 订阅用户状态
func (s *SubscriptionV2ServiceImpl) SubscribeUserStatus(req *SubscribeUserStatusRequestV2) (*core.Result[*SubscribeUserStatusResponseV2], error) {
	pathParams := map[string]string{
		"account_id": req.SubscriberAccountId,
	}

	queryParams := map[string]string{
		"duration": fmt.Sprintf("%d", req.Duration),
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.POST, SubscribeUserStatus, pathParams, queryParams, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*SubscribeUserStatusResponseV2](apiResponse)
}

// UnsubscribeUserStatus 取消订阅用户状态
func (s *SubscriptionV2ServiceImpl) UnsubscribeUserStatus(req *UnsubscribeUserStatusRequestV2) (*core.Result[*UnsubscribeUserStatusResponseV2], error) {
	pathParams := map[string]string{
		"account_id": req.SubscriberAccountId,
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.DELETE, UnsubscribeUserStatus, pathParams, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*UnsubscribeUserStatusResponseV2](apiResponse)
}

// QueryUserStatusSubscription 查询用户状态订阅
func (s *SubscriptionV2ServiceImpl) QueryUserStatusSubscription(req *QueryUserStatusSubscriptionRequestV2) (*core.Result[*QueryUserStatusSubscriptionResponseV2], error) {
	pathParams := map[string]string{
		"account_id": req.SubscriberAccountId,
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.PATCH, QueryUserStatusSubscription, pathParams, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*QueryUserStatusSubscriptionResponseV2](apiResponse)
}
