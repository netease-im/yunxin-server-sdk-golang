package result

import "github.com/netease-im/yunxin-server-sdk-golang/src/core"

// RtcResult RTC结果类型
type RtcResult[T any] struct {
	*core.Result[T]
	HttpCode  int    `json:"httpCode"`  // HTTP状态码
	RequestId string `json:"requestId"` // 请求ID
}

// NewRtcResult 创建RTC结果
func NewRtcResult[T any](endpoint string, code int, httpCode int, requestId string, traceId string, msg string, response T) *RtcResult[T] {
	return &RtcResult[T]{
		Result:    core.NewResult(endpoint, code, traceId, msg, response),
		HttpCode:  httpCode,
		RequestId: requestId,
	}
}

// IsSuccess 判断是否成功
func (r *RtcResult[T]) IsSuccess() bool {
	return (r.GetCode() == 200 || r.GetCode() == 0) && r.HttpCode == 200
}

// GetHttpCode 获取HTTP状态码
func (r *RtcResult[T]) GetHttpCode() int {
	return r.HttpCode
}

// GetRequestId 获取请求ID
func (r *RtcResult[T]) GetRequestId() string {
	return r.RequestId
}
