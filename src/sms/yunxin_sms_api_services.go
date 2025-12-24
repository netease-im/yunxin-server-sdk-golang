package sms

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
)

// YunxinSmsApiServices 云信短信API服务
type YunxinSmsApiServices struct {
	smsApiService SmsApiService
}

// NewYunxinSmsApiServices 创建云信短信API服务实例
func NewYunxinSmsApiServices(yunxinApiHttpClient core.YunxinApiHttpClient) *YunxinSmsApiServices {
	return &YunxinSmsApiServices{
		smsApiService: NewSmsApiService(yunxinApiHttpClient),
	}
}

// GetSmsApiService 获取短信API服务
func (s *YunxinSmsApiServices) GetSmsApiService() SmsApiService {
	return s.smsApiService
}
