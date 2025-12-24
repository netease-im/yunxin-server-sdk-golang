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
	queryParams map[string]string
	Data        string
	TraceId     string
}

// NewExecuteContext 创建执行上下文
func NewExecuteContext(biz BizName, endpoint string, httpMethod http.HttpMethod,
	contextType http.ContextType, apiVersion trace.ApiVersion, uri string, path string,
	queryParams map[string]string, data string, traceId string) ExecuteContext {
	return ExecuteContext{
		Biz:         biz,
		Endpoint:    endpoint,
		HttpMethod:  httpMethod,
		ContextType: contextType,
		ApiVersion:  apiVersion,
		Uri:         uri,
		Path:        path,
		queryParams: queryParams,
		Data:        data,
		TraceId:     traceId,
	}
}

// GetHttpMethod 获取HTTP方法
func (e *ExecuteContext) GetHttpMethod() http.HttpMethod {
	return e.HttpMethod
}

// GetContextType 获取上下文类型
func (e *ExecuteContext) GetContextType() http.ContextType {
	return e.ContextType
}

// GetApiVersion 获取API版本
func (e *ExecuteContext) GetApiVersion() trace.ApiVersion {
	return e.ApiVersion
}

// GetUri 获取URI
func (e *ExecuteContext) GetUri() string {
	return e.Uri
}

// GetPath 获取路径
func (e *ExecuteContext) GetPath() string {
	return e.Path
}

// GetqueryParams 获取查询字符串
func (e *ExecuteContext) GetqueryParams() map[string]string {
	return e.queryParams
}

// GetData 获取数据
func (e *ExecuteContext) GetData() string {
	return e.Data
}
