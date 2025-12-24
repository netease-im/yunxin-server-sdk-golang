package endpoint

import "net/http"

// EndpointFetcher 端点获取器接口
type EndpointFetcher interface {
	// Init 初始化
	Init(httpClient *http.Client)

	// Get 获取端点
	Get() *Endpoints
}
