package account

// CreateAccountRequestV2 创建账号请求
type CreateAccountRequestV2 struct {
	AccountId             string                 `json:"account_id,omitempty"`             // 账号ID
	Token                 string                 `json:"token,omitempty"`                  // Token
	Configuration         *AccountConfiguration  `json:"configuration,omitempty"`          // 账号配置
	UserInformation       *UserInformation       `json:"user_information,omitempty"`       // 用户信息
	AntispamConfiguration *AntispamConfiguration `json:"antispam_configuration,omitempty"` // 反垃圾配置
}

// UpdateAccountRequestV2 更新账号请求
type UpdateAccountRequestV2 struct {
	AccountId           string                `json:"-"`                               // 账号ID，不包含在JSON中，用于构造URL
	Token               string                `json:"token,omitempty"`                 // Token
	Configuration       *AccountConfiguration `json:"configuration,omitempty"`         // 账号配置
	NeedKick            bool                  `json:"need_kick,omitempty"`             // 是否需要踢下线
	KickNotifyExtension string                `json:"kick_notify_extension,omitempty"` // 踢下线通知扩展字段
}

// BatchQueryAccountsRequestV2 批量查询账号请求
type BatchQueryAccountsRequestV2 struct {
	AccountList []string `json:"-"` // 账号列表，用于构造查询参数
}

// DisableAccountRequestV2 禁用/启用账号请求
type DisableAccountRequestV2 struct {
	AccountId           string `json:"-"`                               // 账号ID，用于构造URL
	Enabled             bool   `json:"enabled"`                         // 是否启用
	NeedKick            bool   `json:"need_kick,omitempty"`             // 是否需要踢下线
	KickNotifyExtension string `json:"kick_notify_extension,omitempty"` // 踢下线通知扩展字段
}

// SetPushConfigRequestV2 设置推送配置请求
type SetPushConfigRequestV2 struct {
	AccountId                    string `json:"-"`                                // 账号ID，用于构造URL
	PushEnabledWhenDesktopOnline bool   `json:"push_enabled_when_desktop_online"` // 桌面端在线时是否推送
}

// GetAccountDetailsRequestV2 获取账号详情请求
type GetAccountDetailsRequestV2 struct {
	AccountId string `json:"-"` // 账号ID，用于构造URL
}

// KickAccountRequestV2 踢账号下线请求
type KickAccountRequestV2 struct {
	AccountId           string   `json:"-"`                               // 账号ID，用于构造URL
	Type                int      `json:"type"`                            // 踢下线类型：1-踢所有设备，2-踢device_id_list中的设备，3-保留device_id_list中的设备，踢其他设备
	DeviceIdList        []string `json:"device_id_list,omitempty"`        // 设备ID列表（最多20个）
	KickNotifyExtension string   `json:"kick_notify_extension,omitempty"` // 踢下线通知扩展字段（最大256字符）
}

// RefreshTokenRequestV2 刷新Token请求
type RefreshTokenRequestV2 struct {
	AccountId string `json:"-"` // 账号ID，用于构造URL
}

// AccountConfiguration 账号配置
type AccountConfiguration struct {
	Enabled                      bool `json:"enabled,omitempty"`                          // 是否启用
	P2pChatBanned                bool `json:"p2p_chat_banned,omitempty"`                  // 是否禁止点对点聊天
	TeamChatBanned               bool `json:"team_chat_banned,omitempty"`                 // 是否禁止群聊
	ChatroomChatBanned           bool `json:"chatroom_chat_banned,omitempty"`             // 是否禁止聊天室聊天
	QchatChatBanned              bool `json:"qchat_chat_banned,omitempty"`                // 是否禁止圈组聊天
	PushEnabledWhenDesktopOnline bool `json:"push_enabled_when_desktop_online,omitempty"` // 桌面端在线时是否推送
}

// AntispamConfiguration 反垃圾配置
type AntispamConfiguration struct {
	Enabled       bool                `json:"enabled,omitempty"`         // 是否启用
	BusinessIdMap []BusinessIdMapItem `json:"business_id_map,omitempty"` // 业务ID映射列表
}

// BusinessIdMapItem 业务ID映射项
type BusinessIdMapItem struct {
	Type       int    `json:"type,omitempty"`        // 类型
	BusinessId string `json:"business_id,omitempty"` // 业务ID
}

// UserInformation 用户信息
type UserInformation struct {
	Name      string `json:"name,omitempty"`      // 用户名
	Avatar    string `json:"avatar,omitempty"`    // 头像
	Sign      string `json:"sign,omitempty"`      // 签名
	Email     string `json:"email,omitempty"`     // 邮箱
	Birthday  string `json:"birthday,omitempty"`  // 生日
	Mobile    string `json:"mobile,omitempty"`    // 手机号
	Gender    int    `json:"gender,omitempty"`    // 性别
	Extension string `json:"extension,omitempty"` // 扩展字段
}
