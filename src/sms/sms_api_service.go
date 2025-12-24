package sms

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
)

// SmsApiService 短信服务接口
type SmsApiService interface {
	// SendCode 发送短信验证码
	SendCode(request *SmsSendCodeRequest) (*core.Result[*SmsSendCodeResponse], error)
}

// SmsApiServiceImpl 短信服务实现
type SmsApiServiceImpl struct {
	httpClient core.YunxinApiHttpClient
}

// NewSmsApiService 创建短信服务实例
func NewSmsApiService(httpClient core.YunxinApiHttpClient) SmsApiService {
	return &SmsApiServiceImpl{
		httpClient: httpClient,
	}
}
