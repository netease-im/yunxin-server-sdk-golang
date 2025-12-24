package chatroom_member

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
)

// ChatroomMemberV2Service Chatroom Member V2服务接口
type ChatroomMemberV2Service interface {
	// SetMemberRole 设置成员角色
	SetMemberRole(req *SetMemberRoleRequestV2) (*core.Result[*SetMemberRoleResponseV2], error)

	// UpdateOnlineMemberInfo 更新在线成员信息
	UpdateOnlineMemberInfo(req *UpdateOnlineMemberInfoRequestV2) (*core.Result[*UpdateOnlineMemberInfoResponseV2], error)

	// ToggleChatBan 切换禁言
	ToggleChatBan(req *ToggleChatBanRequestV2) (*core.Result[*ToggleChatBanResponseV2], error)

	// ToggleTempChatBan 切换临时禁言
	ToggleTempChatBan(req *ToggleTempChatBanRequestV2) (*core.Result[*ToggleTempChatBanResponseV2], error)

	// ToggleBlocked 切换屏蔽
	ToggleBlocked(req *ToggleBlockedRequestV2) (*core.Result[*ToggleBlockedResponseV2], error)

	// ModifyMemberTags 修改成员标签
	ModifyMemberTags(req *ModifyMemberTagsRequestV2) (*core.Result[*ModifyMemberTagsResponseV2], error)

	// QueryTaggedMembersCount 查询标签成员数量
	QueryTaggedMembersCount(req *QueryTaggedMembersCountRequestV2) (*core.Result[*QueryTaggedMembersCountResponseV2], error)

	// ListTaggedMembers 列出标签成员
	ListTaggedMembers(req *ListTaggedMembersRequestV2) (*core.Result[*ListTaggedMembersResponseV2], error)

	// QueryChatroomBlacklist 查询聊天室黑名单
	QueryChatroomBlacklist(req *QueryChatroomBlacklistRequestV2) (*core.Result[*QueryChatroomBlacklistResponseV2], error)

	// ToggleTaggedMembersChatBan 切换标签成员禁言
	ToggleTaggedMembersChatBan(req *ToggleTaggedMembersChatBanRequestV2) (*core.Result[*ToggleTaggedMembersChatBanResponseV2], error)

	// BatchQueryChatroomMembers 批量查询聊天室成员
	BatchQueryChatroomMembers(req *BatchQueryChatroomMembersRequestV2) (*core.Result[*BatchQueryChatroomMembersResponseV2], error)

	// AddVirtualMembers 添加虚拟成员
	AddVirtualMembers(req *AddVirtualMembersRequestV2) (*core.Result[*AddVirtualMembersResponseV2], error)

	// DeleteVirtualMembers 删除虚拟成员
	DeleteVirtualMembers(req *DeleteVirtualMembersRequestV2) (*core.Result[*DeleteVirtualMembersResponseV2], error)

	// ClearVirtualMembers 清空虚拟成员
	ClearVirtualMembers(req *ClearVirtualMembersRequestV2) (*core.Result[*ClearVirtualMembersResponseV2], error)

	// QueryVirtualMembers 查询虚拟成员
	QueryVirtualMembers(req *QueryVirtualMembersRequestV2) (*core.Result[*QueryVirtualMembersResponseV2], error)

	// QueryChatBanned 查询禁言成员
	QueryChatBanned(req *QueryChatBannedRequestV2) (*core.Result[*QueryChatBannedResponseV2], error)
}

// ChatroomMemberV2ServiceImpl Chatroom Member V2服务实现
type ChatroomMemberV2ServiceImpl struct {
	httpClient core.YunxinApiHttpClient
}

// NewChatroomMemberV2Service 创建Chatroom Member V2服务实例
func NewChatroomMemberV2Service(httpClient core.YunxinApiHttpClient) ChatroomMemberV2Service {
	return &ChatroomMemberV2ServiceImpl{
		httpClient: httpClient,
	}
}
