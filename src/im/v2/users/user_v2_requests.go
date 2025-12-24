package users

// Gender constants
const (
	GenderUnknown = 0 // 未知
	GenderMale    = 1 // 男性
	GenderFemale  = 2 // 女性
)

// Antispam type constants
const (
	TypeText  = 1 // 文本类型
	TypeImage = 2 // 图片类型
)

// AntispamBusinessIdMap 反垃圾业务ID映射
type AntispamBusinessIdMap struct {
	Type       int    `json:"type"`        // 类型: 1=文本, 2=图片
	BusinessId string `json:"business_id"` // 业务ID
}

// AntispamConfiguration 反垃圾配置
type AntispamConfiguration struct {
	Enabled       bool                    `json:"enabled,omitempty"`         // 是否启用反垃圾
	BusinessIdMap []AntispamBusinessIdMap `json:"business_id_map,omitempty"` // 业务ID映射列表
}

// UpdateUserRequestV2 更新用户信息请求
type UpdateUserRequestV2 struct {
	AccountId             string                 `json:"-"`                                // 账号ID,路径参数 (必填)
	Name                  string                 `json:"name,omitempty"`                   // 用户昵称 (可选)
	Avatar                string                 `json:"avatar,omitempty"`                 // 用户头像URL (可选)
	Sign                  string                 `json:"sign,omitempty"`                   // 用户签名 (可选)
	Email                 string                 `json:"email,omitempty"`                  // 用户邮箱 (可选)
	Birthday              string                 `json:"birthday,omitempty"`               // 用户生日 (可选)
	Mobile                string                 `json:"mobile,omitempty"`                 // 用户手机号 (可选)
	Gender                int                    `json:"gender,omitempty"`                 // 用户性别: 0=未知, 1=男, 2=女 (可选)
	Extension             string                 `json:"extension,omitempty"`              // 自定义扩展字段 (可选)
	AntispamConfiguration *AntispamConfiguration `json:"antispam_configuration,omitempty"` // 反垃圾配置 (可选)
}

// GetUserRequestV2 获取用户信息请求
type GetUserRequestV2 struct {
	AccountId string `json:"-"` // 账号ID,路径参数 (必填)
}

// BatchGetUsersRequestV2 批量获取用户信息请求
type BatchGetUsersRequestV2 struct {
	AccountIds []string `json:"account_ids"` // 账号ID列表 (必填)
}

// GetUserOnlineStatusRequestV2 获取用户在线状态请求
type GetUserOnlineStatusRequestV2 struct {
	AccountIds []string `json:"account_ids"` // 账号ID列表 (必填)
}
