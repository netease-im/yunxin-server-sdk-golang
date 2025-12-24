package user

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
)

// MeetingResult Meeting结果 - 移至user包避免循环依赖
type MeetingResult[T any] struct {
	*core.Result[T]
	HttpCode  int    `json:"http_code"`  // HTTP状态码
	RequestId string `json:"request_id"` // 请求ID
}

// NewMeetingResult 创建Meeting结果
func NewMeetingResult[T any](endpoint string, code int, httpCode int, requestId string, traceId string, msg string, response T) *MeetingResult[T] {
	return &MeetingResult[T]{
		Result:    core.NewResult(endpoint, code, traceId, msg, response),
		HttpCode:  httpCode,
		RequestId: requestId,
	}
}

// IsSuccess 判断是否成功
func (r *MeetingResult[T]) IsSuccess() bool {
	return (r.GetCode() == 200 || r.GetCode() == 0) && r.HttpCode == 200
}

// GetHttpCode 获取HTTP状态码
func (r *MeetingResult[T]) GetHttpCode() int {
	return r.HttpCode
}

// GetRequestId 获取请求ID
func (r *MeetingResult[T]) GetRequestId() string {
	return r.RequestId
}

// MeetingUserService 会议用户服务接口
type MeetingUserService interface {
	// CreateAccount 创建会议账号
	CreateAccount(req *CreateMeetingAccountRequest) (*MeetingResult[*CreateMeetingAccountResponse], error)
}
