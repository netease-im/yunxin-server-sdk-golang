package user

// CreateNeroomAccountResponse 创建Neroom账号响应
type CreateNeroomAccountResponse struct {
	UserUuid  string `json:"user_uuid"`
	UserToken string `json:"user_token"`
	UserName  string `json:"user_name"`
	Icon      string `json:"icon"`
	ImToken   string `json:"im_token"`
}
