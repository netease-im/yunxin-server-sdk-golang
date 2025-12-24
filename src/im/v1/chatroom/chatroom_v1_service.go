package chatroom

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
)

// ChatroomV1Service 聊天室服务接口
type ChatroomV1Service interface {
	// CreateChatroom 创建聊天室
	CreateChatroom(req *CreateChatroomRequestV1) (*core.Result[*CreateChatroomResponseV1], error)

	// UpdateChatroom 更新聊天室信息
	UpdateChatroom(req *UpdateChatroomRequestV1) (*core.Result[*UpdateChatroomResponseV1], error)

	// QueryChatroomAddress 查询聊天室地址
	QueryChatroomAddress(req *QueryChatroomAddressRequestV1) (*core.Result[*QueryChatroomAddressResponseV1], error)

	// QueryChatroomInfo 获取聊天室信息
	QueryChatroomInfo(req *QueryChatroomInfoRequestV1) (*core.Result[*QueryChatroomInfoResponseV1], error)

	// QueryChatroomInfos 批量获取聊天室信息
	QueryChatroomInfos(req *QueryChatroomInfosRequestV1) (*core.Result[*QueryChatroomInfosResponseV1], error)

	// ToggleCloseChatroomStat 切换聊天室关闭状态
	ToggleCloseChatroomStat(req *ToggleCloseChatroomStatRequestV1) (*core.Result[*ToggleCloseChatroomStatResponseV1], error)

	// UpdateDelayClosePolicy 更新聊天室延迟关闭策略
	UpdateDelayClosePolicy(req *UpdateChatroomDelayClosePolicyRequestV1) (*core.Result[*UpdateChatroomDelayClosePolicyResponseV1], error)

	// UpdateChatroomInOutNotification 更新聊天室进出通知
	UpdateChatroomInOutNotification(req *UpdateChatroomInOutNotificationRequestV1) (*core.Result[*UpdateChatroomInOutNotificationResponseV1], error)

	// KickMember 踢出聊天室成员
	KickMember(req *KickMemberRequestV1) (*core.Result[*KickMemberResponseV1], error)

	// SetMemberRole 设置聊天室成员角色
	SetMemberRole(req *SetMemberRoleRequestV1) (*core.Result[*SetMemberRoleResponseV1], error)

	// UpdateMyRoomRole 更新我的聊天室角色
	UpdateMyRoomRole(req *UpdateMyRoomRoleRequestV1) (*core.Result[*UpdateMyRoomRoleResponseV1], error)

	// MembersByPage 分页获取聊天室成员
	MembersByPage(req *QueryMembersByPageRequestV1) (*core.Result[*QueryMembersByPageResponseV1], error)

	// MembersByRoles 按角色获取聊天室成员
	MembersByRoles(req *QueryMembersByRolesRequestV1) (*core.Result[*QueryMembersByRolesResponseV1], error)

	// QueryMembers 查询聊天室成员
	QueryMembers(req *QueryMembersRequestV1) (*core.Result[*QueryMembersResponseV1], error)

	// AddRobot 添加机器人
	AddRobot(req *AddRobotRequestV1) (*core.Result[*AddRobotResponseV1], error)

	// RemoveRobot 移除机器人
	RemoveRobot(req *RemoveRobotRequestV1) (*core.Result[*RemoveRobotResponseV1], error)

	// CleanRobot 清空机器人
	CleanRobot(req *CleanRobotRequestV1) (*core.Result[*CleanRobotResponseV1], error)

	// TemporaryMute 设置临时禁言状态
	TemporaryMute(req *TemporaryMuteRequestV1) (*core.Result[*TemporaryMuteResponseV1], error)

	// MuteRoom 将聊天室整体禁言
	MuteRoom(req *MuteRoomRequestV1) (*core.Result[*MuteRoomResponseV1], error)

	// TagTemporaryMute 标签临时禁言
	TagTemporaryMute(req *TagTemporaryMuteRequestV1) (*core.Result[*TagTemporaryMuteResponseV1], error)

	// TagMembersCount 标签成员计数
	TagMembersCount(req *TagMembersCountRequestV1) (*core.Result[*TagMembersCountResponseV1], error)

	// TagMembersQuery 标签成员查询
	TagMembersQuery(req *TagMembersQueryRequestV1) (*core.Result[*TagMembersQueryResponseV1], error)

	// QueryTagHistoryMsg 查询标签历史消息
	QueryTagHistoryMsg(req *QueryTagHistoryMsgRequestV1) (*core.Result[*QueryTagHistoryMsgResponseV1], error)

	// UpdateChatRoomRoleTag 更新聊天室角色标签
	UpdateChatRoomRoleTag(req *UpdateChatRoomRoleTagRequestV1) (*core.Result[*UpdateChatRoomRoleTagResponseV1], error)

	// QueryUserRoomIds 查询开放状态的聊天室
	QueryUserRoomIds(req *QueryUserRoomIdsRequestV1) (*core.Result[*QueryUserRoomIdsResponseV1], error)

	// QueueInit 初始化队列
	QueueInit(req *QueueInitRequestV1) (*core.Result[*QueueInitResponseV1], error)

	// QueueDrop 删除清理队列
	QueueDrop(req *QueueDropRequestV1) (*core.Result[*QueueDropResponseV1], error)

	// QueueOffer 新加元素到队列
	QueueOffer(req *QueueOfferRequestV1) (*core.Result[*QueueOfferResponseV1], error)

	// QueueBatchOffer 批量添加队列元素
	QueueBatchOffer(req *QueueBatchOfferRequestV1) (*core.Result[*QueueBatchOfferResponseV1], error)

	// QueueBatchUpdate 批量更新队列元素
	QueueBatchUpdate(req *QueueBatchUpdateRequestV1) (*core.Result[*QueueBatchUpdateResponseV1], error)

	// QueueList 排序列出队列中所有元素
	QueueList(req *QueueListRequestV1) (*core.Result[*QueueListResponseV1], error)

	// QueuePoll 从队列中取出元素
	QueuePoll(req *QueuePollRequestV1) (*core.Result[*QueuePollResponseV1], error)

	// QueueGet 获取指定的队列元素
	QueueGet(req *QueueGetRequestV1) (*core.Result[*QueueGetResponseV1], error)
}

// ChatroomV1ServiceImpl 聊天室服务实现
type ChatroomV1ServiceImpl struct {
	httpClient core.YunxinApiHttpClient
}

// NewChatroomV1Service 创建聊天室服务实例
func NewChatroomV1Service(httpClient core.YunxinApiHttpClient) ChatroomV1Service {
	return &ChatroomV1ServiceImpl{
		httpClient: httpClient,
	}
}
