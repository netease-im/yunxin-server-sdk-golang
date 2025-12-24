package conversation

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/util"
)

// CreateConversation 创建会话
func (c *ConversationV2ServiceImpl) CreateConversation(req *CreateConversationRequestV2) (*core.Result[*CreateConversationResponseV2], error) {
	// Validate required parameters
	if req.ConversationId == "" {
		return nil, fmt.Errorf("conversation ID cannot be empty")
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.POST, CreateConversation, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*CreateConversationResponseV2](apiResponse)
}

// UpdateConversation 更新会话
func (c *ConversationV2ServiceImpl) UpdateConversation(req *UpdateConversationRequestV2) (*core.Result[*UpdateConversationResponseV2], error) {
	// Validate required parameters
	if req.ConversationId == "" {
		return nil, fmt.Errorf("conversation ID cannot be empty")
	}

	pathParams := map[string]string{
		"conversation_id": req.ConversationId,
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.PATCH, UpdateConversation, pathParams, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*UpdateConversationResponseV2](apiResponse)
}

// DeleteConversation 删除会话
func (c *ConversationV2ServiceImpl) DeleteConversation(req *DeleteConversationRequestV2) (*core.Result[*DeleteConversationResponseV2], error) {
	// Validate required parameters
	if req.ConversationId == "" {
		return nil, fmt.Errorf("conversation ID cannot be empty")
	}

	pathParams := map[string]string{
		"conversation_id": req.ConversationId,
	}

	queryParams := make(map[string]string)
	if req.ClearMessage {
		queryParams["clear_message"] = "true"
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.DELETE, DeleteConversation, pathParams, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*DeleteConversationResponseV2](apiResponse)
}

// BatchDeleteConversations 批量删除会话
func (c *ConversationV2ServiceImpl) BatchDeleteConversations(req *BatchDeleteConversationsRequestV2) (*core.Result[*BatchDeleteConversationsResponseV2], error) {
	// Validate required parameters
	if len(req.ConversationIds) == 0 {
		return nil, fmt.Errorf("conversation IDs cannot be empty")
	}

	queryParams := map[string]string{
		"conversation_ids": strings.Join(req.ConversationIds, ","),
	}
	if req.ClearMessage {
		queryParams["clear_message"] = "true"
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.DELETE, BatchDeleteConversations, nil, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*BatchDeleteConversationsResponseV2](apiResponse)
}

// GetConversation 获取会话
func (c *ConversationV2ServiceImpl) GetConversation(req *GetConversationRequestV2) (*core.Result[*GetConversationResponseV2], error) {
	// Validate required parameters
	if req.ConversationId == "" {
		return nil, fmt.Errorf("conversation ID cannot be empty")
	}

	pathParams := map[string]string{
		"conversation_id": req.ConversationId,
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.GET, GetConversation, pathParams, nil, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*GetConversationResponseV2](apiResponse)
}

// BatchGetConversations 批量获取会话
func (c *ConversationV2ServiceImpl) BatchGetConversations(req *BatchGetConversationsRequestV2) (*core.Result[*BatchGetConversationsResponseV2], error) {
	// Validate required parameters
	if len(req.ConversationIds) == 0 {
		return nil, fmt.Errorf("conversation IDs cannot be empty")
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.POST, BatchGetConversations, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*BatchGetConversationsResponseV2](apiResponse)
}

// ListConversations 列出会话
func (c *ConversationV2ServiceImpl) ListConversations(req *ListConversationsRequestV2) (*core.Result[*ListConversationsResponseV2], error) {
	// Validate required parameters
	if req.AccountId == "" {
		return nil, fmt.Errorf("account ID cannot be empty")
	}

	queryParams := map[string]string{
		"account_id": req.AccountId,
	}

	if req.PageToken != "" {
		queryParams["page_token"] = req.PageToken
	}

	if req.Limit > 0 {
		queryParams["limit"] = strconv.Itoa(req.Limit)
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.GET, ListConversations, nil, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*ListConversationsResponseV2](apiResponse)
}

// StickTopConversation 置顶会话
func (c *ConversationV2ServiceImpl) StickTopConversation(req *StickTopConversationRequestV2) (*core.Result[*StickTopConversationResponseV2], error) {
	// Validate required parameters
	if req.ConversationId == "" {
		return nil, fmt.Errorf("conversation ID cannot be empty")
	}

	pathParams := map[string]string{
		"conversation_id": req.ConversationId,
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.POST, StickTopConversation, pathParams, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*StickTopConversationResponseV2](apiResponse)
}
