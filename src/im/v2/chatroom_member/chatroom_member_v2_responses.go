package chatroom_member

// SetMemberRoleResponseV2 设置成员角色响应
type SetMemberRoleResponseV2 struct {
	// Success response will only contain status code 200
}

// UpdateOnlineMemberInfoResponseV2 更新在线成员信息响应
type UpdateOnlineMemberInfoResponseV2 struct {
	Extension  string `json:"extension,omitempty"`   // 扩展字段
	RoomId     int64  `json:"room_id,omitempty"`     // 聊天室ID
	AccountId  string `json:"account_id,omitempty"`  // 账号ID
	RoomNick   string `json:"room_nick,omitempty"`   // 聊天室昵称
	RoomAvatar string `json:"room_avatar,omitempty"` // 聊天室头像
	MemberRole int    `json:"member_role,omitempty"` // 成员角色
	IsOnline   bool   `json:"is_online,omitempty"`   // 是否在线
}

// ToggleChatBanResponseV2 切换禁言响应
type ToggleChatBanResponseV2 struct {
	// Success response will only contain status code 200
}

// ToggleTempChatBanResponseV2 切换临时禁言响应
type ToggleTempChatBanResponseV2 struct {
	MuteDuration int64 `json:"mute_duration,omitempty"` // 禁言时长（毫秒）
}

// ToggleBlockedResponseV2 切换屏蔽响应
type ToggleBlockedResponseV2 struct {
	// Success response will only contain status code 200
}

// ModifyMemberTagsResponseV2 修改成员标签响应
type ModifyMemberTagsResponseV2 struct {
	// Success response will only contain status code 200
}

// QueryTaggedMembersCountResponseV2 查询标签成员数量响应
type QueryTaggedMembersCountResponseV2 struct {
	Tag             string `json:"tag,omitempty"`               // 标签
	OnlineUserCount int64  `json:"online_user_count,omitempty"` // 在线用户数
}

// ListTaggedMembersResponseV2 列出标签成员响应
type ListTaggedMembersResponseV2 struct {
	Offset  int64              `json:"offset,omitempty"`   // 偏移量
	HasMore bool               `json:"has_more,omitempty"` // 是否有更多数据
	Items   []TaggedMemberInfo `json:"items,omitempty"`    // 标签成员信息列表
}

// QueryChatroomBlacklistResponseV2 查询聊天室黑名单响应
type QueryChatroomBlacklistResponseV2 struct {
	Items []BlacklistMemberInfo `json:"items,omitempty"` // 黑名单成员信息列表
}

// ToggleTaggedMembersChatBanResponseV2 切换标签成员禁言响应
type ToggleTaggedMembersChatBanResponseV2 struct {
	MuteDuration int64 `json:"mute_duration,omitempty"` // 禁言时长（毫秒）
}

// BatchQueryChatroomMembersResponseV2 批量查询聊天室成员响应
type BatchQueryChatroomMembersResponseV2 struct {
	SuccessList []MemberInfo       `json:"success_list,omitempty"` // 成功列表
	FailedList  []FailedMemberInfo `json:"failed_list,omitempty"`  // 失败列表
}

// AddVirtualMembersResponseV2 添加虚拟成员响应
type AddVirtualMembersResponseV2 struct {
	SuccessList []string           `json:"success_list,omitempty"` // 成功的账号ID列表
	FailedList  []FailedMemberInfo `json:"failed_list,omitempty"`  // 失败列表
}

// DeleteVirtualMembersResponseV2 删除虚拟成员响应
type DeleteVirtualMembersResponseV2 struct {
	SuccessList []string           `json:"success_list,omitempty"` // 成功的账号ID列表
	FailedList  []FailedMemberInfo `json:"failed_list,omitempty"`  // 失败列表
}

// ClearVirtualMembersResponseV2 清空虚拟成员响应
type ClearVirtualMembersResponseV2 struct {
	Count int `json:"count,omitempty"` // 清除数量
}

// QueryVirtualMembersResponseV2 查询虚拟成员响应
type QueryVirtualMembersResponseV2 struct {
	Items []VirtualMemberDetailInfo `json:"items,omitempty"` // 虚拟成员详细信息列表
}

// QueryChatBannedResponseV2 查询禁言成员响应
type QueryChatBannedResponseV2 struct {
	Items []BannedMember `json:"items,omitempty"` // 禁言成员列表
}

// TaggedMemberInfo 标签成员信息
type TaggedMemberInfo struct {
	AccountId  string `json:"account_id,omitempty"`  // 账号ID
	RoomId     int64  `json:"room_id,omitempty"`     // 聊天室ID
	RoomNick   string `json:"room_nick,omitempty"`   // 聊天室昵称
	RoomAvatar string `json:"room_avatar,omitempty"` // 聊天室头像
	Extension  string `json:"extension,omitempty"`   // 扩展字段
	MemberRole int    `json:"member_role,omitempty"` // 成员角色
}

// BlacklistMemberInfo 黑名单成员信息
type BlacklistMemberInfo struct {
	AccountId  string `json:"account_id,omitempty"`  // 账号ID
	RoomNick   string `json:"room_nick,omitempty"`   // 聊天室昵称
	RoomAvatar string `json:"room_avatar,omitempty"` // 聊天室头像
	Extension  string `json:"extension,omitempty"`   // 扩展字段
	BlockTime  int64  `json:"block_time,omitempty"`  // 屏蔽时间
}

// MemberInfo 成员信息
type MemberInfo struct {
	AccountId   string `json:"account_id,omitempty"`   // 账号ID
	RoomNick    string `json:"room_nick,omitempty"`    // 聊天室昵称
	RoomAvatar  string `json:"room_avatar,omitempty"`  // 聊天室头像
	Extension   string `json:"extension,omitempty"`    // 扩展字段
	MemberRole  int    `json:"member_role,omitempty"`  // 成员角色
	MemberLevel int    `json:"member_level,omitempty"` // 成员等级
	EnterTime   int64  `json:"enter_time,omitempty"`   // 进入时间
}

// FailedMemberInfo 失败成员信息
type FailedMemberInfo struct {
	AccountId string `json:"account_id,omitempty"` // 账号ID
	ErrorCode int    `json:"error_code,omitempty"` // 错误码
	ErrorMsg  string `json:"error_msg,omitempty"`  // 错误信息
}

// VirtualMemberDetailInfo 虚拟成员详细信息
type VirtualMemberDetailInfo struct {
	AccountId  string `json:"account_id,omitempty"`  // 账号ID
	RoomNick   string `json:"room_nick,omitempty"`   // 聊天室昵称
	RoomAvatar string `json:"room_avatar,omitempty"` // 聊天室头像
	Extension  string `json:"extension,omitempty"`   // 扩展字段
	EnterTime  int64  `json:"enter_time,omitempty"`  // 进入时间
}

// BannedMember 禁言成员
type BannedMember struct {
	AccountId          string `json:"account_id,omitempty"`            // 账号ID
	RoomNick           string `json:"room_nick,omitempty"`             // 聊天室昵称
	RoomAvatar         string `json:"room_avatar,omitempty"`           // 聊天室头像
	ChatBanned         bool   `json:"chat_banned,omitempty"`           // 是否禁言
	TempChatBanned     bool   `json:"temp_chat_banned,omitempty"`      // 是否临时禁言
	TempChatBannedTime int64  `json:"temp_chat_banned_time,omitempty"` // 临时禁言时间
}
