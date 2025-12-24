package endpoint

import (
	"net/http"
)

// EndpointSelector 端点选择器接口
type EndpointSelector interface {
	// Init 初始化端点选择器
	Init(httpClient *http.Client)

	// Update 更新端点状态
	Update(endpoint string, result RequestResult)

	// SelectEndpoint 选择端点,excludeEndpoints 为需要排除的端点集合
	SelectEndpoint(excludeEndpoints map[string]bool) string

	// Shutdown 关闭选择器(可选实现)
	Shutdown()
}
