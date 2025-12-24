package friend

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/util"
)

// AddFriend 添加好友
func (f *FriendV2ServiceImpl) AddFriend(req *AddFriendRequestV2) (*core.Result[*AddFriendResponseV2], error) {
	// Validate required parameters
	if req.AccountId == "" {
		return nil, fmt.Errorf("account ID cannot be empty")
	}

	if req.FriendAccountId == "" {
		return nil, fmt.Errorf("friend account ID cannot be empty")
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := f.httpClient.ExecuteV2Api(http.POST, Friends, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*AddFriendResponseV2](apiResponse)
}

// DeleteFriend 删除好友
func (f *FriendV2ServiceImpl) DeleteFriend(req *DeleteFriendRequestV2) (*core.Result[*DeleteFriendResponseV2], error) {
	// Validate required parameters
	if req.AccountId == "" {
		return nil, fmt.Errorf("account ID cannot be empty")
	}

	if req.FriendAccountId == "" {
		return nil, fmt.Errorf("friend account ID cannot be empty")
	}

	pathParams := map[string]string{
		"account_id": req.AccountId,
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := f.httpClient.ExecuteV2Api(http.DELETE, DeleteFriend, pathParams, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*DeleteFriendResponseV2](apiResponse)
}

// UpdateFriend 更新好友
func (f *FriendV2ServiceImpl) UpdateFriend(req *UpdateFriendRequestV2) (*core.Result[*UpdateFriendResponseV2], error) {
	// Validate required parameters
	if req.AccountId == "" {
		return nil, fmt.Errorf("account ID cannot be empty")
	}

	if req.FriendAccountId == "" {
		return nil, fmt.Errorf("friend account ID cannot be empty")
	}

	pathParams := map[string]string{
		"account_id": req.AccountId,
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := f.httpClient.ExecuteV2Api(http.PATCH, UpdateFriend, pathParams, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*UpdateFriendResponseV2](apiResponse)
}

// GetFriend 获取好友
func (f *FriendV2ServiceImpl) GetFriend(req *GetFriendRequestV2) (*core.Result[*GetFriendResponseV2], error) {
	// Validate required parameters
	if req.AccountId == "" {
		return nil, fmt.Errorf("account ID cannot be empty")
	}

	if req.FriendAccountId == "" {
		return nil, fmt.Errorf("friend account ID cannot be empty")
	}

	pathParams := map[string]string{
		"account_id": req.AccountId,
	}

	queryParams := map[string]string{
		"friend_account_id": req.FriendAccountId,
	}

	apiResponse, err := f.httpClient.ExecuteV2Api(http.GET, GetFriend, pathParams, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*GetFriendResponseV2](apiResponse)
}

// ListFriends 列出好友
func (f *FriendV2ServiceImpl) ListFriends(req *ListFriendsRequestV2) (*core.Result[*ListFriendsResponseV2], error) {
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

	apiResponse, err := f.httpClient.ExecuteV2Api(http.GET, ListFriends, nil, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*ListFriendsResponseV2](apiResponse)
}

// HandleFriendAddition 处理好友添加
func (f *FriendV2ServiceImpl) HandleFriendAddition(req *HandleFriendAdditionRequestV2) (*core.Result[*HandleFriendAdditionResponseV2], error) {
	// Validate required parameters
	if req.AccountId == "" {
		return nil, fmt.Errorf("account ID cannot be empty")
	}

	if req.FriendAccountId == "" {
		return nil, fmt.Errorf("friend account ID cannot be empty")
	}

	if req.HandleType == 0 {
		return nil, fmt.Errorf("handle type cannot be empty")
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := f.httpClient.ExecuteV2Api(http.POST, HandleFriendAddition, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*HandleFriendAdditionResponseV2](apiResponse)
}
