package chatroom

// CreateChatroomRequestV2 创建聊天室请求
type CreateChatroomRequestV2 struct {
	Creator               string                 `json:"creator"`                          // 聊天室创建者账号ID
	RoomName              string                 `json:"room_name"`                        // 聊天室名称（最大128字符）
	Announcement          string                 `json:"announcement,omitempty"`           // 聊天室公告（最大4096字符）
	LiveUrl               string                 `json:"live_url,omitempty"`               // 直播地址（最大1024字符）
	Extension             string                 `json:"extension,omitempty"`              // 自定义扩展字段（最大4096字符）
	QueueLevel            int                    `json:"queue_level,omitempty"`            // 队列管理权限等级
	DelayClosePolicy      int                    `json:"delay_close_policy,omitempty"`     // 聊天室自动关闭策略
	DelaySeconds          int64                  `json:"delay_seconds,omitempty"`          // 聊天室自动关闭时间（秒）
	InOutNotification     bool                   `json:"in_out_notification,omitempty"`    // 是否启用进出事件通知
	AntispamConfiguration *AntispamConfiguration `json:"antispam_configuration,omitempty"` // 反垃圾配置
}

// GetChatroomInfoRequestV2 获取聊天室信息请求
type GetChatroomInfoRequestV2 struct {
	RoomId                int64 `json:"-"` // 聊天室ID，用于构造URL
	NeedOnlineUserCount   bool  `json:"-"` // 是否需要在线用户数
	OnlineUserCountByType bool  `json:"-"` // 是否按类型统计在线用户数
}

// UpdateChatroomInfoRequestV2 更新聊天室信息请求
type UpdateChatroomInfoRequestV2 struct {
	RoomId                int64                  `json:"-"`                                // 聊天室ID，用于构造URL
	RoomName              string                 `json:"room_name,omitempty"`              // 聊天室名称
	Announcement          string                 `json:"announcement,omitempty"`           // 聊天室公告
	LiveUrl               string                 `json:"live_url,omitempty"`               // 直播地址
	Extension             string                 `json:"extension,omitempty"`              // 自定义扩展字段
	QueueLevel            int                    `json:"queue_level,omitempty"`            // 队列管理权限等级
	NotificationEnabled   bool                   `json:"notification_enabled,omitempty"`   // 是否启用通知
	NotificationExtension string                 `json:"notification_extension,omitempty"` // 通知扩展字段
	AntispamConfiguration *AntispamConfiguration `json:"antispam_configuration,omitempty"` // 反垃圾配置
}

// GetChatroomAddressRequestV2 获取聊天室地址请求
type GetChatroomAddressRequestV2 struct {
	RoomId     int64  `json:"-"` // 聊天室ID，用于构造URL
	AccountId  string `json:"-"` // 账号ID，用于构造查询参数
	ClientIp   string `json:"-"` // 客户端IP，用于构造查询参数
	ClientType int    `json:"-"` // 客户端类型，用于构造查询参数
	IpType     int    `json:"-"` // IP类型，用于构造查询参数
}

// UpdateChatroomStatusRequestV2 更新聊天室状态请求
type UpdateChatroomStatusRequestV2 struct {
	RoomId            int64  `json:"-"`                            // 聊天室ID，用于构造URL
	OperatorAccountId string `json:"operator_account_id"`          // 操作者账号ID
	Valid             bool   `json:"valid"`                        // 是否有效（true: 打开，false: 关闭）
	DelayClosePolicy  int    `json:"delay_close_policy,omitempty"` // 延迟关闭策略
	DelaySeconds      int64  `json:"delay_seconds,omitempty"`      // 延迟关闭时间（秒）
}

// ToggleChatroomMuteRequestV2 切换聊天室禁言请求
type ToggleChatroomMuteRequestV2 struct {
	RoomId                int64  `json:"-"`                                // 聊天室ID，用于构造URL
	OperatorAccountId     string `json:"operator_account_id"`              // 操作者账号ID
	ChatBanned            bool   `json:"chat_banned"`                      // 是否禁言
	NotificationEnabled   bool   `json:"notification_enabled,omitempty"`   // 是否启用通知
	NotificationExtension string `json:"notification_extension,omitempty"` // 通知扩展字段
}

// ToggleInOutNotificationRequestV2 切换进出通知请求
type ToggleInOutNotificationRequestV2 struct {
	RoomId            int64 `json:"room_id"`             // 聊天室ID
	InOutNotification bool  `json:"in_out_notification"` // 是否启用进出通知
}

// QueryOpenChatroomsRequestV2 查询开放聊天室请求
type QueryOpenChatroomsRequestV2 struct {
	CreatorAccountId string `json:"creator_account_id"` // 创建者账号ID
}

// ListOnlineMembersRequestV2 列出在线成员请求
type ListOnlineMembersRequestV2 struct {
	RoomId      int64  `json:"-"` // 聊天室ID，用于构造URL
	MemberRoles string `json:"-"` // 成员角色，用于构造查询参数
	Offset      int64  `json:"-"` // 偏移量，用于构造查询参数
	Limit       int    `json:"-"` // 限制数量，用于构造查询参数
}

// ListFixedMembersRequestV2 列出固定成员请求
type ListFixedMembersRequestV2 struct {
	RoomId      int64  `json:"-"` // 聊天室ID，用于构造URL
	MemberRoles string `json:"-"` // 成员角色，用于构造查询参数
}

// AntispamConfiguration 反垃圾配置
type AntispamConfiguration struct {
	Enabled       bool            `json:"enabled,omitempty"`         // 是否启用反垃圾
	BusinessIdMap []BusinessIdMap `json:"business_id_map,omitempty"` // 反垃圾业务ID映射列表
}

// BusinessIdMap 反垃圾业务ID映射
type BusinessIdMap struct {
	Type       int    `json:"type,omitempty"`        // 检测类型（1: 文本，2: 图片）
	BusinessId string `json:"business_id,omitempty"` // 反垃圾业务ID
}
