package broadcast

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/util"
)

// SendBroadcastNotification 发送广播通知
func (b *BroadcastV2ServiceImpl) SendBroadcastNotification(req *SendBroadcastNotificationRequestV2) (*core.Result[*SendBroadcastNotificationResponseV2], error) {
	// Validate required parameters
	if req.Content == "" {
		return nil, fmt.Errorf("content cannot be empty")
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := b.httpClient.ExecuteV2Api(http.POST, BroadcastNotification, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*SendBroadcastNotificationResponseV2](apiResponse)
}

// DeleteBroadcastNotification 删除广播通知
func (b *BroadcastV2ServiceImpl) DeleteBroadcastNotification(req *DeleteBroadcastNotificationRequestV2) (*core.Result[*DeleteBroadcastNotificationResponseV2], error) {
	// Validate required parameters
	if req.BroadcastId == "" {
		return nil, fmt.Errorf("broadcast ID cannot be empty")
	}

	pathParams := map[string]string{
		"broadcast_id": req.BroadcastId,
	}

	apiResponse, err := b.httpClient.ExecuteV2Api(http.DELETE, DeleteBroadcastNotification, pathParams, nil, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*DeleteBroadcastNotificationResponseV2](apiResponse)
}

// QueryBroadcastNotification 查询广播通知
func (b *BroadcastV2ServiceImpl) QueryBroadcastNotification(req *QueryBroadcastNotificationRequestV2) (*core.Result[*QueryBroadcastNotificationResponseV2], error) {
	// Validate required parameters
	if req.BroadcastId == "" {
		return nil, fmt.Errorf("broadcast ID cannot be empty")
	}

	pathParams := map[string]string{
		"broadcast_id": req.BroadcastId,
	}

	apiResponse, err := b.httpClient.ExecuteV2Api(http.GET, QueryBroadcastNotification, pathParams, nil, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*QueryBroadcastNotificationResponseV2](apiResponse)
}

// QueryBroadcastNotificationList 查询广播通知列表
func (b *BroadcastV2ServiceImpl) QueryBroadcastNotificationList(req *QueryBroadcastNotificationListRequestV2) (*core.Result[*QueryBroadcastNotificationListResponseV2], error) {
	// Build query parameters
	queryParams := make(map[string]string)

	if req.PageToken != "" {
		queryParams["page_token"] = req.PageToken
	}

	if req.Limit > 0 {
		queryParams["limit"] = strconv.Itoa(req.Limit)
	}

	if req.Type > 0 {
		queryParams["type"] = strconv.Itoa(req.Type)
	}

	apiResponse, err := b.httpClient.ExecuteV2Api(http.GET, BroadcastNotification, nil, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*QueryBroadcastNotificationListResponseV2](apiResponse)
}

// SendChatroomBroadcastNotification 发送聊天室广播通知
func (b *BroadcastV2ServiceImpl) SendChatroomBroadcastNotification(req *SendChatroomBroadcastNotificationRequestV2) (*core.Result[*SendChatroomBroadcastNotificationResponseV2], error) {
	// Validate required parameters
	if req.ClientBroadcastId == "" {
		return nil, fmt.Errorf("client broadcast ID cannot be empty")
	}

	if req.SenderId == "" {
		return nil, fmt.Errorf("sender ID cannot be empty")
	}

	if req.Message == nil {
		return nil, fmt.Errorf("message cannot be nil")
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := b.httpClient.ExecuteV2Api(http.POST, ChatroomBroadcastNotification, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*SendChatroomBroadcastNotificationResponseV2](apiResponse)
}
