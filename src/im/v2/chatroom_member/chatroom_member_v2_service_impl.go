package chatroom_member

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/util"
)

// SetMemberRole 设置成员角色
func (c *ChatroomMemberV2ServiceImpl) SetMemberRole(req *SetMemberRoleRequestV2) (*core.Result[*SetMemberRoleResponseV2], error) {
	// Validate required parameters
	if req.RoomId == 0 {
		return nil, fmt.Errorf("chatroom ID cannot be empty")
	}

	if req.MemberRole == 0 {
		return nil, fmt.Errorf("member role cannot be empty")
	}

	pathParams := map[string]string{
		"account_id": req.AccountId,
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.POST, SetMemberRole, pathParams, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*SetMemberRoleResponseV2](apiResponse)
}

// UpdateOnlineMemberInfo 更新在线成员信息
func (c *ChatroomMemberV2ServiceImpl) UpdateOnlineMemberInfo(req *UpdateOnlineMemberInfoRequestV2) (*core.Result[*UpdateOnlineMemberInfoResponseV2], error) {
	// Validate required parameters
	if req.RoomId == 0 {
		return nil, fmt.Errorf("chatroom ID cannot be empty")
	}

	pathParams := map[string]string{
		"account_id": req.AccountId,
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.PATCH, UpdateOnlineMemberInfo, pathParams, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*UpdateOnlineMemberInfoResponseV2](apiResponse)
}

// ToggleChatBan 切换禁言
func (c *ChatroomMemberV2ServiceImpl) ToggleChatBan(req *ToggleChatBanRequestV2) (*core.Result[*ToggleChatBanResponseV2], error) {
	// Validate required parameters
	if req.RoomId == 0 {
		return nil, fmt.Errorf("chatroom ID cannot be empty")
	}

	if req.OperatorAccountId == "" {
		return nil, fmt.Errorf("operator account ID cannot be empty")
	}

	pathParams := map[string]string{
		"account_id": req.AccountId,
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.PATCH, ToggleChatBan, pathParams, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*ToggleChatBanResponseV2](apiResponse)
}

// ToggleTempChatBan 切换临时禁言
func (c *ChatroomMemberV2ServiceImpl) ToggleTempChatBan(req *ToggleTempChatBanRequestV2) (*core.Result[*ToggleTempChatBanResponseV2], error) {
	// Validate required parameters
	if req.RoomId == 0 {
		return nil, fmt.Errorf("chatroom ID cannot be empty")
	}

	if req.OperatorAccountId == "" {
		return nil, fmt.Errorf("operator account ID cannot be empty")
	}

	pathParams := map[string]string{
		"account_id": req.AccountId,
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.PATCH, ToggleTempChatBan, pathParams, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*ToggleTempChatBanResponseV2](apiResponse)
}

// ToggleBlocked 切换屏蔽
func (c *ChatroomMemberV2ServiceImpl) ToggleBlocked(req *ToggleBlockedRequestV2) (*core.Result[*ToggleBlockedResponseV2], error) {
	// Validate required parameters
	if req.RoomId == 0 {
		return nil, fmt.Errorf("chatroom ID cannot be empty")
	}

	if req.OperatorAccountId == "" {
		return nil, fmt.Errorf("operator account ID cannot be empty")
	}

	pathParams := map[string]string{
		"account_id": req.AccountId,
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.PATCH, ToggleBlocked, pathParams, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*ToggleBlockedResponseV2](apiResponse)
}

// ModifyMemberTags 修改成员标签
func (c *ChatroomMemberV2ServiceImpl) ModifyMemberTags(req *ModifyMemberTagsRequestV2) (*core.Result[*ModifyMemberTagsResponseV2], error) {
	// Validate required parameters
	if req.RoomId == 0 {
		return nil, fmt.Errorf("chatroom ID cannot be empty")
	}

	if len(req.Tags) == 0 {
		return nil, fmt.Errorf("tags cannot be empty")
	}

	pathParams := map[string]string{
		"account_id": req.AccountId,
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.PUT, ModifyMemberTags, pathParams, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*ModifyMemberTagsResponseV2](apiResponse)
}

// QueryTaggedMembersCount 查询标签成员数量
func (c *ChatroomMemberV2ServiceImpl) QueryTaggedMembersCount(req *QueryTaggedMembersCountRequestV2) (*core.Result[*QueryTaggedMembersCountResponseV2], error) {
	// Validate required parameters
	if req.RoomId == 0 {
		return nil, fmt.Errorf("chatroom ID cannot be empty")
	}

	if req.Tag == "" {
		return nil, fmt.Errorf("tag cannot be empty")
	}

	pathParams := map[string]string{
		"room_id": strconv.FormatInt(req.RoomId, 10),
	}

	queryParams := map[string]string{
		"tag": req.Tag,
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.GET, QueryTaggedMembersCount, pathParams, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*QueryTaggedMembersCountResponseV2](apiResponse)
}

// ListTaggedMembers 列出标签成员
func (c *ChatroomMemberV2ServiceImpl) ListTaggedMembers(req *ListTaggedMembersRequestV2) (*core.Result[*ListTaggedMembersResponseV2], error) {
	// Validate required parameters
	if req.RoomId == 0 {
		return nil, fmt.Errorf("chatroom ID cannot be empty")
	}

	if req.Tag == "" {
		return nil, fmt.Errorf("tag cannot be empty")
	}

	if req.Limit == 0 {
		return nil, fmt.Errorf("limit cannot be empty")
	}

	pathParams := map[string]string{
		"room_id": strconv.FormatInt(req.RoomId, 10),
	}

	queryParams := map[string]string{
		"tag":    req.Tag,
		"offset": strconv.FormatInt(req.Offset, 10),
		"limit":  strconv.Itoa(req.Limit),
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.GET, ListTagMembers, pathParams, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*ListTaggedMembersResponseV2](apiResponse)
}

