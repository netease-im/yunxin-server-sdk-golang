package signal

// SignalRoomMember 信令房间成员信息
type SignalRoomMember struct {
	AccountId  string `json:"account_id"`  // 账号ID
	Uid        int64  `json:"uid"`         // 用户ID
	DeviceId   string `json:"device_id"`   // 设备ID
	JoinTime   int64  `json:"join_time"`   // 加入时间
	ExpireTime int64  `json:"expire_time"` // 过期时间
}

// SignalRoomInfo 信令房间信息
type SignalRoomInfo struct {
	ChannelName      string             `json:"channel_name"`       // 信令房间名称
	ChannelId        string             `json:"channel_id"`         // 信令房间ID
	CreatorAccountId string             `json:"creator_account_id"` // 创建者账号ID
	ChannelExtension string             `json:"channel_extension"`  // 信令房间扩展字段
	ChannelType      int                `json:"channel_type"`       // 信令房间类型
	CreateTime       int64              `json:"create_time"`        // 创建时间
	ExpireTime       int64              `json:"expire_time"`        // 过期时间
	MemberList       []SignalRoomMember `json:"member_list"`        // 成员列表
}

// CreateSignalRoomResponseV2 创建信令房间响应
type CreateSignalRoomResponseV2 struct {
	SignalRoomInfo
}

// DelaySignalRoomResponseV2 延长信令房间有效期响应
type DelaySignalRoomResponseV2 struct {
	SignalRoomInfo
}

// CloseSignalRoomResponseV2 关闭信令房间响应
type CloseSignalRoomResponseV2 struct {
	// API返回的data为空对象 {}，无需任何字段
}

// QuerySignalRoomResponseV2 查询信令房间响应
type QuerySignalRoomResponseV2 struct {
	SignalRoomInfo
}

// SendSignalRoomControlResponseV2 信令房间发送控制指令响应
type SendSignalRoomControlResponseV2 struct {
	// API返回的data为空对象 {}，无需任何字段
}

// InviteSignalRoomResponseV2 邀请加入信令房间响应
type InviteSignalRoomResponseV2 struct {
	// API返回的data为空对象 {}，无需任何字段
}

// CancelSignalRoomInviteResponseV2 取消邀请加入信令房间响应
type CancelSignalRoomInviteResponseV2 struct {
	// API返回的data为空对象 {}，无需任何字段
}

// KickSignalRoomMemberResponseV2 将成员踢出信令房间响应
type KickSignalRoomMemberResponseV2 struct {
	// API返回的data为空对象 {}，无需任何字段
}
