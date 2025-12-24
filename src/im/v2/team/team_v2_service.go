package team

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
)

// TeamV2Service 群组服务接口
type TeamV2Service interface {
	// CreateTeam 创建群组
	CreateTeam(req *CreateTeamRequestV2) (*core.Result[*CreateTeamResponseV2], error)

	// UpdateTeam 更新群组信息
	UpdateTeam(req *UpdateTeamRequestV2) (*core.Result[*UpdateTeamResponseV2], error)

	// DisbandTeam 解散群组
	DisbandTeam(req *DisbandTeamRequestV2) (*core.Result[*DisbandTeamResponseV2], error)

	// BatchQueryTeamInfo 批量查询群组信息
	BatchQueryTeamInfo(req *BatchQueryTeamInfoRequestV2) (*core.Result[*BatchQueryTeamInfoResponseV2], error)

	// TransferOwner 转让群主
	TransferOwner(req *TransferTeamOwnerRequestV2) (*core.Result[*TransferTeamOwnerResponseV2], error)

	// AddManagers 添加群管理员
	AddManagers(req *AddTeamManagersRequestV2) (*core.Result[*AddTeamManagersResponseV2], error)

	// RemoveManagers 移除群管理员
	RemoveManagers(req *RemoveTeamManagersRequestV2) (*core.Result[*RemoveTeamManagersResponseV2], error)

	// GetTeamInfo 获取群组信息
	GetTeamInfo(req *GetTeamInfoRequestV2) (*core.Result[*GetTeamInfoResponseV2], error)

	// ListTeamMembers 查询群成员列表
	ListTeamMembers(req *ListTeamMembersRequestV2) (*core.Result[*ListTeamMembersResponseV2], error)

	// ListOnlineTeamMembers 查询群在线成员列表
	ListOnlineTeamMembers(req *ListOnlineTeamMembersRequestV2) (*core.Result[*ListOnlineTeamMembersResponseV2], error)

	// BatchQueryTeamOnlineMembersCount 批量查询群在线成员数量
	BatchQueryTeamOnlineMembersCount(req *BatchQueryTeamOnlineMembersCountRequestV2) (*core.Result[*BatchQueryTeamOnlineMembersCountResponseV2], error)
}

// TeamV2ServiceImpl 群组服务实现
type TeamV2ServiceImpl struct {
	httpClient core.YunxinApiHttpClient
}

// NewTeamV2Service 创建群组服务实例
func NewTeamV2Service(httpClient core.YunxinApiHttpClient) TeamV2Service {
	return &TeamV2ServiceImpl{
		httpClient: httpClient,
	}
}
