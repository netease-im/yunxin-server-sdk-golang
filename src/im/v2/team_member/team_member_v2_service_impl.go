package team_member

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/util"
)

// joinStringSlice joins a string slice with a separator
func joinStringSlice(slice []string, sep string) string {
	return strings.Join(slice, sep)
}

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
	// Use query parameters for DELETE request
	queryParams := map[string]string{
		"operator_id":      req.OperatorId,
		"team_id":          strconv.FormatInt(req.TeamId, 10),
		"team_type":        strconv.Itoa(*req.TeamType),
		"kick_account_ids": joinStringSlice(req.KickAccountIds, ","),
	}

	if req.Extension != "" {
		queryParams["extension"] = req.Extension
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.DELETE, KickTeamMembers, nil, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*KickTeamMembersResponseV2](apiResponse)
}

// LeaveTeam 退出群组
func (s *TeamMemberV2ServiceImpl) LeaveTeam(req *LeaveTeamRequestV2) (*core.Result[*LeaveTeamResponseV2], error) {
	// Use query parameters for DELETE request
	queryParams := map[string]string{
		"account_id": req.AccountId,
		"team_id":    strconv.FormatInt(req.TeamId, 10),
		"team_type":  strconv.Itoa(*req.TeamType),
	}

	if req.Extension != "" {
		queryParams["extension"] = req.Extension
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.DELETE, LeaveTeam, nil, queryParams, "")
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

	apiResponse, err := s.httpClient.ExecuteV2Api(http.PATCH, UpdateTeamMember, pathParams, nil, string(requestBody))
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

	apiResponse, err := s.httpClient.ExecuteV2Api(http.PATCH, BatchMuteTeamMembers, nil, nil, string(requestBody))
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
		"team_type": strconv.Itoa(*req.TeamType),
	}

	if req.PageToken != "" {
		queryParams["page_token"] = req.PageToken
	}
	if req.Limit > 0 {
		queryParams["limit"] = strconv.Itoa(req.Limit)
	}
	if req.NeedReturnMemberInfo != nil && *req.NeedReturnMemberInfo > 0 {
		queryParams["need_return_member_info"] = strconv.Itoa(*req.NeedReturnMemberInfo)
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.GET, JoinedTeams, pathParams, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*QueryJoinedTeamsResponseV2](apiResponse)
}
