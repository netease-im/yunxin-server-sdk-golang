package signal

// Signal V1 Request Structures

// CancelSignalRoomInviteRequestV1 CancelSignalRoomInvite请求
type CancelSignalRoomInviteRequestV1 struct {
	ChannelId   string `json:"channelId,omitempty"`
	From        string `json:"from,omitempty"`
	To          string `json:"to,omitempty"`
	RequestId   string `json:"requestId,omitempty"`
	AttachExt   string `json:"attachExt,omitempty"`
	NeedPush    int    `json:"needPush,omitempty"`
	PushTitle   string `json:"pushTitle,omitempty"`
	PushContent string `json:"pushContent,omitempty"`
	PushPayload string `json:"pushPayload,omitempty"`
	NeedBadge   int    `json:"needBadge,omitempty"`
	IsSave      int    `json:"isSave,omitempty"`
	IsRoute     int    `json:"isRoute,omitempty"`
}

// CloseSignalRoomRequestV1 CloseSignalRoom请求
type CloseSignalRoomRequestV1 struct {
	ChannelId string `json:"channelId,omitempty"`
	From      string `json:"from,omitempty"`
	IsSave    int    `json:"isSave,omitempty"`
	AttachExt string `json:"attachExt,omitempty"`
	IsRoute   int    `json:"isRoute,omitempty"`
}

// CreateSignalRoomRequestV1 CreateSignalRoom请求
type CreateSignalRoomRequestV1 struct {
	ChannelName string `json:"channelName,omitempty"`
	Type        int    `json:"type,omitempty"`
	From        string `json:"from,omitempty"`
	Ext         string `json:"ext,omitempty"`
	IsRoute     int    `json:"isRoute,omitempty"`
}

// CtrlSignalRoomRequestV1 CtrlSignalRoom请求
type CtrlSignalRoomRequestV1 struct {
	ChannelId string `json:"channelId,omitempty"`
	From      string `json:"from,omitempty"`
	To        string `json:"to,omitempty"`
	AttachExt string `json:"attachExt,omitempty"`
	IsRoute   int    `json:"isRoute,omitempty"`
}

// DelaySignalRoomRequestV1 DelaySignalRoom请求
type DelaySignalRoomRequestV1 struct {
	ChannelId string `json:"channelId,omitempty"`
}

// GetSignalRoomInfoRequestV1 GetSignalRoomInfo请求
type GetSignalRoomInfoRequestV1 struct {
	ChannelName string `json:"channelName,omitempty"`
	ChannelId   string `json:"channelId,omitempty"`
}

// InviteSignalRoomRequestV1 InviteSignalRoom请求
type InviteSignalRoomRequestV1 struct {
	ChannelId   string `json:"channelId,omitempty"`
	From        string `json:"from,omitempty"`
	To          string `json:"to,omitempty"`
	RequestId   string `json:"requestId,omitempty"`
	AttachExt   string `json:"attachExt,omitempty"`
	NeedPush    int    `json:"needPush,omitempty"`
	PushTitle   string `json:"pushTitle,omitempty"`
	PushContent string `json:"pushContent,omitempty"`
	PushPayload string `json:"pushPayload,omitempty"`
	NeedBadge   int    `json:"needBadge,omitempty"`
	IsSave      int    `json:"isSave,omitempty"`
	IsRoute     int    `json:"isRoute,omitempty"`
}

// KickSignalRoomRequestV1 KickSignalRoom请求
type KickSignalRoomRequestV1 struct {
	ChannelId string `json:"channelId,omitempty"`
	From      string `json:"from,omitempty"`
	To        string `json:"to,omitempty"`
	AttachExt string `json:"attachExt,omitempty"`
	IsSave    int    `json:"isSave,omitempty"`
	IsRoute   int    `json:"isRoute,omitempty"`
}
