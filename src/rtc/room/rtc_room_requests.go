package room

// RtcCreateRoomRequest 创建房间请求
type RtcCreateRoomRequest struct {
	ChannelName string `json:"channelName,omitempty"` // 房间名称
	Mode        int    `json:"mode,omitempty"`        // 模式
	Uid         uint64 `json:"uid,omitempty"`         // 用户ID
}

// RtcGetRoomByCidRequest 根据CID获取房间信息请求
type RtcGetRoomByCidRequest struct {
	Cid int64 `json:"cid"` // 房间ID
}

// RtcGetRoomByCnameRequest 根据CNAME获取房间信息请求
type RtcGetRoomByCnameRequest struct {
	Cname string `json:"cname"` // 房间名称
}

// RtcListRoomMembersRequestV2 V2版本列出房间成员请求
// See https://doc.yunxin.163.com/nertc/server-apis/jUzODcwODE?platform=server
type RtcListRoomMembersRequestV2 struct {
	Cid      int64  `json:"cid"`                // 房间ID
	Uid      uint64 `json:"uid,omitempty"`      // 用户ID
	UserRole *int   `json:"userRole,omitempty"` // 用户角色
}

// RtcListRoomMembersRequestV3 V3版本列出房间成员请求
// See https://doc.yunxin.163.com/nertc/server-apis/jUzODcwODE?platform=server
type RtcListRoomMembersRequestV3 struct {
	Cname    string `json:"cname"`              // 房间名称
	Uid      uint64 `json:"uid,omitempty"`      // 用户ID
	UserRole *int   `json:"userRole,omitempty"` // 用户角色
}

// RtcAddMemberToKicklistRequestV2 V2版本添加成员到踢出列表请求
// See https://doc.yunxin.163.com/nertc/server-apis/TA5MTExNjg?platform=server
type RtcAddMemberToKicklistRequestV2 struct {
	Cid      int64  `json:"cid"`                // 房间ID
	Uid      uint64 `json:"uid"`                // 用户ID
	Duration int64  `json:"duration,omitempty"` // 时长
}

// RtcAddMemberToKicklistRequestV3 V3版本添加成员到踢出列表请求
// See https://doc.yunxin.163.com/nertc/server-apis/TA5MTExNjg?platform=server
type RtcAddMemberToKicklistRequestV3 struct {
	Cname    string `json:"cname"`              // 房间名称
	Uid      uint64 `json:"uid"`                // 用户ID
	Duration int64  `json:"duration,omitempty"` // 时长
}

// RtcDeleteRoomRequestV2 V2版本删除房间请求
// See https://doc.yunxin.163.com/nertc/server-apis/DM2NzQyMTc?platform=server
type RtcDeleteRoomRequestV2 struct {
	Cid int64 `json:"cid"` // 房间ID
}

// RtcDeleteRoomRequestV3 V3版本删除房间请求
// See https://doc.yunxin.163.com/nertc/server-apis/DM2NzQyMTc?platform=server
type RtcDeleteRoomRequestV3 struct {
	Cname string `json:"cname"` // 房间名称
}
