package endpoint

import "net/http"

// LocalEndpointFetcher 本地端点获取器 - 实现EndpointFetcher接口
type LocalEndpointFetcher struct {
	endpoints *Endpoints
}

// NewLocalEndpointFetcher 创建本地端点获取器（自定义端点）
func NewLocalEndpointFetcher(defaultEndpoint string, backupEndpoints ...string) *LocalEndpointFetcher {
	return &LocalEndpointFetcher{
		endpoints: &Endpoints{
			DefaultEndpoint: defaultEndpoint,
			BackupEndpoints: backupEndpoints,
		},
	}
}

func (f *LocalEndpointFetcher) Init(httpClient *http.Client) {
	// 固定端点选择器不需要初始化
}

func (f *LocalEndpointFetcher) Get() *Endpoints {
	return f.endpoints
}