// QueryChatroomBlacklist 查询聊天室黑名单
func (c *ChatroomMemberV2ServiceImpl) QueryChatroomBlacklist(req *QueryChatroomBlacklistRequestV2) (*core.Result[*QueryChatroomBlacklistResponseV2], error) {
	// Validate required parameters
	if req.RoomId == 0 {
		return nil, fmt.Errorf("chatroom ID cannot be empty")
	}

	pathParams := map[string]string{
		"room_id": strconv.FormatInt(req.RoomId, 10),
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.GET, QueryChatroomBlacklist, pathParams, nil, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*QueryChatroomBlacklistResponseV2](apiResponse)
}

// ToggleTaggedMembersChatBan 切换标签成员禁言
func (c *ChatroomMemberV2ServiceImpl) ToggleTaggedMembersChatBan(req *ToggleTaggedMembersChatBanRequestV2) (*core.Result[*ToggleTaggedMembersChatBanResponseV2], error) {
	// Validate required parameters
	if req.RoomId == 0 {
		return nil, fmt.Errorf("chatroom ID cannot be empty")
	}

	if req.OperatorAccountId == "" {
		return nil, fmt.Errorf("operator account ID cannot be empty")
	}

	if req.TargetTag == "" {
		return nil, fmt.Errorf("target tag cannot be empty")
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.PATCH, ToggleTaggedMembersChatBan, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*ToggleTaggedMembersChatBanResponseV2](apiResponse)
}

// BatchQueryChatroomMembers 批量查询聊天室成员
func (c *ChatroomMemberV2ServiceImpl) BatchQueryChatroomMembers(req *BatchQueryChatroomMembersRequestV2) (*core.Result[*BatchQueryChatroomMembersResponseV2], error) {
	// Validate required parameters
	if req.RoomId == 0 {
		return nil, fmt.Errorf("chatroom ID cannot be empty")
	}

	if len(req.AccountIds) == 0 {
		return nil, fmt.Errorf("account IDs cannot be empty")
	}

	pathParams := map[string]string{
		"room_id": strconv.FormatInt(req.RoomId, 10),
	}

	queryParams := map[string]string{
		"account_ids": strings.Join(req.AccountIds, ","),
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.GET, BatchQueryChatroomMembers, pathParams, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*BatchQueryChatroomMembersResponseV2](apiResponse)
}

// AddVirtualMembers 添加虚拟成员
func (c *ChatroomMemberV2ServiceImpl) AddVirtualMembers(req *AddVirtualMembersRequestV2) (*core.Result[*AddVirtualMembersResponseV2], error) {
	// Validate required parameters
	if req.RoomId == 0 {
		return nil, fmt.Errorf("chatroom ID cannot be empty")
	}

	if len(req.VirtualMembers) == 0 {
		return nil, fmt.Errorf("virtual members cannot be empty")
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.POST, AddVirtualMembers, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*AddVirtualMembersResponseV2](apiResponse)
}

// DeleteVirtualMembers 删除虚拟成员
func (c *ChatroomMemberV2ServiceImpl) DeleteVirtualMembers(req *DeleteVirtualMembersRequestV2) (*core.Result[*DeleteVirtualMembersResponseV2], error) {
	// Validate required parameters
	if req.RoomId == 0 {
		return nil, fmt.Errorf("chatroom ID cannot be empty")
	}

	if len(req.AccountIds) == 0 {
		return nil, fmt.Errorf("account IDs cannot be empty")
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.DELETE, DeleteVirtualMembers, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*DeleteVirtualMembersResponseV2](apiResponse)
}

// ClearVirtualMembers 清空虚拟成员
func (c *ChatroomMemberV2ServiceImpl) ClearVirtualMembers(req *ClearVirtualMembersRequestV2) (*core.Result[*ClearVirtualMembersResponseV2], error) {
	// Validate required parameters
	if req.RoomId == 0 {
		return nil, fmt.Errorf("chatroom ID cannot be empty")
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.POST, ClearVirtualMembers, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*ClearVirtualMembersResponseV2](apiResponse)
}

// QueryVirtualMembers 查询虚拟成员
func (c *ChatroomMemberV2ServiceImpl) QueryVirtualMembers(req *QueryVirtualMembersRequestV2) (*core.Result[*QueryVirtualMembersResponseV2], error) {
	// Validate required parameters
	if req.RoomId == 0 {
		return nil, fmt.Errorf("chatroom ID cannot be empty")
	}

	queryParams := map[string]string{
		"room_id": strconv.FormatInt(req.RoomId, 10),
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.GET, QueryVirtualMembers, nil, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*QueryVirtualMembersResponseV2](apiResponse)
}

// QueryChatBanned 查询禁言成员
func (c *ChatroomMemberV2ServiceImpl) QueryChatBanned(req *QueryChatBannedRequestV2) (*core.Result[*QueryChatBannedResponseV2], error) {
	// Validate required parameters
	if req.RoomId == 0 {
		return nil, fmt.Errorf("chatroom ID cannot be empty")
	}

	pathParams := map[string]string{
		"room_id": strconv.FormatInt(req.RoomId, 10),
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.GET, QueryChatBanned, pathParams, nil, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*QueryChatBannedResponseV2](apiResponse)
}
