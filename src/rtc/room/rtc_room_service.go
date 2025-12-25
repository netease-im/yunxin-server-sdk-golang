package room

import . "github.com/netease-im/yunxin-server-sdk-golang/src/rtc/result"

// RtcRoomService RTC房间服务接口
type RtcRoomService interface {
	// CreateRoom 创建房间
	CreateRoom(request *RtcCreateRoomRequest) (*RtcResult[*RtcCreateRoomResponse], error)

	// GetRoomByCid 根据CID获取房间信息
	GetRoomByCid(request *RtcGetRoomByCidRequest) (*RtcResult[*RtcGetRoomResponse], error)

	// GetRoomByCname 根据CNAME获取房间信息
	GetRoomByCname(request *RtcGetRoomByCnameRequest) (*RtcResult[*RtcGetRoomResponse], error)

	// ListRoomMembersV2 V2版本列出房间成员
	ListRoomMembersV2(request *RtcListRoomMembersRequestV2) (*RtcResult[*RtcListRoomMembersResponse], error)

	// ListRoomMembersV3 V3版本列出房间成员
	ListRoomMembersV3(request *RtcListRoomMembersRequestV3) (*RtcResult[*RtcListRoomMembersResponse], error)

	// AddMemberToKicklistV2 V2版本添加成员到踢出列表
	AddMemberToKicklistV2(request *RtcAddMemberToKicklistRequestV2) (*RtcResult[*RtcAddMemberToKicklistResponse], error)

	// AddMemberToKicklistV3 V3版本添加成员到踢出列表
	AddMemberToKicklistV3(request *RtcAddMemberToKicklistRequestV3) (*RtcResult[*RtcAddMemberToKicklistResponse], error)

	// DeleteRoomV2 V2版本删除房间
	DeleteRoomV2(request *RtcDeleteRoomRequestV2) (*RtcResult[*RtcDeleteRoomResponse], error)

	// DeleteRoomV3 V3版本删除房间
	DeleteRoomV3(request *RtcDeleteRoomRequestV3) (*RtcResult[*RtcDeleteRoomResponse], error)
}
