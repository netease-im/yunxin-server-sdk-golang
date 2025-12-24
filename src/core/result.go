package core

// Result 结果类型
type Result[T any] struct {
	Endpoint string
	Code     int
	TraceId  string
	Msg      string
	Response T
}

// NewResult 创建结果对象
func NewResult[T any](endpoint string, code int, traceId string, msg string, response T) *Result[T] {
	return &Result[T]{
		Endpoint: endpoint,
		Code:     code,
		TraceId:  traceId,
		Msg:      msg,
		Response: response,
	}
}

// IsSuccess 判断是否成功
func (r *Result[T]) IsSuccess() bool {
	return r.Code == 200
}

// GetCode 获取状态码
func (r *Result[T]) GetCode() int {
	return r.Code
}

// GetEndpoint 获取端点
func (r *Result[T]) GetEndpoint() string {
	return r.Endpoint
}

// GetTraceId 获取跟踪ID
func (r *Result[T]) GetTraceId() string {
	return r.TraceId
}

// GetMsg 获取消息
func (r *Result[T]) GetMsg() string {
	return r.Msg
}

// GetResponse 获取响应
func (r *Result[T]) GetResponse() T {
	return r.Response
}
