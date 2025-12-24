package user

// CreateNeroomAccountRequest 创建Neroom账号请求
type CreateNeroomAccountRequest struct {
	UserUuid   string `json:"userUuid"`             // 用户UUID，全局唯一
	Username   string `json:"username"`             // 用户名
	Avatar     string `json:"avatar,omitempty"`     // 用户头像URL
	Role       int    `json:"role,omitempty"`       // 用户角色，可选
	ExpireTime int64  `json:"expireTime,omitempty"` // 账号过期时间戳，可选
}
