package account

// CreateAccountResponseV1 创建账户响应
type CreateAccountResponseV1 struct {
	Accid string `json:"accid"` // 账号ID
	Name  string `json:"name"`  // 用户昵称
	Token string `json:"token"` // Token
}

// UpdateTokenResponseV1 更新Token响应
type UpdateTokenResponseV1 struct {
	// API返回的data为空对象 {}，无需任何字段
}

// RefreshTokenResponseV1 刷新Token响应
type RefreshTokenResponseV1 struct {
	Accid string `json:"accid"` // 账号ID
	Token string `json:"token"` // 新Token
}

// BlockAccountResponseV1 封禁账户响应
type BlockAccountResponseV1 struct {
	// API返回的data为空对象 {}，无需任何字段
}

// UnBlockAccountResponseV1 解除账户封禁响应
type UnBlockAccountResponseV1 struct {
	// API返回的data为空对象 {}，无需任何字段
}

// MuteAccountResponseV1 账户禁言响应
type MuteAccountResponseV1 struct {
	// API返回的data为空对象 {}，无需任何字段
}

// MuteModuleResponseV1 模块禁言响应
type MuteModuleResponseV1 struct {
	MuteP2P   bool `json:"muteP2P"`   // 是否禁言单聊
	MuteTeam  bool `json:"muteTeam"`  // 是否禁言群聊
	MuteRoom  bool `json:"muteRoom"`  // 是否禁言聊天室
	MuteQChat bool `json:"muteQChat"` // 是否禁言圈组
}

// SetDonnopResponseV1 设置免打扰响应
type SetDonnopResponseV1 struct {
	// API返回的data为空对象 {}，无需任何字段
}

// Status 在线状态
type Status struct {
	LoginTime  int64 `json:"loginTime"`  // 登录时间
	ClientType int   `json:"clientType"` // 客户端类型
}

// OnlineStatus 账号在线状态
type OnlineStatus struct {
	Accid      string   `json:"accid"`      // 账号ID
	StatusList []Status `json:"statusList"` // 状态列表
}

// QueryAccountOnlineStatusResponseV1 查询账户在线状态响应
type QueryAccountOnlineStatusResponseV1 struct {
	InvalidAccids    []string       `json:"invalidAccids"`    // 无效的账号ID列表
	OnlineStatusList []OnlineStatus `json:"onlineStatusList"` // 在线状态列表
}

// UserInfo 用户信息
type UserInfo struct {
	Accid     string `json:"accid"`     // 账号ID
	Name      string `json:"name"`      // 用户昵称
	Icon      string `json:"icon"`      // 用户头像URL
	Sign      string `json:"sign"`      // 用户签名
	Email     string `json:"email"`     // 用户邮箱
	Birth     string `json:"birth"`     // 用户生日
	Mobile    string `json:"mobile"`    // 用户手机号
	Ex        string `json:"ex"`        // 扩展字段
	Gender    int    `json:"gender"`    // 用户性别
	Valid     bool   `json:"valid"`     // 账号是否有效
	Mute      bool   `json:"mute"`      // 是否禁言
	MuteP2P   bool   `json:"muteP2P"`   // 是否禁言单聊
	MuteQChat bool   `json:"muteQChat"` // 是否禁言圈组
	MuteTeam  bool   `json:"muteTeam"`  // 是否禁言群聊
	MuteRoom  bool   `json:"muteRoom"`  // 是否禁言聊天室
}

// QueryUserInfosResponseV1 获取用户信息响应
type QueryUserInfosResponseV1 struct {
	UserInfoList []UserInfo `json:"userInfoList"` // 用户信息列表
}

// UpdateUinfoResponseV1 更新用户信息响应
type UpdateUinfoResponseV1 struct {
	// API返回的data为空对象 {}，无需任何字段
}
