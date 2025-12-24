package chatroom

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/util"
)

// CreateChatroom 创建聊天室
func (c *ChatroomV2ServiceImpl) CreateChatroom(req *CreateChatroomRequestV2) (*core.Result[*CreateChatroomResponseV2], error) {
	// Validate required parameters
	if req.Creator == "" {
		return nil, fmt.Errorf("creator cannot be empty")
	}

	if req.RoomName == "" {
		return nil, fmt.Errorf("room name cannot be empty")
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.POST, CreateChatroom, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*CreateChatroomResponseV2](apiResponse)
}

// GetChatroomInfo 获取聊天室信息
func (c *ChatroomV2ServiceImpl) GetChatroomInfo(req *GetChatroomInfoRequestV2) (*core.Result[*GetChatroomInfoResponseV2], error) {
	// Validate required parameters
	if req.RoomId == 0 {
		return nil, fmt.Errorf("chatroom ID cannot be empty")
	}

	pathParams := map[string]string{
		"room_id": strconv.FormatInt(req.RoomId, 10),
	}

	queryParams := make(map[string]string)
	if req.NeedOnlineUserCount {
		queryParams["need_online_user_count"] = "true"
	}
	if req.OnlineUserCountByType {
		queryParams["online_user_count_by_type"] = "true"
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.GET, GetChatroomInfo, pathParams, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*GetChatroomInfoResponseV2](apiResponse)
}

// UpdateChatroomInfo 更新聊天室信息
func (c *ChatroomV2ServiceImpl) UpdateChatroomInfo(req *UpdateChatroomInfoRequestV2) (*core.Result[*UpdateChatroomInfoResponseV2], error) {
	// Validate required parameters
	if req.RoomId == 0 {
		return nil, fmt.Errorf("chatroom ID cannot be empty")
	}

	pathParams := map[string]string{
		"room_id": strconv.FormatInt(req.RoomId, 10),
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.PATCH, UpdateChatroomInfo, pathParams, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*UpdateChatroomInfoResponseV2](apiResponse)
}

// GetChatroomAddress 获取聊天室地址
func (c *ChatroomV2ServiceImpl) GetChatroomAddress(req *GetChatroomAddressRequestV2) (*core.Result[*GetChatroomAddressResponseV2], error) {
	// Validate required parameters
	if req.RoomId == 0 {
		return nil, fmt.Errorf("chatroom ID cannot be empty")
	}

	if req.AccountId == "" {
		return nil, fmt.Errorf("account ID cannot be empty")
	}

	if req.ClientIp == "" {
		return nil, fmt.Errorf("client IP cannot be empty")
	}

	pathParams := map[string]string{
		"room_id": strconv.FormatInt(req.RoomId, 10),
	}

	queryParams := map[string]string{
		"account_id": req.AccountId,
		"client_ip":  req.ClientIp,
	}

	if req.ClientType > 0 {
		queryParams["client_type"] = strconv.Itoa(req.ClientType)
	}

	if req.IpType > 0 {
		queryParams["ip_type"] = strconv.Itoa(req.IpType)
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.GET, GetChatroomAddress, pathParams, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*GetChatroomAddressResponseV2](apiResponse)
}

// UpdateChatroomStatus 更新聊天室状态
func (c *ChatroomV2ServiceImpl) UpdateChatroomStatus(req *UpdateChatroomStatusRequestV2) (*core.Result[*UpdateChatroomStatusResponseV2], error) {
	// Validate required parameters
	if req.OperatorAccountId == "" {
		return nil, fmt.Errorf("operator account ID cannot be empty")
	}

	if req.RoomId == 0 {
		return nil, fmt.Errorf("chatroom ID cannot be empty")
	}

	pathParams := map[string]string{
		"room_id": strconv.FormatInt(req.RoomId, 10),
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.PATCH, UpdateChatroomStatus, pathParams, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*UpdateChatroomStatusResponseV2](apiResponse)
}

// ToggleChatroomMute 切换聊天室禁言
func (c *ChatroomV2ServiceImpl) ToggleChatroomMute(req *ToggleChatroomMuteRequestV2) (*core.Result[*ToggleChatroomMuteResponseV2], error) {
	// Validate required parameters
	if req.OperatorAccountId == "" {
		return nil, fmt.Errorf("operator account ID cannot be empty")
	}

	if req.RoomId == 0 {
		return nil, fmt.Errorf("chatroom ID cannot be empty")
	}

	pathParams := map[string]string{
		"room_id": strconv.FormatInt(req.RoomId, 10),
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.PATCH, ToggleChatroomMute, pathParams, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*ToggleChatroomMuteResponseV2](apiResponse)
}

// ToggleInOutNotification 切换进出通知
func (c *ChatroomV2ServiceImpl) ToggleInOutNotification(req *ToggleInOutNotificationRequestV2) (*core.Result[*ToggleInOutNotificationResponseV2], error) {
	// Validate required parameters
	if req.RoomId == 0 {
		return nil, fmt.Errorf("chatroom ID cannot be empty")
	}

	pathParams := map[string]string{
		"room_id": strconv.FormatInt(req.RoomId, 10),
	}

	queryParams := map[string]string{
		"in_out_notification": fmt.Sprintf("%t", req.InOutNotification),
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.PATCH, ToggleInOutNotification, pathParams, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*ToggleInOutNotificationResponseV2](apiResponse)
}

// QueryOpenChatrooms 查询开放聊天室
func (c *ChatroomV2ServiceImpl) QueryOpenChatrooms(req *QueryOpenChatroomsRequestV2) (*core.Result[*QueryOpenChatroomsResponseV2], error) {
	// Validate required parameters
	if req.CreatorAccountId == "" {
		return nil, fmt.Errorf("creator account ID cannot be empty")
	}

	queryParams := map[string]string{
		"creator_account_id": req.CreatorAccountId,
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.GET, QueryOpenChatrooms, nil, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*QueryOpenChatroomsResponseV2](apiResponse)
}

// ListOnlineMembers 列出在线成员
func (c *ChatroomV2ServiceImpl) ListOnlineMembers(req *ListOnlineMembersRequestV2) (*core.Result[*ListOnlineMembersResponseV2], error) {
	// Validate required parameters
	if req.RoomId == 0 {
		return nil, fmt.Errorf("chatroom ID cannot be empty")
	}

	if req.Limit == 0 {
		return nil, fmt.Errorf("limit cannot be empty")
	}

	pathParams := map[string]string{
		"room_id": strconv.FormatInt(req.RoomId, 10),
	}

	queryParams := map[string]string{
		"offset": strconv.FormatInt(req.Offset, 10),
		"limit":  strconv.Itoa(req.Limit),
	}

	if req.MemberRoles != "" {
		queryParams["member_roles"] = req.MemberRoles
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.GET, ListOnlineMembers, pathParams, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*ListOnlineMembersResponseV2](apiResponse)
}

// ListFixedMembers 列出固定成员
func (c *ChatroomV2ServiceImpl) ListFixedMembers(req *ListFixedMembersRequestV2) (*core.Result[*ListFixedMembersResponseV2], error) {
	// Validate required parameters
	if req.RoomId == 0 {
		return nil, fmt.Errorf("chatroom ID cannot be empty")
	}

	pathParams := map[string]string{
		"room_id": strconv.FormatInt(req.RoomId, 10),
	}

	queryParams := make(map[string]string)
	if req.MemberRoles != "" {
		queryParams["member_roles"] = req.MemberRoles
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.GET, ListFixedMembers, pathParams, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*ListFixedMembersResponseV2](apiResponse)
}
