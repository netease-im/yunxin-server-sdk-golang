package team_member

// FailedOperation 失败的操作信息
type FailedOperation struct {
	AccountId string `json:"account_id"` // 账号ID
	ErrorCode int    `json:"error_code"` // 错误码
	ErrorMsg  string `json:"error_msg"`  // 错误信息
}

// TeamConfiguration 群组配置
type TeamConfiguration struct {
	JoinMode            int `json:"join_mode"`             // 加入群组权限
	AgreeMode           int `json:"agree_mode"`            // 邀请他人权限
	InviteMode          int `json:"invite_mode"`           // 邀请他人需要被邀请人同意模式
	UpdateTeamInfoMode  int `json:"update_team_info_mode"` // 群资料修改权限
	UpdateExtensionMode int `json:"update_extension_mode"` // 群扩展字段修改权限
}

// MemberInfo 成员信息
type MemberInfo struct {
	AccountId          string `json:"account_id"`           // 账号ID
	TeamMemberType     int    `json:"team_member_type"`     // 成员类型
	TeamNick           string `json:"team_nick"`            // 群昵称
	ChatBanned         bool   `json:"chat_banned"`          // 是否禁言
	MessageNotifyState int    `json:"message_notify_state"` // 消息通知状态
	JoinTime           int64  `json:"join_time"`            // 加入时间
	Extension          string `json:"extension"`            // 自定义扩展字段
	ServerExtension    string `json:"server_extension"`     // 服务端扩展字段
}

// JoinedTeamInfo 已加入的群组信息
type JoinedTeamInfo struct {
	TeamId         int64              `json:"team_id"`                 // 群组ID
	OwnerAccountId string             `json:"owner_account_id"`        // 群主账号ID
	Name           string             `json:"name"`                    // 群名称
	Icon           string             `json:"icon"`                    // 群头像URL
	MemberCount    int                `json:"member_count"`            // 群成员数量
	MembersLimit   int                `json:"members_limit"`           // 群成员人数上限
	CreateTime     int64              `json:"create_time"`             // 创建时间
	UpdateTime     int64              `json:"update_time"`             // 更新时间
	Configuration  *TeamConfiguration `json:"configuration,omitempty"` // 群配置
	MemberInfo     *MemberInfo        `json:"member_info,omitempty"`   // 成员信息
}

// InviteTeamMembersResponseV2 邀请成员加入群组响应
type InviteTeamMembersResponseV2 struct {
	SuccessList []string          `json:"success_list"` // 成功邀请的账号ID列表
	FailedList  []FailedOperation `json:"failed_list"`  // 邀请失败的列表
}

// KickTeamMembersResponseV2 踢出群成员响应
type KickTeamMembersResponseV2 struct {
	SuccessList []string          `json:"success_list"` // 成功踢出的账号ID列表
	FailedList  []FailedOperation `json:"failed_list"`  // 踢出失败的列表
}

// LeaveTeamResponseV2 退出群组响应
type LeaveTeamResponseV2 struct {
	// API返回的data为空对象 {}，无需任何字段
}

// UpdateTeamMemberResponseV2 更新群成员信息响应
type UpdateTeamMemberResponseV2 struct {
	MemberInfo MemberInfo `json:"member_info"` // 更新后的成员信息
}

// BatchMuteTeamMembersResponseV2 批量禁言群成员响应
type BatchMuteTeamMembersResponseV2 struct {
	SuccessList []string          `json:"success_list"` // 成功禁言的账号ID列表
	FailedList  []FailedOperation `json:"failed_list"`  // 禁言失败的列表
}

// QueryJoinedTeamsResponseV2 查询用户已加入的群组响应
type QueryJoinedTeamsResponseV2 struct {
	HasMore      bool             `json:"has_more"`       // 是否有更多数据
	NextToken    string           `json:"next_token"`     // 下一页token
	TeamInfoList []JoinedTeamInfo `json:"team_info_list"` // 群组信息列表
}
