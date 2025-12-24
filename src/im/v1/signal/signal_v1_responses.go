package signal

// Signal V1 Response Structures

// CancelSignalRoomInviteResponseV1 CancelSignalRoomInvite响应
type CancelSignalRoomInviteResponseV1 struct {
}

// CloseSignalRoomResponseV1 CloseSignalRoom响应
type CloseSignalRoomResponseV1 struct {
}

// CreateSignalRoomResponseV1 CreateSignalRoom响应
type CreateSignalRoomResponseV1 struct {
	ChannelName       string   `json:"channelName,omitempty"`
	Type              string   `json:"type,omitempty"`
	ChannelId         string   `json:"channelId,omitempty"`
	ChannelCreateTime int64    `json:"channelCreateTime,omitempty"`
	ChannelExpireTime int64    `json:"channelExpireTime,omitempty"`
	Creator           string   `json:"creator,omitempty"`
	Ext               string   `json:"ext,omitempty"`
	Members           []Member `json:"members,omitempty"`
}

// Member Member
type Member struct {
	Accid      string `json:"accid,omitempty"`
	Uid        int64  `json:"uid,omitempty"`
	DeviceId   string `json:"deviceId,omitempty"`
	ExpireTime int64  `json:"expireTime,omitempty"`
	CreateTime int64  `json:"createTime,omitempty"`
}

// CtrlSignalRoomResponseV1 CtrlSignalRoom响应
type CtrlSignalRoomResponseV1 struct {
}

// DelaySignalRoomResponseV1 DelaySignalRoom响应
type DelaySignalRoomResponseV1 struct {
	ChannelName       string   `json:"channelName,omitempty"`
	Type              string   `json:"type,omitempty"`
	ChannelId         string   `json:"channelId,omitempty"`
	ChannelCreateTime int64    `json:"channelCreateTime,omitempty"`
	ChannelExpireTime int64    `json:"channelExpireTime,omitempty"`
	Creator           string   `json:"creator,omitempty"`
	Ext               string   `json:"ext,omitempty"`
	Members           []Member `json:"members,omitempty"`
}

// GetSignalRoomInfoResponseV1 GetSignalRoomInfo响应
type GetSignalRoomInfoResponseV1 struct {
	ChannelName       string   `json:"channelName,omitempty"`
	Type              string   `json:"type,omitempty"`
	ChannelId         string   `json:"channelId,omitempty"`
	ChannelCreateTime int64    `json:"channelCreateTime,omitempty"`
	ChannelExpireTime int64    `json:"channelExpireTime,omitempty"`
	Creator           string   `json:"creator,omitempty"`
	Ext               string   `json:"ext,omitempty"`
	Members           []Member `json:"members,omitempty"`
}

// InviteSignalRoomResponseV1 InviteSignalRoom响应
type InviteSignalRoomResponseV1 struct {
}

// KickSignalRoomResponseV1 KickSignalRoom响应
type KickSignalRoomResponseV1 struct {
}
