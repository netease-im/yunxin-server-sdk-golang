package system_notification

import (
	"encoding/json"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/util"
)

// SendCustomNotification 发送自定义系统通知
func (s *CustomNotificationV2ServiceImpl) SendCustomNotification(req *SendCustomNotificationRequestV2) (*core.Result[*SendCustomNotificationResponseV2], error) {
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.POST, CustomNotification, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*SendCustomNotificationResponseV2](apiResponse)
}

// SendBatchCustomNotification 批量发送自定义系统通知
func (s *CustomNotificationV2ServiceImpl) SendBatchCustomNotification(req *SendBatchCustomNotificationRequestV2) (*core.Result[*SendBatchCustomNotificationResponseV2], error) {
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.POST, BatchCustomNotification, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*SendBatchCustomNotificationResponseV2](apiResponse)
}
