package signal

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
)

// SignalV2Service 信令服务接口 V2
type SignalV2Service interface {
	// CreateRoom 创建信令房间
	CreateRoom(req *CreateSignalRoomRequestV2) (*core.Result[*CreateSignalRoomResponseV2], error)

	// DelayRoom 延长信令房间有效期
	DelayRoom(req *DelaySignalRoomRequestV2) (*core.Result[*DelaySignalRoomResponseV2], error)

	// CloseRoom 关闭信令房间
	CloseRoom(req *CloseSignalRoomRequestV2) (*core.Result[*CloseSignalRoomResponseV2], error)

	// QueryRoom 查询信令房间
	QueryRoom(req *QuerySignalRoomRequestV2) (*core.Result[*QuerySignalRoomResponseV2], error)

	// SendControl 信令房间发送控制指令
	SendControl(req *SendSignalRoomControlRequestV2) (*core.Result[*SendSignalRoomControlResponseV2], error)

	// Invite 邀请加入信令房间
	Invite(req *InviteSignalRoomRequestV2) (*core.Result[*InviteSignalRoomResponseV2], error)

	// CancelInvite 取消邀请加入信令房间
	CancelInvite(req *CancelSignalRoomInviteRequestV2) (*core.Result[*CancelSignalRoomInviteResponseV2], error)

	// KickMember 将成员踢出信令房间
	KickMember(req *KickSignalRoomMemberRequestV2) (*core.Result[*KickSignalRoomMemberResponseV2], error)
}

// SignalV2ServiceImpl 信令服务实现 V2
type SignalV2ServiceImpl struct {
	httpClient core.YunxinApiHttpClient
}

// NewSignalV2Service 创建信令服务实例
func NewSignalV2Service(httpClient core.YunxinApiHttpClient) SignalV2Service {
	return &SignalV2ServiceImpl{
		httpClient: httpClient,
	}
}
