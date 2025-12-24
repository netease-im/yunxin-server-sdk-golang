package team_member

// InviteTeamMembersRequestV2 邀请成员加入群组请求
type InviteTeamMembersRequestV2 struct {
	OperatorId       string   `json:"operator_id"`         // 操作者账号ID (必填)
	TeamId           int64    `json:"team_id"`             // 群组ID (必填)
	TeamType         int      `json:"team_type"`           // 群类型 (必填)
	InviteAccountIds []string `json:"invite_account_ids"`  // 要邀请的账号ID列表 (必填)
	Msg              string   `json:"msg,omitempty"`       // 邀请附言 (可选)
	Extension        string   `json:"extension,omitempty"` // 自定义扩展字段 (可选)
}

// KickTeamMembersRequestV2 踢出群成员请求
type KickTeamMembersRequestV2 struct {
	OperatorId     string   `json:"operator_id"`         // 操作者账号ID (必填)
	TeamId         int64    `json:"team_id"`             // 群组ID (必填)
	TeamType       int      `json:"team_type"`           // 群类型 (必填)
	KickAccountIds []string `json:"kick_account_ids"`    // 要踢出的账号ID列表 (必填)
	Extension      string   `json:"extension,omitempty"` // 自定义扩展字段 (可选)
}

// LeaveTeamRequestV2 退出群组请求
type LeaveTeamRequestV2 struct {
	AccountId string `json:"account_id"`          // 要退出的账号ID (必填)
	TeamId    int64  `json:"team_id"`             // 群组ID (必填)
	TeamType  int    `json:"team_type"`           // 群类型 (必填)
	Extension string `json:"extension,omitempty"` // 自定义扩展字段 (可选)
}

// UpdateTeamMemberRequestV2 更新群成员信息请求
type UpdateTeamMemberRequestV2 struct {
	AccountId          string `json:"-"`                              // 要更新的账号ID,路径参数 (必填)
	OperatorId         string `json:"operator_id"`                    // 操作者账号ID (必填)
	TeamId             int64  `json:"team_id"`                        // 群组ID (必填)
	TeamType           int    `json:"team_type"`                      // 群类型 (必填)
	TeamNick           string `json:"team_nick,omitempty"`            // 群昵称 (可选)
	ChatBanned         bool   `json:"chat_banned,omitempty"`          // 是否禁言 (可选)
	MessageNotifyState int    `json:"message_notify_state,omitempty"` // 消息通知状态 (可选)
	Extension          string `json:"extension,omitempty"`            // 自定义扩展字段 (可选)
	ServerExtension    string `json:"server_extension,omitempty"`     // 服务端扩展字段 (可选)
}

// BatchMuteTeamMembersRequestV2 批量禁言群成员请求
type BatchMuteTeamMembersRequestV2 struct {
	OperatorId        string   `json:"operator_id"`          // 操作者账号ID (必填)
	TeamId            int64    `json:"team_id"`              // 群组ID (必填)
	TeamType          int      `json:"team_type"`            // 群类型 (必填)
	ChatBanAccountIds []string `json:"chat_ban_account_ids"` // 要禁言的账号ID列表 (必填)
	ChatBanned        bool     `json:"chat_banned"`          // 是否禁言: true=禁言, false=取消禁言 (必填)
}

// QueryJoinedTeamsRequestV2 查询用户已加入的群组请求
type QueryJoinedTeamsRequestV2 struct {
	AccountId            string `json:"-"` // 账号ID,路径参数 (必填)
	TeamType             int    `json:"-"` // 群类型,查询参数 (必填)
	PageToken            string `json:"-"` // 分页token,查询参数 (可选)
	Limit                int    `json:"-"` // 每页数量,查询参数 (可选)
	NeedReturnMemberInfo int    `json:"-"` // 是否返回成员信息: 1=返回, 0=不返回,查询参数 (可选)
}
