package account

// BatchQueryAccountsResponseV2 批量查询账号响应
type BatchQueryAccountsResponseV2 struct {
	SuccessList []AccountInfo   `json:"success_list,omitempty"` // 成功列表
	FailedList  []FailedAccount `json:"failed_list,omitempty"`  // 失败列表
}

// Account 账号信息
type AccountInfo struct {
	AccountId     string                    `json:"account_id,omitempty"`    // 账号ID
	Configuration *AccountConfigurationResp `json:"configuration,omitempty"` // 账号配置
}

// ResponseConfiguration 响应中的配置信息
type ResponseConfiguration struct {
	Enabled                      bool `json:"enabled,omitempty"`                          // 是否启用
	P2pChatBanned                bool `json:"p2p_chat_banned,omitempty"`                  // 是否禁止点对点聊天
	TeamChatBanned               bool `json:"team_chat_banned,omitempty"`                 // 是否禁止群聊
	ChatroomChatBanned           bool `json:"chatroom_chat_banned,omitempty"`             // 是否禁止聊天室聊天
	QchatChatBanned              bool `json:"qchat_chat_banned,omitempty"`                // 是否禁止圈组聊天
	PushEnabledWhenDesktopOnline bool `json:"push_enabled_when_desktop_online,omitempty"` // 桌面端在线时是否推送
}

// FailedAccount 失败的账号
type FailedAccount struct {
	AccountID string `json:"account_id,omitempty"`
	ErrorCode int    `json:"error_code,omitempty"`
	ErrorMsg  string `json:"error_msg,omitempty"`
}

// CreateAccountResponseV2 创建账号响应
type CreateAccountResponseV2 struct {
	AccountId       string                    `json:"account_id,omitempty"`       // 账号ID
	Token           string                    `json:"token,omitempty"`            // Token
	Configuration   *AccountConfigurationResp `json:"configuration,omitempty"`    // 账号配置
	UserInformation *UserInformationResp      `json:"user_information,omitempty"` // 用户信息
}

// DisableAccountResponseV2 禁用/启用账号响应
type DisableAccountResponseV2 struct {
	AccountId     string                    `json:"account_id,omitempty"`    // 账号ID
	Enabled       bool                      `json:"enabled,omitempty"`       // 是否启用
	Configuration *AccountConfigurationResp `json:"configuration,omitempty"` // 账号配置
}

// GetAccountDetailsResponseV2 获取账号详情响应
type GetAccountDetailsResponseV2 struct {
	AccountId     string                    `json:"account_id,omitempty"`    // 账号ID
	Token         string                    `json:"token,omitempty"`         // Token
	Configuration *AccountConfigurationResp `json:"configuration,omitempty"` // 账号配置
}

// KickAccountResponseV2 踢账号下线响应
type KickAccountResponseV2 struct {
	SuccessList []KickedDevice `json:"success_list,omitempty"` // 成功列表
	FailedList  []FailedKick   `json:"failed_list,omitempty"`  // 失败列表
}

// RefreshTokenResponseV2 刷新token响应
type RefreshTokenResponseV2 struct {
	AccountID string `json:"account_id,omitempty"`
	Token     string `json:"token,omitempty"`
}

// SetPushConfigResponseV2 设置推送配置响应
type SetPushConfigResponseV2 struct {
	AccountID                    string `json:"account_id,omitempty"`
	PushEnabledWhenDesktopOnline bool   `json:"push_enabled_when_desktop_online,omitempty"`
}

// UpdateAccountResponseV2 更新账号响应
type UpdateAccountResponseV2 struct {
	AccountID     string                    `json:"account_id,omitempty"`    // 账号ID
	Token         string                    `json:"token,omitempty"`         // Token
	Configuration *AccountConfigurationResp `json:"configuration,omitempty"` // 账号配置
}

// AccountConfigurationResp 账号配置响应
type AccountConfigurationResp struct {
	Enabled                      bool `json:"enabled,omitempty"`                          // 是否启用
	P2pChatBanned                bool `json:"p2p_chat_banned,omitempty"`                  // 是否禁止点对点聊天
	TeamChatBanned               bool `json:"team_chat_banned,omitempty"`                 // 是否禁止群聊
	ChatroomChatBanned           bool `json:"chatroom_chat_banned,omitempty"`             // 是否禁止聊天室聊天
	QchatChatBanned              bool `json:"qchat_chat_banned,omitempty"`                // 是否禁止圈组聊天
	PushEnabledWhenDesktopOnline bool `json:"push_enabled_when_desktop_online,omitempty"` // 桌面端在线时是否推送
}

// UserInformationResp 用户信息响应
type UserInformationResp struct {
	Name      string `json:"name,omitempty"`      // 用户名
	Avatar    string `json:"avatar,omitempty"`    // 头像
	Sign      string `json:"sign,omitempty"`      // 签名
	Email     string `json:"email,omitempty"`     // 邮箱
	Birthday  string `json:"birthday,omitempty"`  // 生日
	Mobile    string `json:"mobile,omitempty"`    // 手机号
	Gender    int    `json:"gender,omitempty"`    // 性别
	Extension string `json:"extension,omitempty"` // 扩展字段
}

// KickedDevice 被踢下线的设备信息
type KickedDevice struct {
	DeviceId   string `json:"device_id,omitempty"`   // 设备ID
	ClientType int    `json:"client_type,omitempty"` // 客户端类型
	Consid     string `json:"consid,omitempty"`      // 连接ID
	LoginTime  int64  `json:"login_time,omitempty"`  // 登录时间
}

// FailedKick 踢下线失败信息
type FailedKick struct {
	DeviceId  string `json:"device_id,omitempty"`  // 设备ID
	ErrorCode int    `json:"error_code,omitempty"` // 错误码
	ErrorMsg  string `json:"error_msg,omitempty"`  // 错误信息
}
