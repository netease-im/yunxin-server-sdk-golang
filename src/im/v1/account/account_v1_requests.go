package account

// CreateAccountRequestV1 创建账户请求
type CreateAccountRequestV1 struct {
	Accid  string `json:"accid"`            // 账号ID (必填)
	Token  string `json:"token,omitempty"`  // Token (可选)
	Name   string `json:"name,omitempty"`   // 用户昵称 (可选)
	Icon   string `json:"icon,omitempty"`   // 用户头像URL (可选)
	Sign   string `json:"sign,omitempty"`   // 用户签名 (可选)
	Email  string `json:"email,omitempty"`  // 用户邮箱 (可选)
	Birth  string `json:"birth,omitempty"`  // 用户生日 (可选)
	Mobile string `json:"mobile,omitempty"` // 用户手机号 (可选)
	Gender int    `json:"gender,omitempty"` // 用户性别 (可选)
	Ex     string `json:"ex,omitempty"`     // 扩展字段 (可选)
}

// UpdateTokenRequestV1 更新Token请求
type UpdateTokenRequestV1 struct {
	Accid string `json:"accid"` // 账号ID (必填)
	Token string `json:"token"` // Token (必填)
}

// RefreshTokenRequestV1 刷新Token请求
type RefreshTokenRequestV1 struct {
	Accid string `json:"accid"` // 账号ID (必填)
}

// BlockAccountRequestV1 封禁账户请求
type BlockAccountRequestV1 struct {
	Accid               string `json:"accid"`                         // 账号ID (必填)
	NeedKick            bool   `json:"needkick,omitempty"`            // 是否踢出在线设备 (可选)
	KickNotifyExt       string `json:"kickNotifyExt,omitempty"`       // 踢出通知扩展字段 (可选)
	NeedUnbindPushToken bool   `json:"needUnbindPushToken,omitempty"` // 是否解绑推送token (可选)
}

// UnBlockAccountRequestV1 解除账户封禁请求
type UnBlockAccountRequestV1 struct {
	Accid string `json:"accid"` // 账号ID (必填)
}

// MuteAccountRequestV1 账户禁言请求
type MuteAccountRequestV1 struct {
	Accid string `json:"accid"` // 账号ID (必填)
	Mute  bool   `json:"mute"`  // 是否禁言: true=禁言, false=取消禁言 (必填)
}

// MuteModuleRequestV1 模块禁言请求
type MuteModuleRequestV1 struct {
	Accid     string `json:"accid"`               // 账号ID (必填)
	MuteP2P   bool   `json:"muteP2P,omitempty"`   // 是否禁言单聊 (可选)
	MuteTeam  bool   `json:"muteTeam,omitempty"`  // 是否禁言群聊 (可选)
	MuteRoom  bool   `json:"muteRoom,omitempty"`  // 是否禁言聊天室 (可选)
	MuteQChat bool   `json:"muteQChat,omitempty"` // 是否禁言圈组 (可选)
}

// SetDonnopRequestV1 设置免打扰请求
type SetDonnopRequestV1 struct {
	Accid      string `json:"accid"`      // 账号ID (必填)
	DonnopOpen bool   `json:"donnopOpen"` // 是否开启免打扰 (必填)
}

// QueryAccountOnlineStatusRequestV1 查询账户在线状态请求
type QueryAccountOnlineStatusRequestV1 struct {
	Accids []string `json:"accids"` // 账号ID列表 (必填)
}

// QueryUserInfosRequestV1 获取用户信息请求
type QueryUserInfosRequestV1 struct {
	Accids     []string `json:"accids"`               // 账号ID列表 (必填)
	MuteStatus bool     `json:"muteStatus,omitempty"` // 是否返回禁言状态 (可选)
}

// UpdateUinfoRequestV1 更新用户信息请求
type UpdateUinfoRequestV1 struct {
	Accid  string `json:"accid"`            // 账号ID (必填)
	Name   string `json:"name,omitempty"`   // 用户昵称 (可选)
	Icon   string `json:"icon,omitempty"`   // 用户头像URL (可选)
	Sign   string `json:"sign,omitempty"`   // 用户签名 (可选)
	Email  string `json:"email,omitempty"`  // 用户邮箱 (可选)
	Birth  string `json:"birth,omitempty"`  // 用户生日 (可选)
	Mobile string `json:"mobile,omitempty"` // 用户手机号 (可选)
	Gender int    `json:"gender,omitempty"` // 用户性别 (可选)
	Ex     string `json:"ex,omitempty"`     // 扩展字段 (可选)
	Bid    string `json:"bid,omitempty"`    // 业务ID (可选)
}
