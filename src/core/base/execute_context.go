package base

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/trace"
)

// ExecuteContext 执行上下文 - 实现ExecuteContext接口
type ExecuteContext struct {
	Biz         BizName
	Endpoint    string
	HttpMethod  http.HttpMethod
	ContextType http.ContextType
	ApiVersion  trace.ApiVersion
	Uri         string
	Path        string
	QueryParams map[string]string
	Data        string
	TraceId     string
}

// NewExecuteContext 创建执行上下文
func NewExecuteContext(biz BizName, endpoint string, httpMethod http.HttpMethod, contextType http.ContextType,
	apiVersion trace.ApiVersion, uri string, path string, queryParams map[string]string, data string, traceId string) ExecuteContext {
	return ExecuteContext{
		Biz:         biz,
		Endpoint:    endpoint,
		HttpMethod:  httpMethod,
		ContextType: contextType,
		ApiVersion:  apiVersion,
		Uri:         uri,
		Path:        path,
		QueryParams: queryParams,
		Data:        data,
		TraceId:     traceId,
	}
}
