package meeting

import "github.com/netease-im/yunxin-server-sdk-golang/src/core"

// MeetingResult 会议结果类型
type MeetingResult[T any] struct {
	*core.Result[T]
	RequestId string `json:"requestId"` // 请求ID
	Ts        int64  `json:"ts"`        // 时间戳
	Cost      string `json:"cost"`      // 耗时
}

// NewMeetingResult 创建会议结果
func NewMeetingResult[T any](endpoint string, code int, traceId, requestId string,
	ts int64, cost, msg string, response T) *MeetingResult[T] {
	return &MeetingResult[T]{
		Result:    core.NewResult(endpoint, code, traceId, msg, response),
		RequestId: requestId,
		Ts:        ts,
		Cost:      cost,
	}
}

// IsSuccess 判断是否成功
func (mr *MeetingResult[T]) IsSuccess() bool {
	return mr.Code == 0
}

// GetRequestId 获取请求ID
func (mr *MeetingResult[T]) GetRequestId() string {
	return mr.RequestId
}

// GetTs 获取时间戳
func (mr *MeetingResult[T]) GetTs() int64 {
	return mr.Ts
}

// GetCost 获取耗时
func (mr *MeetingResult[T]) GetCost() string {
	return mr.Cost
}
