package metrics

import (
	"fmt"
	"strings"
)

// PrometheusConverter Prometheus指标转换器
type PrometheusConverter struct{}

// Convert 将统计数据转换为Prometheus格式
func Convert(stats *Stats) string {
	var builder strings.Builder

	// 端点统计
	builder.WriteString("# HELP endpoint Stats\n")
	builder.WriteString("# TYPE endpoint gauge\n")

	bizName := stats.BizName
	for _, endpointStats := range stats.EndpointStatsList {
		endpoint := endpointStats.Endpoint
		result := endpointStats.Result
		count := endpointStats.Count
		avg := endpointStats.Avg
		max := endpointStats.Max
		p50 := endpointStats.P50
		p75 := endpointStats.P75
		p90 := endpointStats.P90
		p99 := endpointStats.P99
		p999 := endpointStats.P999

		builder.WriteString(fmt.Sprintf("yunxin_sdk_endpoint_stats{biz=\"%s\",endpoint=\"%s\",result=\"%s\",type=\"count\"} %d\n",
			bizName, endpoint, result, count))
		builder.WriteString(fmt.Sprintf("yunxin_sdk_endpoint_stats{biz=\"%s\",endpoint=\"%s\",result=\"%s\",type=\"avg\"} %f\n",
			bizName, endpoint, result, avg))
		builder.WriteString(fmt.Sprintf("yunxin_sdk_endpoint_stats{biz=\"%s\",endpoint=\"%s\",result=\"%s\",type=\"max\"} %f\n",
			bizName, endpoint, result, max))
		builder.WriteString(fmt.Sprintf("yunxin_sdk_endpoint_stats{biz=\"%s\",endpoint=\"%s\",result=\"%s\",type=\"p50\"} %f\n",
			bizName, endpoint, result, p50))
		builder.WriteString(fmt.Sprintf("yunxin_sdk_endpoint_stats{biz=\"%s\",endpoint=\"%s\",result=\"%s\",type=\"p75\"} %f\n",
			bizName, endpoint, result, p75))
		builder.WriteString(fmt.Sprintf("yunxin_sdk_endpoint_stats{biz=\"%s\",endpoint=\"%s\",result=\"%s\",type=\"p90\"} %f\n",
			bizName, endpoint, result, p90))
		builder.WriteString(fmt.Sprintf("yunxin_sdk_endpoint_stats{biz=\"%s\",endpoint=\"%s\",result=\"%s\",type=\"p99\"} %f\n",
			bizName, endpoint, result, p99))
		builder.WriteString(fmt.Sprintf("yunxin_sdk_endpoint_stats{biz=\"%s\",endpoint=\"%s\",result=\"%s\",type=\"p999\"} %f\n",
			bizName, endpoint, result, p999))
	}

	// URI统计
	builder.WriteString("# HELP uri Stats\n")
	builder.WriteString("# TYPE uri gauge\n")

	for _, uriStats := range stats.UriStatsList {
		endpoint := uriStats.Endpoint
		method := uriStats.Method
		apiVersion := uriStats.ApiVersion
		contextType := uriStats.ContextType
		uri := uriStats.Uri
		result := uriStats.Result
		count := uriStats.Count
		avg := uriStats.Avg
		max := uriStats.Max
		p50 := uriStats.P50
		p75 := uriStats.P75
		p90 := uriStats.P90
		p99 := uriStats.P99
		p999 := uriStats.P999

		builder.WriteString(fmt.Sprintf("yunxin_sdk_uri_stats{biz=\"%s\",endpoint=\"%s\",method=\"%s\",api_version=\"%s\",context_type=\"%s\",uri=\"%s\",result=\"%s\",type=\"count\"} %d\n",
			bizName, endpoint, method.String(), apiVersion.String(), contextType.String(), uri, result, count))
		builder.WriteString(fmt.Sprintf("yunxin_sdk_uri_stats{biz=\"%s\",endpoint=\"%s\",method=\"%s\",api_version=\"%s\",context_type=\"%s\",uri=\"%s\",result=\"%s\",type=\"avg\"} %f\n",
			bizName, endpoint, method.String(), apiVersion.String(), contextType.String(), uri, result, avg))
		builder.WriteString(fmt.Sprintf("yunxin_sdk_uri_stats{biz=\"%s\",endpoint=\"%s\",method=\"%s\",api_version=\"%s\",context_type=\"%s\",uri=\"%s\",result=\"%s\",type=\"max\"} %f\n",
			bizName, endpoint, method.String(), apiVersion.String(), contextType.String(), uri, result, max))
		builder.WriteString(fmt.Sprintf("yunxin_sdk_uri_stats{biz=\"%s\",endpoint=\"%s\",method=\"%s\",api_version=\"%s\",context_type=\"%s\",uri=\"%s\",result=\"%s\",type=\"p50\"} %f\n",
			bizName, endpoint, method.String(), apiVersion.String(), contextType.String(), uri, result, p50))
		builder.WriteString(fmt.Sprintf("yunxin_sdk_uri_stats{biz=\"%s\",endpoint=\"%s\",method=\"%s\",api_version=\"%s\",context_type=\"%s\",uri=\"%s\",result=\"%s\",type=\"p75\"} %f\n",
			bizName, endpoint, method.String(), apiVersion.String(), contextType.String(), uri, result, p75))
		builder.WriteString(fmt.Sprintf("yunxin_sdk_uri_stats{biz=\"%s\",endpoint=\"%s\",method=\"%s\",api_version=\"%s\",context_type=\"%s\",uri=\"%s\",result=\"%s\",type=\"p90\"} %f\n",
			bizName, endpoint, method.String(), apiVersion.String(), contextType.String(), uri, result, p90))
		builder.WriteString(fmt.Sprintf("yunxin_sdk_uri_stats{biz=\"%s\",endpoint=\"%s\",method=\"%s\",api_version=\"%s\",context_type=\"%s\",uri=\"%s\",result=\"%s\",type=\"p99\"} %f\n",
			bizName, endpoint, method.String(), apiVersion.String(), contextType.String(), uri, result, p99))
		builder.WriteString(fmt.Sprintf("yunxin_sdk_uri_stats{biz=\"%s\",endpoint=\"%s\",method=\"%s\",api_version=\"%s\",context_type=\"%s\",uri=\"%s\",result=\"%s\",type=\"p999\"} %f\n",
			bizName, endpoint, method.String(), apiVersion.String(), contextType.String(), uri, result, p999))
	}

	return builder.String()
}
