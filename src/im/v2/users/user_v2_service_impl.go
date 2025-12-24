package users

import (
	"encoding/json"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/util"
)

// UpdateUser 更新用户信息
func (s *UserV2ServiceImpl) UpdateUser(req *UpdateUserRequestV2) (*core.Result[*UpdateUserResponseV2], error) {
	pathParams := map[string]string{
		"account_id": req.AccountId,
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.PUT, UpdateUser, pathParams, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*UpdateUserResponseV2](apiResponse)
}

// GetUser 获取用户信息
func (s *UserV2ServiceImpl) GetUser(req *GetUserRequestV2) (*core.Result[*GetUserResponseV2], error) {
	pathParams := map[string]string{
		"account_id": req.AccountId,
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.GET, GetUser, pathParams, nil, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*GetUserResponseV2](apiResponse)
}

// BatchGetUsers 批量获取用户信息
func (s *UserV2ServiceImpl) BatchGetUsers(req *BatchGetUsersRequestV2) (*core.Result[*BatchGetUsersResponseV2], error) {
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.POST, BatchGetUsers, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*BatchGetUsersResponseV2](apiResponse)
}

// GetUsersOnlineStatus 获取用户在线状态
func (s *UserV2ServiceImpl) GetUsersOnlineStatus(req *GetUserOnlineStatusRequestV2) (*core.Result[*GetUserOnlineStatusResponseV2], error) {
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.POST, GetUsersOnlineStatus, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*GetUserOnlineStatusResponseV2](apiResponse)
}
