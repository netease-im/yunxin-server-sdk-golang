package super_team

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
)

// SuperTeamV1Service super_team服务接口
type SuperTeamV1Service interface {
	// SuperTeamCreate SuperTeamCreate
	SuperTeamCreate(req *SuperTeamCreateRequestV1) (*core.Result[*CreateSuperTeamResponseV1], error)

	// SuperTeamUpdate SuperTeamUpdate
	SuperTeamUpdate(req *SuperTeamUpdateRequestV1) (*core.Result[*SuperTeamUpdateResponseV1], error)

	// SuperTeamDismiss SuperTeamDismiss
	SuperTeamDismiss(req *SuperTeamDismissRequestV1) (*core.Result[*SuperTeamDismissResponseV1], error)

	// SuperTeamInvite SuperTeamInvite
	SuperTeamInvite(req *SuperTeamInviteRequestV1) (*core.Result[*SuperTeamInviteResponseV1], error)

	// SuperTeamKickMember SuperTeamKickMember
	SuperTeamKickMember(req *SuperTeamKickMemberRequestV1) (*core.Result[*SuperTeamKickMemberResponseV1], error)

	// SuperTeamMemberLeave SuperTeamMemberLeave
	SuperTeamMemberLeave(req *SuperTeamMemberLeaveRequestV1) (*core.Result[*SuperTeamMemberLeaveResponseV1], error)

	// SuperTeamAddManager SuperTeamAddManager
	SuperTeamAddManager(req *SuperTeamAddManagerRequestV1) (*core.Result[*SuperTeamAddManagerResponseV1], error)

	// SuperTeamRemoveManager SuperTeamRemoveManager
	SuperTeamRemoveManager(req *SuperTeamRemoveManagerRequestV1) (*core.Result[*SuperTeamRemoveManagerResponseV1], error)

	// SuperTeamChangeOwner SuperTeamChangeOwner
	SuperTeamChangeOwner(req *SuperTeamChangeOwnerRequestV1) (*core.Result[*SuperTeamChangeOwnerResponseV1], error)

	// GetSuperTeam GetSuperTeam
	GetSuperTeam(req *GetSuperTeamRequestV1) (*core.Result[*GetSuperTeamResponseV1], error)

	// GetSuperTeamMember GetSuperTeamMember
	GetSuperTeamMember(req *GetSuperTeamMemberRequestV1) (*core.Result[*GetSuperTeamMemberResponseV1], error)

	// GetJoinSuperTeam GetJoinSuperTeam
	GetJoinSuperTeam(req *GetJoinSuperTeamRequestV1) (*core.Result[*GetJoinSuperTeamResponseV1], error)

	// SuperTeamUpdateMemberInfo SuperTeamUpdateMemberInfo
	SuperTeamUpdateMemberInfo(req *SuperTeamUpdateMemberInfoRequestV1) (*core.Result[*SuperTeamUpdateMemberInfoResponseV1], error)

	// SuperTeamUpdateNick SuperTeamUpdateNick
	SuperTeamUpdateNick(req *SuperTeamUpdateNickRequestV1) (*core.Result[*SuperTeamUpdateNickResponseV1], error)

	// SuperTeamMute SuperTeamMute
	SuperTeamMute(req *SuperTeamMuteRequestV1) (*core.Result[*SuperTeamMuteResponseV1], error)

	// SuperTeamMuteTlist SuperTeamMuteTlist
	SuperTeamMuteTlist(req *SuperTeamMuteTlistRequestV1) (*core.Result[*SuperTeamMuteTlistResponseV1], error)

	// GetSuperTeamMuteMember GetSuperTeamMuteMember
	GetSuperTeamMuteMember(req *GetSuperTeamMuteMemberRequestV1) (*core.Result[*GetSuperTeamMuteMemberResponseV1], error)

	// SuperTeamChangeLevel SuperTeamChangeLevel
	SuperTeamChangeLevel(req *SuperTeamChangeLevelRequestV1) (*core.Result[*SuperTeamChangeLevelResponseV1], error)

	// SendSuperTeamMessage SendSuperTeamMessage
	SendSuperTeamMessage(req *SendSuperTeamMessageRequestV1) (*core.Result[*SendSuperTeamMessageResponseV1], error)

	// SendAttachSuperTeamMessage SendAttachSuperTeamMessage
	SendAttachSuperTeamMessage(req *SendAttachSuperTeamMessageRequestV1) (*core.Result[*SendAttachSuperTeamMessageResponseV1], error)

	// RecallSuperTeamMessage RecallSuperTeamMessage
	RecallSuperTeamMessage(req *RecallSuperTeamMessageRequestV1) (*core.Result[*RecallSuperTeamMessageResponseV1], error)

	// GetSuperTeamMessage GetSuperTeamMessage
	GetSuperTeamMessage(req *GetSuperTeamMessageRequestV1) (*core.Result[*GetSuperTeamMessageResponseV1], error)

	// GetSuperTeamMessageByIds GetSuperTeamMessageByIds
	GetSuperTeamMessageByIds(req *GetSuperTeamMessageByIdsRequestV1) (*core.Result[*GetSuperTeamMessageByIdsResponseV1], error)
}

// SuperTeamV1ServiceImpl super_team服务实现
type SuperTeamV1ServiceImpl struct {
	httpClient core.YunxinApiHttpClient
}

// NewSuperTeamV1Service 创建super_team服务实例
func NewSuperTeamV1Service(httpClient core.YunxinApiHttpClient) SuperTeamV1Service {
	return &SuperTeamV1ServiceImpl{
		httpClient: httpClient,
	}
}
