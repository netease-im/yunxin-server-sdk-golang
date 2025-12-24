package result

import "github.com/netease-im/yunxin-server-sdk-golang/src/core"

// NeroomResult Neroom结果类型
type NeroomResult[T any] struct {
	*core.Result[T]
	RequestId string `json:"requestId"` // 请求ID
	Ts        int64  `json:"ts"`        // 时间戳
	Cost      string `json:"cost"`      // 耗时
}

// NewNeroomResult 创建Neroom结果
func NewNeroomResult[T any](endpoint string, code int, traceId, requestId string,
	ts int64, cost, msg string, response T) *NeroomResult[T] {
	return &NeroomResult[T]{
		Result:    core.NewResult(endpoint, code, traceId, msg, response),
		RequestId: requestId,
		Ts:        ts,
		Cost:      cost,
	}
}

// IsSuccess 判断是否成功
func (nr *NeroomResult[T]) IsSuccess() bool {
	return nr.Code == 0
}
