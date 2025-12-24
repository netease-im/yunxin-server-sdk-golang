package team_member

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
)

// TeamMemberV2Service 群成员服务接口
type TeamMemberV2Service interface {
	// InviteTeamMembers 邀请成员加入群组
	InviteTeamMembers(req *InviteTeamMembersRequestV2) (*core.Result[*InviteTeamMembersResponseV2], error)

	// KickTeamMembers 踢出群成员
	KickTeamMembers(req *KickTeamMembersRequestV2) (*core.Result[*KickTeamMembersResponseV2], error)

	// LeaveTeam 退出群组
	LeaveTeam(req *LeaveTeamRequestV2) (*core.Result[*LeaveTeamResponseV2], error)

	// UpdateTeamMember 更新群成员信息
	UpdateTeamMember(req *UpdateTeamMemberRequestV2) (*core.Result[*UpdateTeamMemberResponseV2], error)

	// BatchMuteTeamMembers 批量禁言群成员
	BatchMuteTeamMembers(req *BatchMuteTeamMembersRequestV2) (*core.Result[*BatchMuteTeamMembersResponseV2], error)

	// QueryJoinedTeams 查询用户已加入的群组
	QueryJoinedTeams(req *QueryJoinedTeamsRequestV2) (*core.Result[*QueryJoinedTeamsResponseV2], error)
}

// TeamMemberV2ServiceImpl 群成员服务实现
type TeamMemberV2ServiceImpl struct {
	httpClient core.YunxinApiHttpClient
}

// NewTeamMemberV2Service 创建群成员服务实例
func NewTeamMemberV2Service(httpClient core.YunxinApiHttpClient) TeamMemberV2Service {
	return &TeamMemberV2ServiceImpl{
		httpClient: httpClient,
	}
}
