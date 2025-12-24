package metrics

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/trace"
)

// EndpointStats 端点统计信息
type EndpointStats struct {
	Endpoint string  `json:"endpoint"` // 端点地址
	Result   string  `json:"result"`   // 结果状态
	Count    int64   `json:"count"`    // 请求次数
	Avg      float64 `json:"avg"`      // 平均响应时间
	Max      float64 `json:"max"`      // 最大响应时间
	P50      float64 `json:"p50"`      // 50%分位数
	P75      float64 `json:"p75"`      // 75%分位数
	P90      float64 `json:"p90"`      // 90%分位数
	P95      float64 `json:"p95"`      // 95%分位数
	P99      float64 `json:"p99"`      // 99%分位数
	P999     float64 `json:"p999"`     // 99.9%分位数
}

// UriStats URI统计信息
type UriStats struct {
	Endpoint    string           `json:"endpoint"`    // 端点地址
	Method      http.HttpMethod  `json:"method"`      // HTTP方法
	ApiVersion  trace.ApiVersion `json:"apiVersion"`  // API版本
	ContextType http.ContextType `json:"contextType"` // 内容类型
	Uri         string           `json:"uri"`         // URI路径
	Result      string           `json:"result"`      // 结果状态
	Count       int64            `json:"count"`       // 请求次数
	Avg         float64          `json:"avg"`         // 平均响应时间
	Max         float64          `json:"max"`         // 最大响应时间
	P50         float64          `json:"p50"`         // 50%分位数
	P75         float64          `json:"p75"`         // 75%分位数
	P90         float64          `json:"p90"`         // 90%分位数
	P95         float64          `json:"p95"`         // 95%分位数
	P99         float64          `json:"p99"`         // 99%分位数
	P999        float64          `json:"p999"`        // 99.9%分位数
}
