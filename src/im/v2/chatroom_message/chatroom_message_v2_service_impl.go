package chatroom_message

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/util"
)

// SendChatroomMessage 发送聊天室消息
func (c *ChatroomMessageV2ServiceImpl) SendChatroomMessage(req *SendChatroomMessageRequestV2) (*core.Result[*SendChatroomMessageResponseV2], error) {
	// Validate required parameters
	if req.RoomId == 0 {
		return nil, fmt.Errorf("room ID cannot be empty")
	}

	if req.SenderId == "" {
		return nil, fmt.Errorf("sender ID cannot be empty")
	}

	if req.Message == nil {
		return nil, fmt.Errorf("message cannot be nil")
	}

	pathParams := map[string]string{
		"room_id": strconv.FormatInt(req.RoomId, 10),
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.POST, SendChatroomMessage, pathParams, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*SendChatroomMessageResponseV2](apiResponse)
}

// BatchSendChatroomMessages 批量发送聊天室消息
func (c *ChatroomMessageV2ServiceImpl) BatchSendChatroomMessages(req *BatchSendChatroomMessagesRequestV2) (*core.Result[*BatchSendChatroomMessagesResponseV2], error) {
	// Validate required parameters
	if req.RoomId == 0 {
		return nil, fmt.Errorf("room ID cannot be empty")
	}

	if req.SenderId == "" {
		return nil, fmt.Errorf("sender ID cannot be empty")
	}

	if len(req.Messages) == 0 {
		return nil, fmt.Errorf("messages cannot be empty")
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.POST, BatchSendChatroomMessages, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*BatchSendChatroomMessagesResponseV2](apiResponse)
}

// RecallChatroomMessage 撤回聊天室消息
func (c *ChatroomMessageV2ServiceImpl) RecallChatroomMessage(req *RecallChatroomMessageRequestV2) (*core.Result[*RecallChatroomMessageResponseV2], error) {
	// Validate required parameters
	if req.RoomId == 0 {
		return nil, fmt.Errorf("room ID cannot be empty")
	}

	if req.MessageClientId == "" {
		return nil, fmt.Errorf("message client ID cannot be empty")
	}

	pathParams := map[string]string{
		"room_id":           strconv.FormatInt(req.RoomId, 10),
		"message_client_id": req.MessageClientId,
	}

	queryParams := make(map[string]string)
	if req.OperatorId != "" {
		queryParams["operator_id"] = req.OperatorId
	}
	if req.SenderId != "" {
		queryParams["sender_id"] = req.SenderId
	}
	if req.CreateTime > 0 {
		queryParams["create_time"] = strconv.FormatInt(req.CreateTime, 10)
	}
	if req.NotificationEnabled {
		queryParams["notification_enabled"] = "true"
		if req.NotificationExtension != "" {
			queryParams["notification_extension"] = req.NotificationExtension
		}
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.DELETE, RecallOrDeleteChatroomMessage, pathParams, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*RecallChatroomMessageResponseV2](apiResponse)
}

// QueryChatroomHistoryMessages 查询聊天室历史消息
func (c *ChatroomMessageV2ServiceImpl) QueryChatroomHistoryMessages(req *QueryChatroomHistoryMessagesRequestV2) (*core.Result[*QueryChatroomHistoryMessagesResponseV2], error) {
	// Validate required parameters
	if req.RoomId == 0 {
		return nil, fmt.Errorf("room ID cannot be empty")
	}

	if req.SenderId == "" {
		return nil, fmt.Errorf("sender ID cannot be empty")
	}

	if req.Limit == 0 {
		return nil, fmt.Errorf("limit cannot be empty")
	}

	pathParams := map[string]string{
		"room_id": strconv.FormatInt(req.RoomId, 10),
	}

	queryParams := map[string]string{
		"sender_id": req.SenderId,
		"limit":     strconv.Itoa(req.Limit),
	}

	if req.PageToken != "" {
		queryParams["page_token"] = req.PageToken
	}

	if req.Descending {
		queryParams["descending"] = "true"
	}

	if req.MessageTypes != "" {
		queryParams["message_types"] = req.MessageTypes
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.GET, QueryChatroomHistoryMessages, pathParams, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*QueryChatroomHistoryMessagesResponseV2](apiResponse)
}
