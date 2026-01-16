package user

// CreateNeroomAccountRequest 创建Neroom账号请求
type CreateNeroomAccountRequest struct {
	UserUuid  string `json:"user_uuid"`
	UserToken string `json:"user_token"`
	UserName  string `json:"user_name"`
	Icon      string `json:"icon"`
	ImToken   string `json:"im_token"`
}
