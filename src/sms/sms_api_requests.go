package sms

// SmsSendCodeRequest 发送短信验证码请求
type SmsSendCodeRequest struct {
	Mobile     string `json:"mobile"`               // 手机号码 (必填)
	Templateid int    `json:"templateid,omitempty"` // 模板ID (可选)
	CodeLen    int    `json:"codeLen,omitempty"`    // 验证码长度 (可选)
	AuthCode   string `json:"authCode,omitempty"`   // 自定义验证码 (可选)
	NeedUp     bool   `json:"needUp,omitempty"`     // 是否需要大写 (可选)
}
