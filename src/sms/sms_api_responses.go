package sms

// SmsSendCodeResponse 发送短信验证码响应
type SmsSendCodeResponse struct {
	Sendid   int64  `json:"sendid"`   // 发送ID
	AuthCode string `json:"authCode"` // 验证码
}
