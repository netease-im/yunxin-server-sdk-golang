package core

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/metrics"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/trace"
)

// HttpClient HTTP客户端接口
type HttpClient interface {
	// Execute 执行HTTP请求
	// method: HTTP方法
	// contextType: 内容类型
	// apiVersion: API版本
	// uri: URI，仅用于监控指标
	// path: 请求路径
	// queryParams: 查询字符串参数
	// data: 请求数据
	// 返回: HTTP响应和错误
	Execute(method http.HttpMethod, contextType http.ContextType, apiVersion trace.ApiVersion,
		uri string, path string, queryParams map[string]string, data string) (*http.HttpResponse, error)

	// Shutdown 关闭客户端
	Shutdown()

	// GetStats 获取统计信息
	GetStats() *metrics.Stats
}
