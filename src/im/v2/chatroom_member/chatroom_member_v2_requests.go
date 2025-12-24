package chatroom_member

// SetMemberRoleRequestV2 设置成员角色请求
type SetMemberRoleRequestV2 struct {
	AccountId             string                 `json:"-"`                                // 账号ID，用于构造URL
	RoomId                int64                  `json:"room_id"`                          // 聊天室ID
	OperatorAccountId     string                 `json:"operator_account_id"`              // 操作者账号ID
	MemberRole            int                    `json:"member_role"`                      // 成员角色
	RoomNick              string                 `json:"room_nick,omitempty"`              // 聊天室昵称
	RoomAvatar            string                 `json:"room_avatar,omitempty"`            // 聊天室头像
	MemberLevel           int                    `json:"member_level,omitempty"`           // 成员等级
	Extension             string                 `json:"extension,omitempty"`              // 扩展字段
	NotificationEnabled   bool                   `json:"notification_enabled,omitempty"`   // 是否启用通知
	NotificationExtension string                 `json:"notification_extension,omitempty"` // 通知扩展字段
	AntispamConfiguration *AntispamConfiguration `json:"antispam_configuration,omitempty"` // 反垃圾配置
}

// UpdateOnlineMemberInfoRequestV2 更新在线成员信息请求
type UpdateOnlineMemberInfoRequestV2 struct {
	AccountId             string                 `json:"-"`                                // 账号ID，用于构造URL
	RoomId                int64                  `json:"room_id"`                          // 聊天室ID
	Persistence           bool                   `json:"persistence,omitempty"`            // 是否持久化
	RoomNick              string                 `json:"room_nick,omitempty"`              // 聊天室昵称
	RoomAvatar            string                 `json:"room_avatar,omitempty"`            // 聊天室头像
	Extension             string                 `json:"extension,omitempty"`              // 扩展字段
	NotificationEnabled   bool                   `json:"notification_enabled,omitempty"`   // 是否启用通知
	NotificationExtension string                 `json:"notification_extension,omitempty"` // 通知扩展字段
	AntispamConfiguration *AntispamConfiguration `json:"antispam_configuration,omitempty"` // 反垃圾配置
}

// ToggleChatBanRequestV2 切换禁言请求
type ToggleChatBanRequestV2 struct {
	AccountId             string `json:"-"`                                // 账号ID，用于构造URL
	RoomId                int64  `json:"room_id"`                          // 聊天室ID
	OperatorAccountId     string `json:"operator_account_id"`              // 操作者账号ID
	ChatBanned            bool   `json:"chat_banned"`                      // 是否禁言
	NotificationEnabled   bool   `json:"notification_enabled,omitempty"`   // 是否启用通知
	NotificationExtension string `json:"notification_extension,omitempty"` // 通知扩展字段
}

// ToggleTempChatBanRequestV2 切换临时禁言请求
type ToggleTempChatBanRequestV2 struct {
	AccountId             string `json:"-"`                                // 账号ID，用于构造URL
	RoomId                int64  `json:"room_id"`                          // 聊天室ID
	OperatorAccountId     string `json:"operator_account_id"`              // 操作者账号ID
	ChatBanned            bool   `json:"chat_banned"`                      // 是否禁言
	ChatBannedDuration    int64  `json:"chat_banned_duration,omitempty"`   // 禁言时长（毫秒）
	NotificationEnabled   bool   `json:"notification_enabled,omitempty"`   // 是否启用通知
	NotificationExtension string `json:"notification_extension,omitempty"` // 通知扩展字段
}

// ToggleBlockedRequestV2 切换屏蔽请求
type ToggleBlockedRequestV2 struct {
	AccountId             string `json:"-"`                                // 账号ID，用于构造URL
	RoomId                int64  `json:"room_id"`                          // 聊天室ID
	OperatorAccountId     string `json:"operator_account_id"`              // 操作者账号ID
	Blocked               bool   `json:"blocked"`                          // 是否屏蔽
	NotificationEnabled   bool   `json:"notification_enabled,omitempty"`   // 是否启用通知
	NotificationExtension string `json:"notification_extension,omitempty"` // 通知扩展字段
}

