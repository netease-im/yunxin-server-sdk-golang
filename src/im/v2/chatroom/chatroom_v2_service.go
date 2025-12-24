package chatroom

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
)

// ChatroomV2Service Chatroom V2服务接口
type ChatroomV2Service interface {
	// CreateChatroom 创建聊天室
	CreateChatroom(req *CreateChatroomRequestV2) (*core.Result[*CreateChatroomResponseV2], error)

	// GetChatroomInfo 获取聊天室信息
	GetChatroomInfo(req *GetChatroomInfoRequestV2) (*core.Result[*GetChatroomInfoResponseV2], error)

	// UpdateChatroomInfo 更新聊天室信息
	UpdateChatroomInfo(req *UpdateChatroomInfoRequestV2) (*core.Result[*UpdateChatroomInfoResponseV2], error)

	// GetChatroomAddress 获取聊天室地址
	GetChatroomAddress(req *GetChatroomAddressRequestV2) (*core.Result[*GetChatroomAddressResponseV2], error)

	// UpdateChatroomStatus 更新聊天室状态
	UpdateChatroomStatus(req *UpdateChatroomStatusRequestV2) (*core.Result[*UpdateChatroomStatusResponseV2], error)

	// ToggleChatroomMute 切换聊天室禁言
	ToggleChatroomMute(req *ToggleChatroomMuteRequestV2) (*core.Result[*ToggleChatroomMuteResponseV2], error)

	// ToggleInOutNotification 切换进出通知
	ToggleInOutNotification(req *ToggleInOutNotificationRequestV2) (*core.Result[*ToggleInOutNotificationResponseV2], error)

	// QueryOpenChatrooms 查询开放聊天室
	QueryOpenChatrooms(req *QueryOpenChatroomsRequestV2) (*core.Result[*QueryOpenChatroomsResponseV2], error)

	// ListOnlineMembers 列出在线成员
	ListOnlineMembers(req *ListOnlineMembersRequestV2) (*core.Result[*ListOnlineMembersResponseV2], error)

	// ListFixedMembers 列出固定成员
	ListFixedMembers(req *ListFixedMembersRequestV2) (*core.Result[*ListFixedMembersResponseV2], error)
}

// ChatroomV2ServiceImpl Chatroom V2服务实现
type ChatroomV2ServiceImpl struct {
	httpClient core.YunxinApiHttpClient
}

// NewChatroomV2Service 创建Chatroom V2服务实例
func NewChatroomV2Service(httpClient core.YunxinApiHttpClient) ChatroomV2Service {
	return &ChatroomV2ServiceImpl{
		httpClient: httpClient,
	}
}
