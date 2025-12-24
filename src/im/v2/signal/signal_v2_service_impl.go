package signal

import (
	"encoding/json"
	"fmt"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/util"
)

// CreateRoom 创建信令房间
func (s *SignalV2ServiceImpl) CreateRoom(req *CreateSignalRoomRequestV2) (*core.Result[*CreateSignalRoomResponseV2], error) {
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.POST, CreateRoom, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*CreateSignalRoomResponseV2](apiResponse)
}

// DelayRoom 延长信令房间有效期
func (s *SignalV2ServiceImpl) DelayRoom(req *DelaySignalRoomRequestV2) (*core.Result[*DelaySignalRoomResponseV2], error) {
	queryParams := map[string]string{
		"channel_id": req.ChannelId,
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.PUT, DelayRoom, nil, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*DelaySignalRoomResponseV2](apiResponse)
}

// CloseRoom 关闭信令房间
func (s *SignalV2ServiceImpl) CloseRoom(req *CloseSignalRoomRequestV2) (*core.Result[*CloseSignalRoomResponseV2], error) {
	queryParams := map[string]string{
		"channel_id":          req.ChannelId,
		"operator_account_id": req.OperatorAccountId,
	}

	if req.ServerExtension != "" {
		queryParams["server_extension"] = req.ServerExtension
	}
	if req.OfflineEnabled != nil {
		queryParams["offline_enabled"] = fmt.Sprintf("%t", *req.OfflineEnabled)
	}
	if req.RouteEnabled != nil {
		queryParams["route_enabled"] = fmt.Sprintf("%t", *req.RouteEnabled)
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.DELETE, CloseRoom, nil, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*CloseSignalRoomResponseV2](apiResponse)
}

// QueryRoom 查询信令房间
func (s *SignalV2ServiceImpl) QueryRoom(req *QuerySignalRoomRequestV2) (*core.Result[*QuerySignalRoomResponseV2], error) {
	queryParams := map[string]string{}

	// channel_id和channel_name二选一，优先使用channel_id
	if req.ChannelId != "" {
		queryParams["channel_id"] = req.ChannelId
	} else if req.ChannelName != "" {
		queryParams["channel_name"] = req.ChannelName
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.GET, QueryRoom, nil, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*QuerySignalRoomResponseV2](apiResponse)
}

// SendControl 信令房间发送控制指令
func (s *SignalV2ServiceImpl) SendControl(req *SendSignalRoomControlRequestV2) (*core.Result[*SendSignalRoomControlResponseV2], error) {
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.POST, SendControl, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*SendSignalRoomControlResponseV2](apiResponse)
}

// Invite 邀请加入信令房间
func (s *SignalV2ServiceImpl) Invite(req *InviteSignalRoomRequestV2) (*core.Result[*InviteSignalRoomResponseV2], error) {
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.POST, Invite, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*InviteSignalRoomResponseV2](apiResponse)
}

// CancelInvite 取消邀请加入信令房间
func (s *SignalV2ServiceImpl) CancelInvite(req *CancelSignalRoomInviteRequestV2) (*core.Result[*CancelSignalRoomInviteResponseV2], error) {
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.POST, CancelInvite, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*CancelSignalRoomInviteResponseV2](apiResponse)
}

// KickMember 将成员踢出信令房间
func (s *SignalV2ServiceImpl) KickMember(req *KickSignalRoomMemberRequestV2) (*core.Result[*KickSignalRoomMemberResponseV2], error) {
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.POST, KickMember, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*KickSignalRoomMemberResponseV2](apiResponse)
}
