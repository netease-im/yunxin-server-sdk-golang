package chatroom_queue

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/util"
)

// InitializeChatroomQueue 初始化聊天室队列
func (c *ChatroomQueueV2ServiceImpl) InitializeChatroomQueue(req *InitializeChatroomQueueRequestV2) (*core.Result[*InitializeChatroomQueueResponseV2], error) {
	// Validate required parameters
	if req.RoomId == 0 {
		return nil, fmt.Errorf("room ID cannot be empty")
	}

	pathParams := map[string]string{
		"room_id": strconv.FormatInt(req.RoomId, 10),
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.POST, InitializeChatroomQueue, pathParams, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*InitializeChatroomQueueResponseV2](apiResponse)
}

// QueryChatroomQueueElements 查询聊天室队列元素
func (c *ChatroomQueueV2ServiceImpl) QueryChatroomQueueElements(req *QueryChatroomQueueElementsRequestV2) (*core.Result[*QueryChatroomQueueElementsResponseV2], error) {
	// Validate required parameters
	if req.RoomId == 0 {
		return nil, fmt.Errorf("room ID cannot be empty")
	}

	pathParams := map[string]string{
		"room_id": strconv.FormatInt(req.RoomId, 10),
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.POST, QueryChatroomQueueElements, pathParams, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*QueryChatroomQueueElementsResponseV2](apiResponse)
}

// UpdateChatroomQueue 更新聊天室队列
func (c *ChatroomQueueV2ServiceImpl) UpdateChatroomQueue(req *UpdateChatroomQueueRequestV2) (*core.Result[*UpdateChatroomQueueResponseV2], error) {
	// Validate required parameters
	if req.RoomId == 0 {
		return nil, fmt.Errorf("room ID cannot be empty")
	}

	if req.OperatorAccountId == "" {
		return nil, fmt.Errorf("operator account ID cannot be empty")
	}

	pathParams := map[string]string{
		"room_id": strconv.FormatInt(req.RoomId, 10),
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.PUT, UpdateChatroomQueue, pathParams, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*UpdateChatroomQueueResponseV2](apiResponse)
}

// DeleteChatroomQueue 删除聊天室队列
func (c *ChatroomQueueV2ServiceImpl) DeleteChatroomQueue(req *DeleteChatroomQueueRequestV2) (*core.Result[*DeleteChatroomQueueResponseV2], error) {
	// Validate required parameters
	if req.RoomId == 0 {
		return nil, fmt.Errorf("room ID cannot be empty")
	}

	pathParams := map[string]string{
		"room_id": strconv.FormatInt(req.RoomId, 10),
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.DELETE, DeleteChatroomQueue, pathParams, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*DeleteChatroomQueueResponseV2](apiResponse)
}

// PollChatroomQueueElement 轮询聊天室队列元素
func (c *ChatroomQueueV2ServiceImpl) PollChatroomQueueElement(req *PollChatroomQueueElementRequestV2) (*core.Result[*PollChatroomQueueElementResponseV2], error) {
	// Validate required parameters
	if req.RoomId == 0 {
		return nil, fmt.Errorf("room ID cannot be empty")
	}

	if req.ElementKey == "" {
		return nil, fmt.Errorf("element key cannot be empty")
	}

	pathParams := map[string]string{
		"room_id": strconv.FormatInt(req.RoomId, 10),
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.POST, PollChatroomQueueElement, pathParams, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*PollChatroomQueueElementResponseV2](apiResponse)
}
