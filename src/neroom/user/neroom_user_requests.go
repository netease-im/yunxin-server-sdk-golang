package user

// CreateNeroomAccountRequest 创建Neroom账号请求
type CreateNeroomAccountRequest struct {
	UserUuid  string `json:"user_uuid,omitempty"`
	UserToken string `json:"user_token,omitempty"`
	UserName  string `json:"user_name,omitempty"`
	Icon      string `json:"icon,omitempty"`
	ImToken   string `json:"im_token,omitempty"`
}
