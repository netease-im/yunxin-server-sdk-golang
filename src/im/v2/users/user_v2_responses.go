package users

// Client type constants
const (
	ClientTypeAndroid = 1  // Android客户端
	ClientTypeIOS     = 2  // iOS客户端
	ClientTypePC      = 4  // PC客户端
	ClientTypeWeb     = 16 // Web客户端
	ClientTypeMAC     = 64 // Mac客户端
	ClientTypeHarmony = 65 // 鸿蒙客户端
)

// UserInfo 用户信息
type UserInfo struct {
	AccountId string `json:"account_id"` // 账号ID
	Name      string `json:"name"`       // 用户昵称
	Avatar    string `json:"avatar"`     // 用户头像URL
	Sign      string `json:"sign"`       // 用户签名
	Email     string `json:"email"`      // 用户邮箱
	Birthday  string `json:"birthday"`   // 用户生日
	Mobile    string `json:"mobile"`     // 用户手机号
	Gender    int    `json:"gender"`     // 用户性别
	Extension string `json:"extension"`  // 自定义扩展字段
}

// FailedInfo 失败信息
type FailedInfo struct {
	AccountId string `json:"account_id"` // 账号ID
	ErrorCode int    `json:"error_code"` // 错误码
	ErrorMsg  string `json:"error_msg"`  // 错误信息
}

// ClientStatus 客户端状态
type ClientStatus struct {
	ClientType int   `json:"client_type"` // 客户端类型
	LoginTime  int64 `json:"login_time"`  // 登录时间
}

// UserOnlineStatus 用户在线状态
type UserOnlineStatus struct {
	AccountId    string         `json:"account_id"`    // 账号ID
	OnlineStatus []ClientStatus `json:"online_status"` // 在线状态列表
}

// UpdateUserResponseV2 更新用户信息响应
type UpdateUserResponseV2 struct {
	UserInfo UserInfo `json:"user_info"` // 用户信息
}

// GetUserResponseV2 获取用户信息响应
type GetUserResponseV2 struct {
	UserInfo UserInfo `json:"user_info"` // 用户信息
}

// BatchGetUsersResponseV2 批量获取用户信息响应
type BatchGetUsersResponseV2 struct {
	SuccessList []UserInfo   `json:"success_list"` // 成功获取的用户信息列表
	FailedList  []FailedInfo `json:"failed_list"`  // 失败的列表
}

// GetUserOnlineStatusResponseV2 获取用户在线状态响应
type GetUserOnlineStatusResponseV2 struct {
	SuccessList []UserOnlineStatus `json:"success_list"` // 成功获取的在线状态列表
	FailedList  []FailedInfo       `json:"failed_list"`  // 失败的列表
}
