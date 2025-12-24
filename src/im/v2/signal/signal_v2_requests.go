package signal

// RouteConfig 抄送配置
type RouteConfig struct {
	RouteEnabled bool `json:"route_enabled"` // 是否抄送，默认false
}

// PushConfig 推送配置
type PushConfig struct {
	PushEnabled bool   `json:"push_enabled"` // 是否推送，默认false
	PushTitle   string `json:"push_title"`   // 推送标题，长度上限32字符
	PushContent string `json:"push_content"` // 推送文案，长度上限500字符
	PushPayload string `json:"push_payload"` // 推送对应的 payload，必须是 JSON 格式，长度上限 4096 位字符
}

// CreateSignalRoomRequestV2 创建信令房间请求
type CreateSignalRoomRequestV2 struct {
	ChannelName      string       `json:"channel_name"`           // 信令房间名称
	CreatorAccountId string       `json:"creator_account_id"`     // 创建者账号ID
	ChannelType      int          `json:"channel_type"`           // 信令房间类型
	ChannelExtension string       `json:"channel_extension"`      // 信令房间扩展字段
	RouteConfig      *RouteConfig `json:"route_config,omitempty"` // 抄送配置
}

// DelaySignalRoomRequestV2 延长信令房间有效期请求
type DelaySignalRoomRequestV2 struct {
	ChannelId string `json:"-"` // 信令房间id,查询参数
}

// CloseSignalRoomRequestV2 关闭信令房间请求
type CloseSignalRoomRequestV2 struct {
	ChannelId         string `json:"-"` // 信令房间id,查询参数
	OperatorAccountId string `json:"-"` // 操作者，必须是信令房间创建者或者房间里的人,查询参数
	ServerExtension   string `json:"-"` // 关闭通知的扩展字段，最长4096个字符,查询参数
	OfflineEnabled    *bool  `json:"-"` // 关闭通知是否存离线，默认false,查询参数
	RouteEnabled      *bool  `json:"-"` // 是否抄送，默认false,查询参数
}

// QuerySignalRoomRequestV2 查询信令房间请求
type QuerySignalRoomRequestV2 struct {
	ChannelId   string `json:"-"` // 信令房间id,查询参数
	ChannelName string `json:"-"` // 信令房间名称,查询参数
}

// SendSignalRoomControlRequestV2 信令房间发送控制指令请求
type SendSignalRoomControlRequestV2 struct {
	ChannelId         string       `json:"channel_id"`                  // 信令房间id
	OperatorAccountId string       `json:"operator_account_id"`         // 操作者，必须是信令房间创建者或者房间里的人
	TargetAccountId   string       `json:"target_account_id,omitempty"` // 控制指令的目标成员，如果缺失，则表示发送给房间里所有人
	ServerExtension   string       `json:"server_extension,omitempty"`  // 发送通知时的扩展字段，最长4096个字符
	RouteConfig       *RouteConfig `json:"route_config,omitempty"`      // 抄送配置
}

// InviteSignalRoomRequestV2 邀请加入信令房间请求
type InviteSignalRoomRequestV2 struct {
	ChannelId        string       `json:"channel_id"`                 // 信令房间id
	InviterAccountId string       `json:"inviter_account_id"`         // 操作者，必须是信令房间创建者或者房间里的人
	InviteeAccountId string       `json:"invitee_account_id"`         // 被邀请者账号
	RequestId        string       `json:"request_id"`                 // 邀请ID，长度上限128字符
	ServerExtension  string       `json:"server_extension,omitempty"` // 发送通知时的扩展字段，最长4096个字符
	RouteConfig      *RouteConfig `json:"route_config,omitempty"`     // 抄送配置
	OfflineEnabled   bool         `json:"offline_enabled"`            // 关闭通知是否存离线，默认false
	UnreadEnabled    bool         `json:"unread_enabled"`             // 是否计入未读数，默认true
	PushConfig       *PushConfig  `json:"push_config,omitempty"`      // 推送配置
}

// CancelSignalRoomInviteRequestV2 取消邀请加入信令房间请求
type CancelSignalRoomInviteRequestV2 struct {
	ChannelId        string       `json:"channel_id"`                 // 信令房间id
	InviterAccountId string       `json:"inviter_account_id"`         // 操作者，必须是信令房间创建者或者房间里的人
	InviteeAccountId string       `json:"invitee_account_id"`         // 被邀请者账号
	RequestId        string       `json:"request_id"`                 // 邀请ID，长度上限128字符
	ServerExtension  string       `json:"server_extension,omitempty"` // 发送通知时的扩展字段，最长4096个字符
	OfflineEnabled   bool         `json:"offline_enabled"`            // 关闭通知是否存离线，默认false
	UnreadEnabled    bool         `json:"unread_enabled"`             // 是否计入未读数，默认true
	PushConfig       *PushConfig  `json:"push_config,omitempty"`      // 推送配置
	RouteConfig      *RouteConfig `json:"route_config,omitempty"`     // 抄送配置
}

// KickSignalRoomMemberRequestV2 将成员踢出信令房间请求
type KickSignalRoomMemberRequestV2 struct {
	ChannelId         string       `json:"channel_id"`                 // 信令房间id
	OperatorAccountId string       `json:"operator_account_id"`        // 操作者，必须是信令房间创建者
	TargetAccountId   string       `json:"target_account_id"`          // 被踢的目标成员
	ServerExtension   string       `json:"server_extension,omitempty"` // 发送通知时的扩展字段，最长4096个字符
	OfflineEnabled    bool         `json:"offline_enabled"`            // 关闭通知是否存离线，默认false
	RouteConfig       *RouteConfig `json:"route_config,omitempty"`     // 抄送配置
}
