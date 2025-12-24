package chatroom

// CreateChatroomResponseV2 创建聊天室响应
type CreateChatroomResponseV2 struct {
	Valid        bool       `json:"valid,omitempty"`        // 聊天室状态（true: 打开，false: 关闭）
	Creator      string     `json:"creator,omitempty"`      // 聊天室创建者账号ID
	RoomName     string     `json:"room_name,omitempty"`    // 聊天室名称
	Announcement string     `json:"announcement,omitempty"` // 聊天室公告
	LiveUrl      string     `json:"live_url,omitempty"`     // 直播地址
	Extension    string     `json:"extension,omitempty"`    // 自定义扩展字段
	ChatBanned   bool       `json:"chat_banned,omitempty"`  // 是否全员禁言
	RoomId       int64      `json:"room_id,omitempty"`      // 聊天室ID
	QueueLevel   int        `json:"queue_level,omitempty"`  // 队列管理权限等级
	DelayInfo    *DelayInfo `json:"delay_info,omitempty"`   // 自动关闭信息
}

// GetChatroomInfoResponseV2 获取聊天室信息响应
type GetChatroomInfoResponseV2 struct {
	Valid                  bool       `json:"valid,omitempty"`                     // 聊天室状态
	Creator                string     `json:"creator,omitempty"`                   // 创建者
	RoomName               string     `json:"room_name,omitempty"`                 // 聊天室名称
	Announcement           string     `json:"announcement,omitempty"`              // 聊天室公告
	LiveUrl                string     `json:"live_url,omitempty"`                  // 直播地址
	Extension              string     `json:"extension,omitempty"`                 // 扩展字段
	ChatBanned             bool       `json:"chat_banned,omitempty"`               // 是否全员禁言
	RoomId                 int64      `json:"room_id,omitempty"`                   // 聊天室ID
	QueueLevel             int        `json:"queue_level,omitempty"`               // 队列管理权限等级
	InOutNotification      int        `json:"in_out_notification,omitempty"`       // 进出通知
	CdnMessageEnable       bool       `json:"cdn_message_enable,omitempty"`        // CDN消息启用
	OnlineUserCount        int64      `json:"online_user_count,omitempty"`         // 在线用户数
	IosOnlineUserCount     int64      `json:"ios_online_user_count,omitempty"`     // iOS在线用户数
	AndroidOnlineUserCount int64      `json:"android_online_user_count,omitempty"` // Android在线用户数
	PcOnlineUserCount      int64      `json:"pc_online_user_count,omitempty"`      // PC在线用户数
	WebOnlineUserCount     int64      `json:"web_online_user_count,omitempty"`     // Web在线用户数
	MacOnlineUserCount     int64      `json:"mac_online_user_count,omitempty"`     // Mac在线用户数
	HarmonyOnlineUserCount int64      `json:"harmony_online_user_count,omitempty"` // Harmony在线用户数
	DelayInfo              *DelayInfo `json:"delay_info,omitempty"`                // 延迟关闭信息
}

// UpdateChatroomInfoResponseV2 更新聊天室信息响应
type UpdateChatroomInfoResponseV2 struct {
	// Success response will only contain status code 200
}

// GetChatroomAddressResponseV2 获取聊天室地址响应
type GetChatroomAddressResponseV2 struct {
	Items []string `json:"items,omitempty"` // 地址列表
}

// UpdateChatroomStatusResponseV2 更新聊天室状态响应
type UpdateChatroomStatusResponseV2 struct {
	// Success response will only contain status code 200
}

// ToggleChatroomMuteResponseV2 切换聊天室禁言响应
type ToggleChatroomMuteResponseV2 struct {
	// Success response will only contain status code 200
}

// ToggleInOutNotificationResponseV2 切换进出通知响应
type ToggleInOutNotificationResponseV2 struct {
	// Success response will only contain status code 200
}

// QueryOpenChatroomsResponseV2 查询开放聊天室响应
type QueryOpenChatroomsResponseV2 struct {
	Items []string `json:"items,omitempty"` // 聊天室ID列表
}

// ListOnlineMembersResponseV2 列出在线成员响应
type ListOnlineMembersResponseV2 struct {
	HasMore bool                 `json:"has_more,omitempty"` // 是否有更多数据
	Offset  int64                `json:"offset,omitempty"`   // 偏移量
	Items   []ChatroomMemberInfo `json:"items,omitempty"`    // 成员信息列表
}

// ListFixedMembersResponseV2 列出固定成员响应
type ListFixedMembersResponseV2 struct {
	Items []ChatroomMemberInfo `json:"items,omitempty"` // 成员信息列表
}

// DelayInfo 延迟关闭信息
type DelayInfo struct {
	DelaySeconds     int64 `json:"delay_seconds,omitempty"`      // 延迟关闭时间（秒）
	DelayCloseEnable bool  `json:"delay_close_enable,omitempty"` // 是否启用延迟关闭
	StartTime        int64 `json:"start_time,omitempty"`         // 开始时间戳
	DelayClosePolicy int   `json:"delay_close_policy,omitempty"` // 延迟关闭策略
	Status           int   `json:"status,omitempty"`             // 当前状态
}

// ChatroomMemberInfo 聊天室成员信息
type ChatroomMemberInfo struct {
	AccountId          string       `json:"account_id,omitempty"`            // 账号ID
	RoomNick           string       `json:"room_nick,omitempty"`             // 聊天室昵称
	RoomAvatar         string       `json:"room_avatar,omitempty"`           // 聊天室头像
	Extension          string       `json:"extension,omitempty"`             // 扩展字段
	MemberRole         int          `json:"member_role,omitempty"`           // 成员角色
	MemberLevel        int          `json:"member_level,omitempty"`          // 成员等级
	IsOnline           bool         `json:"is_online,omitempty"`             // 是否在线
	EnterTime          int64        `json:"enter_time,omitempty"`            // 进入时间
	Blocked            bool         `json:"blocked,omitempty"`               // 是否被屏蔽
	ChatBanned         bool         `json:"chat_banned,omitempty"`           // 是否被禁言
	TempChatBanned     bool         `json:"temp_chat_banned,omitempty"`      // 是否临时禁言
	TempChatBannedTime int64        `json:"temp_chat_banned_time,omitempty"` // 临时禁言时间
	Tags               string       `json:"tags,omitempty"`                  // 标签
	NotifyTargetTags   string       `json:"notify_target_tags,omitempty"`    // 通知目标标签
	OnlineInfoList     []OnlineInfo `json:"online_info_list,omitempty"`      // 在线信息列表
	OnlineCount        int          `json:"online_count,omitempty"`          // 在线数量
}

// OnlineInfo 在线信息
type OnlineInfo struct {
	RoomNick   string `json:"room_nick,omitempty"`   // 聊天室昵称
	RoomAvatar string `json:"room_avatar,omitempty"` // 聊天室头像
	EnterTime  int64  `json:"enter_time,omitempty"`  // 进入时间
	ClientType int    `json:"client_type,omitempty"` // 客户端类型
	Robot      bool   `json:"robot,omitempty"`       // 是否机器人
}
