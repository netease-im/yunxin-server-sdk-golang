package team

// Configuration 群配置
type Configuration struct {
	JoinMode            int `json:"join_mode,omitempty"`             // 加入群组权限
	AgreeMode           int `json:"agree_mode,omitempty"`            // 邀请他人权限
	InviteMode          int `json:"invite_mode,omitempty"`           // 邀请他人需要被邀请人同意模式
	UpdateTeamInfoMode  int `json:"update_team_info_mode,omitempty"` // 群资料修改权限
	UpdateExtensionMode int `json:"update_extension_mode,omitempty"` // 群扩展字段修改权限
}

// BusinessIdMap 反垃圾业务ID映射
type BusinessIdMap struct {
	Type       int    `json:"type"`        // 类型
	BusinessId string `json:"business_id"` // 业务ID
}

// AntispamConfiguration 反垃圾配置
type AntispamConfiguration struct {
	Enabled       bool            `json:"enabled,omitempty"`         // 是否启用反垃圾
	BusinessIdMap []BusinessIdMap `json:"business_id_map,omitempty"` // 业务ID映射列表
}

// CreateTeamRequestV2 创建群组请求
type CreateTeamRequestV2 struct {
	OwnerAccountId        string                 `json:"owner_account_id"`                 // 群主账号ID (必填)
	TeamType              int                    `json:"team_type"`                        // 群类型: 1=高级群, 2=超大群 (必填)
	Name                  string                 `json:"name"`                             // 群名称 (必填)
	Icon                  string                 `json:"icon,omitempty"`                   // 群头像URL (可选)
	Announcement          string                 `json:"announcement,omitempty"`           // 群公告 (可选)
	Intro                 string                 `json:"intro,omitempty"`                  // 群简介 (可选)
	MembersLimit          int                    `json:"members_limit,omitempty"`          // 群成员人数上限 (可选)
	ServerExtension       string                 `json:"server_extension,omitempty"`       // 服务端扩展字段 (可选)
	CustomerExtension     string                 `json:"customer_extension,omitempty"`     // 客户端扩展字段 (可选)
	InviteAccountIds      []string               `json:"invite_account_ids,omitempty"`     // 邀请的账号ID列表 (可选)
	InviteMsg             string                 `json:"invite_msg,omitempty"`             // 邀请附言 (可选)
	Extension             string                 `json:"extension,omitempty"`              // 自定义扩展字段 (可选)
	Configuration         *Configuration         `json:"configuration,omitempty"`          // 群配置 (可选)
	AntispamConfiguration *AntispamConfiguration `json:"antispam_configuration,omitempty"` // 反垃圾配置 (可选)
}

// UpdateTeamRequestV2 更新群组请求
type UpdateTeamRequestV2 struct {
	TeamId            string         `json:"-"`                            // 群组ID,路径参数 (必填)
	OperatorId        string         `json:"operator_id"`                  // 操作者账号ID (必填)
	TeamType          int            `json:"team_type"`                    // 群类型 (必填)
	Name              string         `json:"name,omitempty"`               // 群名称 (可选)
	Icon              string         `json:"icon,omitempty"`               // 群头像URL (可选)
	Announcement      string         `json:"announcement,omitempty"`       // 群公告 (可选)
	Intro             string         `json:"intro,omitempty"`              // 群简介 (可选)
	MembersLimit      int            `json:"members_limit,omitempty"`      // 群成员人数上限 (可选)
	ServerExtension   string         `json:"server_extension,omitempty"`   // 服务端扩展字段 (可选)
	CustomerExtension string         `json:"customer_extension,omitempty"` // 客户端扩展字段 (可选)
	Extension         string         `json:"extension,omitempty"`          // 自定义扩展字段 (可选)
	Configuration     *Configuration `json:"configuration,omitempty"`      // 群配置 (可选)
}

// DisbandTeamRequestV2 解散群组请求
type DisbandTeamRequestV2 struct {
	TeamId     string `json:"-"` // 群组ID,路径参数 (必填)
	TeamType   int    `json:"-"` // 群类型,查询参数 (必填)
	OperatorId string `json:"-"` // 操作者账号ID,查询参数 (必填)
}

// BatchQueryTeamInfoRequestV2 批量查询群组信息请求
type BatchQueryTeamInfoRequestV2 struct {
	TeamIds  []int64 `json:"team_ids"`  // 群组ID列表 (必填)
	TeamType int     `json:"team_type"` // 群类型 (必填)
}

// TransferTeamOwnerRequestV2 转让群主请求
type TransferTeamOwnerRequestV2 struct {
	TeamId            string `json:"-"`                    // 群组ID,路径参数 (必填)
	TeamType          int    `json:"team_type"`            // 群类型: 1=高级群, 2=超大群 (必填)
	NewOwnerAccountId string `json:"new_owner_account_id"` // 新群主账号ID (必填)
	Leave             int    `json:"leave"`                // 原群主是否退出: 1=退出, 2=不退出成为普通成员 (必填)
	Extension         string `json:"extension,omitempty"`  // 自定义扩展字段 (可选)
}

// AddTeamManagersRequestV2 添加群管理员请求
type AddTeamManagersRequestV2 struct {
	TeamId     string   `json:"-"`                   // 群组ID,路径参数 (必填)
	TeamType   int      `json:"team_type"`           // 群类型 (必填)
	OperatorId string   `json:"operator_id"`         // 操作者账号ID (必填)
	Managers   []string `json:"managers"`            // 要添加的管理员账号ID列表 (必填)
	Extension  string   `json:"extension,omitempty"` // 自定义扩展字段 (可选)
}

// RemoveTeamManagersRequestV2 移除群管理员请求
type RemoveTeamManagersRequestV2 struct {
	TeamId     string   `json:"-"`                   // 群组ID,路径参数 (必填)
	TeamType   int      `json:"team_type"`           // 群类型 (必填)
	OperatorId string   `json:"operator_id"`         // 操作者账号ID (必填)
	Managers   []string `json:"managers"`            // 要移除的管理员账号ID列表 (必填)
	Extension  string   `json:"extension,omitempty"` // 自定义扩展字段 (可选)
}

// GetTeamInfoRequestV2 获取群组信息请求
type GetTeamInfoRequestV2 struct {
	TeamId   string `json:"-"` // 群组ID,路径参数 (必填)
	TeamType int    `json:"-"` // 群类型,查询参数 (必填)
}

// ListTeamMembersRequestV2 查询群成员列表请求
type ListTeamMembersRequestV2 struct {
	TeamId    string `json:"-"` // 群组ID,路径参数 (必填)
	TeamType  int    `json:"-"` // 群类型,查询参数 (必填)
	Offset    int    `json:"-"` // 偏移量,查询参数 (可选)
	Limit     int    `json:"-"` // 限制数量,查询参数 (可选)
	NextToken string `json:"-"` // 下一页token,查询参数 (可选)
}

// ListOnlineTeamMembersRequestV2 查询群在线成员列表请求
type ListOnlineTeamMembersRequestV2 struct {
	TeamId string `json:"-"` // 群组ID,路径参数 (必填)
}

// BatchQueryTeamOnlineMembersCountRequestV2 批量查询群在线成员数量请求
type BatchQueryTeamOnlineMembersCountRequestV2 struct {
	TeamIds []int64 `json:"team_ids"` // 群组ID列表 (必填)
}
