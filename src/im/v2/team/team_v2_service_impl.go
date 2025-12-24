package team

import (
	"encoding/json"
	"strconv"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/util"
)

// CreateTeam 创建群组
func (s *TeamV2ServiceImpl) CreateTeam(req *CreateTeamRequestV2) (*core.Result[*CreateTeamResponseV2], error) {
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.POST, CreateTeam, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*CreateTeamResponseV2](apiResponse)
}

// UpdateTeam 更新群组信息
func (s *TeamV2ServiceImpl) UpdateTeam(req *UpdateTeamRequestV2) (*core.Result[*UpdateTeamResponseV2], error) {
	pathParams := map[string]string{
		"team_id": req.TeamId,
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.PUT, UpdateTeam, pathParams, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*UpdateTeamResponseV2](apiResponse)
}

// DisbandTeam 解散群组
func (s *TeamV2ServiceImpl) DisbandTeam(req *DisbandTeamRequestV2) (*core.Result[*DisbandTeamResponseV2], error) {
	pathParams := map[string]string{
		"team_id": req.TeamId,
	}

	queryParams := map[string]string{
		"team_type":   strconv.Itoa(req.TeamType),
		"operator_id": req.OperatorId,
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.DELETE, DisbandTeam, pathParams, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*DisbandTeamResponseV2](apiResponse)
}

// BatchQueryTeamInfo 批量查询群组信息
func (s *TeamV2ServiceImpl) BatchQueryTeamInfo(req *BatchQueryTeamInfoRequestV2) (*core.Result[*BatchQueryTeamInfoResponseV2], error) {
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.POST, BatchQueryTeams, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*BatchQueryTeamInfoResponseV2](apiResponse)
}

// TransferOwner 转让群主
func (s *TeamV2ServiceImpl) TransferOwner(req *TransferTeamOwnerRequestV2) (*core.Result[*TransferTeamOwnerResponseV2], error) {
	pathParams := map[string]string{
		"team_id": req.TeamId,
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.POST, TransferOwner, pathParams, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*TransferTeamOwnerResponseV2](apiResponse)
}

// AddManagers 添加群管理员
func (s *TeamV2ServiceImpl) AddManagers(req *AddTeamManagersRequestV2) (*core.Result[*AddTeamManagersResponseV2], error) {
	pathParams := map[string]string{
		"team_id": req.TeamId,
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.POST, AddManager, pathParams, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*AddTeamManagersResponseV2](apiResponse)
}

// RemoveManagers 移除群管理员
func (s *TeamV2ServiceImpl) RemoveManagers(req *RemoveTeamManagersRequestV2) (*core.Result[*RemoveTeamManagersResponseV2], error) {
	pathParams := map[string]string{
		"team_id": req.TeamId,
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.POST, RemoveManager, pathParams, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*RemoveTeamManagersResponseV2](apiResponse)
}

// GetTeamInfo 获取群组信息
func (s *TeamV2ServiceImpl) GetTeamInfo(req *GetTeamInfoRequestV2) (*core.Result[*GetTeamInfoResponseV2], error) {
	pathParams := map[string]string{
		"team_id": req.TeamId,
	}

	queryParams := map[string]string{
		"team_type": strconv.Itoa(req.TeamType),
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.GET, GetTeamInfo, pathParams, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*GetTeamInfoResponseV2](apiResponse)
}

// ListTeamMembers 查询群成员列表
func (s *TeamV2ServiceImpl) ListTeamMembers(req *ListTeamMembersRequestV2) (*core.Result[*ListTeamMembersResponseV2], error) {
	pathParams := map[string]string{
		"team_id": req.TeamId,
	}

	queryParams := map[string]string{
		"team_type": strconv.Itoa(req.TeamType),
	}

	if req.Offset > 0 {
		queryParams["offset"] = strconv.Itoa(req.Offset)
	}
	if req.Limit > 0 {
		queryParams["limit"] = strconv.Itoa(req.Limit)
	}
	if req.NextToken != "" {
		queryParams["next_token"] = req.NextToken
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.GET, ListTeamMembers, pathParams, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*ListTeamMembersResponseV2](apiResponse)
}

// ListOnlineTeamMembers 查询群在线成员列表
func (s *TeamV2ServiceImpl) ListOnlineTeamMembers(req *ListOnlineTeamMembersRequestV2) (*core.Result[*ListOnlineTeamMembersResponseV2], error) {
	pathParams := map[string]string{
		"team_id": req.TeamId,
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.GET, ListOnlineTeamMembers, pathParams, nil, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*ListOnlineTeamMembersResponseV2](apiResponse)
}

// BatchQueryTeamOnlineMembersCount 批量查询群在线成员数量
func (s *TeamV2ServiceImpl) BatchQueryTeamOnlineMembersCount(req *BatchQueryTeamOnlineMembersCountRequestV2) (*core.Result[*BatchQueryTeamOnlineMembersCountResponseV2], error) {
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.POST, BatchQueryTeamOnlineMembersCount, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*BatchQueryTeamOnlineMembersCountResponseV2](apiResponse)
}
