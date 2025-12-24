package team_member

import (
	"encoding/json"
	"strconv"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/util"
)

// InviteTeamMembers 邀请成员加入群组
func (s *TeamMemberV2ServiceImpl) InviteTeamMembers(req *InviteTeamMembersRequestV2) (*core.Result[*InviteTeamMembersResponseV2], error) {
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.POST, InviteTeamMembers, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*InviteTeamMembersResponseV2](apiResponse)
}

// KickTeamMembers 踢出群成员
func (s *TeamMemberV2ServiceImpl) KickTeamMembers(req *KickTeamMembersRequestV2) (*core.Result[*KickTeamMembersResponseV2], error) {
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.POST, KickTeamMembers, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*KickTeamMembersResponseV2](apiResponse)
}

// LeaveTeam 退出群组
func (s *TeamMemberV2ServiceImpl) LeaveTeam(req *LeaveTeamRequestV2) (*core.Result[*LeaveTeamResponseV2], error) {
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.POST, LeaveTeam, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*LeaveTeamResponseV2](apiResponse)
}

// UpdateTeamMember 更新群成员信息
func (s *TeamMemberV2ServiceImpl) UpdateTeamMember(req *UpdateTeamMemberRequestV2) (*core.Result[*UpdateTeamMemberResponseV2], error) {
	pathParams := map[string]string{
		"account_id": req.AccountId,
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.PUT, UpdateTeamMember, pathParams, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*UpdateTeamMemberResponseV2](apiResponse)
}

// BatchMuteTeamMembers 批量禁言群成员
func (s *TeamMemberV2ServiceImpl) BatchMuteTeamMembers(req *BatchMuteTeamMembersRequestV2) (*core.Result[*BatchMuteTeamMembersResponseV2], error) {
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.POST, BatchMuteTeamMembers, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*BatchMuteTeamMembersResponseV2](apiResponse)
}

// QueryJoinedTeams 查询用户已加入的群组
func (s *TeamMemberV2ServiceImpl) QueryJoinedTeams(req *QueryJoinedTeamsRequestV2) (*core.Result[*QueryJoinedTeamsResponseV2], error) {
	pathParams := map[string]string{
		"account_id": req.AccountId,
	}

	queryParams := map[string]string{
		"team_type": strconv.Itoa(req.TeamType),
	}

	if req.PageToken != "" {
		queryParams["page_token"] = req.PageToken
	}
	if req.Limit > 0 {
		queryParams["limit"] = strconv.Itoa(req.Limit)
	}
	if req.NeedReturnMemberInfo > 0 {
		queryParams["need_return_member_info"] = strconv.Itoa(req.NeedReturnMemberInfo)
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.GET, JoinedTeams, pathParams, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*QueryJoinedTeamsResponseV2](apiResponse)
}
