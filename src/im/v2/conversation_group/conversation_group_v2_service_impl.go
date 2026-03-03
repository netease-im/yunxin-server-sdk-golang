package conversation_group

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/util"
)

// CreateConversationGroup 创建会话分组
func (c *ConversationGroupV2ServiceImpl) CreateConversationGroup(req *CreateConversationGroupRequestV2) (*core.Result[*CreateConversationGroupResponseV2], error) {
	// Validate required parameters
	if req.AccountId == "" {
		return nil, fmt.Errorf("account ID cannot be empty")
	}

	if req.Name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.POST, CreateConversationGroup, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*CreateConversationGroupResponseV2](apiResponse)
}

// UpdateConversationGroup 更新会话分组
func (c *ConversationGroupV2ServiceImpl) UpdateConversationGroup(req *UpdateConversationGroupRequestV2) (*core.Result[*UpdateConversationGroupResponseV2], error) {
	// Validate required parameters
	if req.GroupId == 0 {
		return nil, fmt.Errorf("group ID cannot be empty")
	}

	if req.AccountId == "" {
		return nil, fmt.Errorf("account ID cannot be empty")
	}

	pathParams := map[string]string{
		"group_id": strconv.FormatInt(req.GroupId, 10),
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.PATCH, UpdateConversationGroup, pathParams, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*UpdateConversationGroupResponseV2](apiResponse)
}

// DeleteConversationGroup 删除会话分组
func (c *ConversationGroupV2ServiceImpl) DeleteConversationGroup(req *DeleteConversationGroupRequestV2) (*core.Result[*DeleteConversationGroupResponseV2], error) {
	// Validate required parameters
	if req.GroupId == 0 {
		return nil, fmt.Errorf("group ID cannot be empty")
	}

	if req.AccountId == "" {
		return nil, fmt.Errorf("account ID cannot be empty")
	}

	pathParams := map[string]string{
		"group_id": strconv.FormatInt(req.GroupId, 10),
	}

	queryParams := map[string]string{
		"account_id": req.AccountId,
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.DELETE, DeleteConversationGroup, pathParams, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*DeleteConversationGroupResponseV2](apiResponse)
}

// GetConversationGroup 获取会话分组
func (c *ConversationGroupV2ServiceImpl) GetConversationGroup(req *GetConversationGroupRequestV2) (*core.Result[*GetConversationGroupResponseV2], error) {
	// Validate required parameters
	if req.GroupId == 0 {
		return nil, fmt.Errorf("group ID cannot be empty")
	}

	if req.AccountId == "" {
		return nil, fmt.Errorf("account ID cannot be empty")
	}

	pathParams := map[string]string{
		"group_id": strconv.FormatInt(req.GroupId, 10),
	}

	queryParams := map[string]string{
		"account_id": req.AccountId,
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.GET, GetConversationGroup, pathParams, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*GetConversationGroupResponseV2](apiResponse)
}

// BatchGetConversationGroups 批量获取会话分组
func (c *ConversationGroupV2ServiceImpl) BatchGetConversationGroups(req *BatchGetConversationGroupsRequestV2) (*core.Result[*BatchGetConversationGroupsResponseV2], error) {
	// Validate required parameters
	if req.AccountId == "" {
		return nil, fmt.Errorf("account ID cannot be empty")
	}

	if len(req.GroupIds) == 0 {
		return nil, fmt.Errorf("group IDs cannot be empty")
	}

	queryParams := map[string]string{
		"account_id": req.AccountId,
	}

	// Convert group IDs to comma-separated string
	groupIdsStr := ""
	for i, id := range req.GroupIds {
		if i > 0 {
			groupIdsStr += ","
		}
		groupIdsStr += strconv.FormatInt(id, 10)
	}
	queryParams["group_ids"] = groupIdsStr

	apiResponse, err := c.httpClient.ExecuteV2Api(http.GET, BatchGetConversationGroups, nil, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*BatchGetConversationGroupsResponseV2](apiResponse)
}

// ListAllConversationGroups 列出所有会话分组
func (c *ConversationGroupV2ServiceImpl) ListAllConversationGroups(req *ListAllConversationGroupsRequestV2) (*core.Result[*ListAllConversationGroupsResponseV2], error) {
	// Validate required parameters
	if req.AccountId == "" {
		return nil, fmt.Errorf("account ID cannot be empty")
	}

	queryParams := map[string]string{
		"account_id": req.AccountId,
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.GET, ListAllConversationGroups, nil, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*ListAllConversationGroupsResponseV2](apiResponse)
}
