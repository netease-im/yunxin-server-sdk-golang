package team

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
)

// TeamV1Service team服务接口
type TeamV1Service interface {
	// CreateTeam CreateTeam
	CreateTeam(req *CreateTeamRequestV1) (*core.Result[*CreateTeamResponseV1], error)

	// AddTeam AddTeam
	AddTeam(req *AddTeamRequestV1) (*core.Result[*AddTeamResponseV1], error)

	// KickTeam KickTeam
	KickTeam(req *KickTeamRequestV1) (*core.Result[*KickTeamResponseV1], error)

	// LeaveTeam LeaveTeam
	LeaveTeam(req *LeaveTeamRequestV1) (*core.Result[*LeaveTeamResponseV1], error)

	// UpdateTeam UpdateTeam
	UpdateTeam(req *UpdateTeamRequestV1) (*core.Result[*UpdateTeamResponseV1], error)

	// QueryTeam QueryTeam
	QueryTeam(req *QueryTeamRequestV1) (*core.Result[*QueryTeamResponseV1], error)

	// QueryTeamInfoDetails QueryTeamInfoDetails
	QueryTeamInfoDetails(req *QueryTeamInfoDetailsRequestV1) (*core.Result[*QueryTeamInfoDetailsResponseV1], error)

	// AddManagerTeam AddManagerTeam
	AddManagerTeam(req *AddManagerTeamRequestV1) (*core.Result[*AddManagerTeamResponseV1], error)

	// RemoveManagerTeam RemoveManagerTeam
	RemoveManagerTeam(req *RemoveManagerTeamRequestV1) (*core.Result[*RemoveManagerTeamResponseV1], error)

	// ChangeOwnerTeam ChangeOwnerTeam
	ChangeOwnerTeam(req *ChangeOwnerTeamRequestV1) (*core.Result[*ChangeOwnerTeamResponseV1], error)

	// DismissTeam DismissTeam
	DismissTeam(req *DismissTeamRequestV1) (*core.Result[*DismissTeamResponseV1], error)

	// MuteTeam MuteTeam
	MuteTeam(req *MuteTeamRequestV1) (*core.Result[*MuteTeamResponseV1], error)

	// MuteTeamAllMember MuteTeamAllMember
	MuteTeamAllMember(req *MuteTeamAllMemberRequestV1) (*core.Result[*MuteTeamAllMemberResponseV1], error)

	// MuteTeamTargetMember MuteTeamTargetMember
	MuteTeamTargetMember(req *MuteTeamTargetMemberRequestV1) (*core.Result[*MuteTeamTargetMemberResponseV1], error)

	// QueryMuteTeamMembers QueryMuteTeamMembers
	QueryMuteTeamMembers(req *QueryMuteTeamMembersRequestV1) (*core.Result[*QueryMuteTeamMembersResponseV1], error)

	// UpdateTeamNick UpdateTeamNick
	UpdateTeamNick(req *UpdateTeamNickRequestV1) (*core.Result[*UpdateTeamNickResponseV1], error)

	// JoinsTeam JoinsTeam
	JoinsTeam(req *JoinsTeamRequestV1) (*core.Result[*JoinsTeamResponseV1], error)

	// QueryOnlineTeamMember QueryOnlineTeamMember
	QueryOnlineTeamMember(req *QueryOnlineTeamMemberRequestV1) (*core.Result[*QueryOnlineTeamMemberResponseV1], error)

	// BatchQueryOnlineTeamMemberCount BatchQueryOnlineTeamMemberCount
	BatchQueryOnlineTeamMemberCount(req *BatchQueryOnlineTeamMemberCountRequestV1) (*core.Result[*BatchQueryOnlineTeamMemberCountResponseV1], error)

	// QueryTeamMsgMarkReadInfo QueryTeamMsgMarkReadInfo
	QueryTeamMsgMarkReadInfo(req *QueryTeamMsgMarkReadInfoRequestV1) (*core.Result[*QueryTeamMsgMarkReadInfoResponseV1], error)

	// QueryAllJoinedTeamMemberInfoByAccId QueryAllJoinedTeamMemberInfoByAccId
	QueryAllJoinedTeamMemberInfoByAccId(req *QueryAllJoinedTeamMemberInfoByAccIdRequestV1) (*core.Result[*QueryAllJoinedTeamMemberInfoByAccIdResponseV1], error)
}

// TeamV1ServiceImpl team服务实现
type TeamV1ServiceImpl struct {
	httpClient core.YunxinApiHttpClient
}

// NewTeamV1Service 创建team服务实例
func NewTeamV1Service(httpClient core.YunxinApiHttpClient) TeamV1Service {
	return &TeamV1ServiceImpl{
		httpClient: httpClient,
	}
}
