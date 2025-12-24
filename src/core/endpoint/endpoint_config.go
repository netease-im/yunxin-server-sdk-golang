package endpoint

import "github.com/netease-im/yunxin-server-sdk-golang/src/core/base"

// EndpointConfig 端点配置
type EndpointConfig struct {
	RetryPolicy      base.RetryPolicy `json:"retry_policy"`
	EndpointSelector EndpointSelector `json:"endpoint_selector"`
}

// NewDefaultEndpointConfig 创建默认端点配置
func NewDefaultEndpointConfig() *EndpointConfig {
	return &EndpointConfig{}
}
