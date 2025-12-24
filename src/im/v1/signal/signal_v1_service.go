package signal

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
)

// SignalV1Service signal服务接口
type SignalV1Service interface {
	// CreateSignalRoom CreateSignalRoom
	CreateSignalRoom(req *CreateSignalRoomRequestV1) (*core.Result[*CreateSignalRoomResponseV1], error)

	// CloseSignalRoom CloseSignalRoom
	CloseSignalRoom(req *CloseSignalRoomRequestV1) (*core.Result[*CloseSignalRoomResponseV1], error)

	// InviteSignalRoom InviteSignalRoom
	InviteSignalRoom(req *InviteSignalRoomRequestV1) (*core.Result[*InviteSignalRoomResponseV1], error)

	// CancelSignalRoomInvite CancelSignalRoomInvite
	CancelSignalRoomInvite(req *CancelSignalRoomInviteRequestV1) (*core.Result[*CancelSignalRoomInviteResponseV1], error)

	// KickSignalRoom KickSignalRoom
	KickSignalRoom(req *KickSignalRoomRequestV1) (*core.Result[*KickSignalRoomResponseV1], error)

	// CtrlSignalRoom CtrlSignalRoom
	CtrlSignalRoom(req *CtrlSignalRoomRequestV1) (*core.Result[*CtrlSignalRoomResponseV1], error)

	// DelaySignalRoom DelaySignalRoom
	DelaySignalRoom(req *DelaySignalRoomRequestV1) (*core.Result[*DelaySignalRoomResponseV1], error)

	// GetSignalRoomInfo GetSignalRoomInfo
	GetSignalRoomInfo(req *GetSignalRoomInfoRequestV1) (*core.Result[*GetSignalRoomInfoResponseV1], error)
}

// SignalV1ServiceImpl signal服务实现
type SignalV1ServiceImpl struct {
	httpClient core.YunxinApiHttpClient
}

// NewSignalV1Service 创建signal服务实例
func NewSignalV1Service(httpClient core.YunxinApiHttpClient) SignalV1Service {
	return &SignalV1ServiceImpl{
		httpClient: httpClient,
	}
}