// ModifyMemberTagsRequestV2 修改成员标签请求
type ModifyMemberTagsRequestV2 struct {
	AccountId        string   `json:"-"`                            // 账号ID，用于构造URL
	RoomId           int64    `json:"room_id"`                      // 聊天室ID
	Tags             []string `json:"tags,omitempty"`               // 标签列表
	NotifyTargetTags string   `json:"notify_target_tags,omitempty"` // 通知目标标签
	NotifyExtension  string   `json:"notify_extension,omitempty"`   // 通知扩展字段
}

// QueryTaggedMembersCountRequestV2 查询标签成员数量请求
type QueryTaggedMembersCountRequestV2 struct {
	RoomId int64  `json:"-"` // 聊天室ID，用于构造URL
	Tag    string `json:"-"` // 标签，用于构造查询参数
}

// ListTaggedMembersRequestV2 列出标签成员请求
type ListTaggedMembersRequestV2 struct {
	RoomId int64  `json:"-"` // 聊天室ID，用于构造URL
	Tag    string `json:"-"` // 标签，用于构造查询参数
	Offset int64  `json:"-"` // 偏移量，用于构造查询参数
	Limit  int    `json:"-"` // 限制数量，用于构造查询参数
}

// QueryChatroomBlacklistRequestV2 查询聊天室黑名单请求
type QueryChatroomBlacklistRequestV2 struct {
	RoomId int64 `json:"-"` // 聊天室ID，用于构造URL
}

// ToggleTaggedMembersChatBanRequestV2 切换标签成员禁言请求
type ToggleTaggedMembersChatBanRequestV2 struct {
	RoomId                 int64  `json:"room_id"`                            // 聊天室ID
	OperatorAccountId      string `json:"operator_account_id"`                // 操作者账号ID
	TargetTag              string `json:"target_tag"`                         // 目标标签
	ChatBanned             bool   `json:"chat_banned"`                        // 是否禁言
	ChatBannedDuration     int64  `json:"chat_banned_duration,omitempty"`     // 禁言时长（毫秒）
	NotificationEnabled    bool   `json:"notification_enabled,omitempty"`     // 是否启用通知
	NotificationExtension  string `json:"notification_extension,omitempty"`   // 通知扩展字段
	NotificationTargetTags string `json:"notification_target_tags,omitempty"` // 通知目标标签
}

// BatchQueryChatroomMembersRequestV2 批量查询聊天室成员请求
type BatchQueryChatroomMembersRequestV2 struct {
	RoomId     int64    `json:"-"` // 聊天室ID，用于构造URL
	AccountIds []string `json:"-"` // 账号ID列表，用于构造查询参数
}

// AddVirtualMembersRequestV2 添加虚拟成员请求
type AddVirtualMembersRequestV2 struct {
	RoomId                int64               `json:"room_id"`                          // 聊天室ID
	VirtualMembers        []VirtualMemberInfo `json:"virtual_members"`                  // 虚拟成员列表
	Extension             string              `json:"extension,omitempty"`              // 扩展字段
	NotificationEnabled   bool                `json:"notification_enabled,omitempty"`   // 是否启用通知
	NotificationExtension string              `json:"notification_extension,omitempty"` // 通知扩展字段
}

// DeleteVirtualMembersRequestV2 删除虚拟成员请求
type DeleteVirtualMembersRequestV2 struct {
	RoomId              int64    `json:"room_id"`                        // 聊天室ID
	AccountIds          []string `json:"account_ids"`                    // 账号ID列表
	NotificationEnabled bool     `json:"notification_enabled,omitempty"` // 是否启用通知
}

// ClearVirtualMembersRequestV2 清空虚拟成员请求
type ClearVirtualMembersRequestV2 struct {
	RoomId              int64 `json:"room_id"`                        // 聊天室ID
	NotificationEnabled bool  `json:"notification_enabled,omitempty"` // 是否启用通知
}

// QueryVirtualMembersRequestV2 查询虚拟成员请求
type QueryVirtualMembersRequestV2 struct {
	RoomId int64 `json:"-"` // 聊天室ID，用于构造查询参数
}

// QueryChatBannedRequestV2 查询禁言成员请求
type QueryChatBannedRequestV2 struct {
	RoomId int64 `json:"-"` // 聊天室ID，用于构造URL
}

// VirtualMemberInfo 虚拟成员信息
type VirtualMemberInfo struct {
	AccountId  string `json:"account_id"`            // 账号ID
	RoomNick   string `json:"room_nick,omitempty"`   // 聊天室昵称
	RoomAvatar string `json:"room_avatar,omitempty"` // 聊天室头像
	Extension  string `json:"extension,omitempty"`   // 扩展字段
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
