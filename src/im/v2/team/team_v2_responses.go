package team

// FailedAccount 失败的账号信息
type FailedAccount struct {
	AccountId string `json:"account_id"` // 账号ID
	ErrorCode int    `json:"error_code"` // 错误码
	ErrorMsg  string `json:"error_msg"`  // 错误信息
}

// TeamInfo 群组信息
type TeamInfo struct {
	TeamType          *int           `json:"team_type"`               // 群类型
	TeamId            int64          `json:"team_id"`                 // 群组ID
	OwnerAccountId    string         `json:"owner_account_id"`        // 群主账号ID
	MemberCount       *int           `json:"member_count"`            // 群成员数量
	Name              string         `json:"name"`                    // 群名称
	Icon              string         `json:"icon"`                    // 群头像URL
	Announcement      string         `json:"announcement"`            // 群公告
	Intro             string         `json:"intro"`                   // 群简介
	MembersLimit      int            `json:"members_limit"`           // 群成员人数上限
	ServerExtension   string         `json:"server_extension"`        // 服务端扩展字段
	CustomerExtension string         `json:"customer_extension"`      // 客户端扩展字段
	CreateTime        int64          `json:"create_time"`             // 创建时间
	UpdateTime        int64          `json:"update_time"`             // 更新时间
	Configuration     *Configuration `json:"configuration,omitempty"` // 群配置
}

// TeamMember 群成员信息
type TeamMember struct {
	TeamId           int64  `json:"team_id"`            // 群组ID
	AccountId        string `json:"account_id"`         // 账号ID
	MemberRole       *int   `json:"member_role"`        // 成员角色: 0=普通成员, 1=群主, 2=管理员
	TeamNick         string `json:"team_nick"`          // 成员在群内的昵称
	ServerExtension  string `json:"server_extension"`   // 服务端扩展字段
	ChatBanned       *bool  `json:"chat_banned"`        // 是否被禁言
	JoinTime         int64  `json:"join_time"`          // 加入时间
	UpdateTime       int64  `json:"update_time"`        // 更新时间
	InvitorAccountId string `json:"invitor_account_id"` // 邀请者账号ID
}

// OnlineMemberCount 在线成员数量信息
type OnlineMemberCount struct {
	TeamId            int64 `json:"team_id"`             // 群组ID
	OnlineMemberCount *int  `json:"online_member_count"` // 在线成员数量
}

// CreateTeamResponseV2 创建群组响应
type CreateTeamResponseV2 struct {
	FailedList []FailedAccount `json:"failed_list"` // 邀请失败的账号列表
	TeamInfo   TeamInfo        `json:"team_info"`   // 群组信息
}

// UpdateTeamResponseV2 更新群组响应
type UpdateTeamResponseV2 struct {
	TeamInfo TeamInfo `json:"team_info"` // 群组信息
}

// DisbandTeamResponseV2 解散群组响应
type DisbandTeamResponseV2 struct {
	// API返回的data为空对象 {}，无需任何字段
}

// BatchQueryTeamInfoResponseV2 批量查询群组信息响应
type BatchQueryTeamInfoResponseV2 struct {
	TeamInfoList []TeamInfo `json:"team_info_list"` // 群组信息列表
}

// TransferTeamOwnerResponseV2 转让群主响应
type TransferTeamOwnerResponseV2 struct {
	// API返回的data为空对象 {}，无需任何字段
}

// AddTeamManagersResponseV2 添加群管理员响应
type AddTeamManagersResponseV2 struct {
	// API返回的data为空对象 {}，无需任何字段
}

// RemoveTeamManagersResponseV2 移除群管理员响应
type RemoveTeamManagersResponseV2 struct {
	// API返回的data为空对象 {}，无需任何字段
}

// GetTeamInfoResponseV2 获取群组信息响应
type GetTeamInfoResponseV2 struct {
	TeamInfo TeamInfo `json:"team_info"` // 群组信息
}

// ListTeamMembersResponseV2 查询群成员列表响应
type ListTeamMembersResponseV2 struct {
	Items     []TeamMember `json:"items"`      // 群成员列表
	NextToken string       `json:"next_token"` // 下一页token
	HasMore   *bool        `json:"has_more"`   // 是否有更多数据
}

// LoginInfo 登录信息
type LoginInfo struct {
	LoginTime  int64 `json:"login_time"`  // 登录时间
	ClientType *int  `json:"client_type"` // 客户端类型
}

// OnlineTeamMember 在线群成员信息
type OnlineTeamMember struct {
	AccountId    string      `json:"account_id"`    // 账号ID
	OnlineStatus []LoginInfo `json:"online_status"` // 在线状态列表(每个设备一个登录信息)
}

// ListOnlineTeamMembersResponseV2 查询群在线成员列表响应
type ListOnlineTeamMembersResponseV2 struct {
	Items []OnlineTeamMember `json:"items"` // 在线成员列表
}

// SuccessTeam 成功查询的群组在线信息
type SuccessTeam struct {
	TeamId       int64 `json:"team_id"`       // 群组ID
	OnlineCount  *int  `json:"online_count"`  // 在线成员数量
	OfflineCount *int  `json:"offline_count"` // 离线成员数量
}

// FailedTeam 查询失败的群组信息
type FailedTeam struct {
	TeamId    int64  `json:"team_id"`    // 群组ID
	ErrorCode *int   `json:"error_code"` // 错误码
	ErrorMsg  string `json:"error_msg"`  // 错误信息
}

// BatchQueryTeamOnlineMembersCountResponseV2 批量查询群在线成员数量响应
type BatchQueryTeamOnlineMembersCountResponseV2 struct {
	SuccessList []SuccessTeam `json:"success_list"` // 成功查询的群组列表
	FailedList  []FailedTeam  `json:"failed_list"`  // 查询失败的群组列表
}
