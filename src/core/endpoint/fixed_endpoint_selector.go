package endpoint

import (
	"net/http"
)

// FixedEndpointSelector 固定端点选择器
type FixedEndpointSelector struct {
	endpoint string
}

// NewFixedEndpointSelector 创建一个新的固定端点选择器实例
func NewFixedEndpointSelector(endpoint string) *FixedEndpointSelector {
	return &FixedEndpointSelector{
		endpoint: endpoint,
	}
}

// Init 初始化端点选择器
func (f *FixedEndpointSelector) Init(httpClient *http.Client) {
	// 固定端点选择器不需要初始化
}

// Update 更新端点请求结果
func (f *FixedEndpointSelector) Update(endpoint string, result RequestResult) {
	// 固定端点选择器不需要更新
}

// SelectEndpoint 选择端点
func (f *FixedEndpointSelector) SelectEndpoint(excludeEndpoints map[string]bool) string {
	return f.endpoint
}

func (d *FixedEndpointSelector) Shutdown() {
	// 固定端点选择器不需要shutdown
}
