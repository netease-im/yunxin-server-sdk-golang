package base

type RetryPolicy interface {
	// MaxRetry 最大重试次数
	MaxRetry() int

	// RetryInterval 重试间隔(毫秒)
	RetryInterval(retryContext ExecuteContext, retry int) int64

	// OnError 错误处理策略
	OnError(retryContext ExecuteContext, retry int, err error) *RetryAction
}

// RetryAction 重试动作
type RetryAction struct {
	retry        bool
	nextEndpoint bool
}

var (
	NoRetry      = &RetryAction{retry: false, nextEndpoint: false}
	RetryCurrent = &RetryAction{retry: true, nextEndpoint: false}
	RetryNext    = &RetryAction{retry: true, nextEndpoint: true}
)

var intervalArray = []int64{0, 5, 10, 50, 100, 500, 1000, 5000}

// GetRetryInterval 获取重试间隔
func GetRetryInterval(retry int) int64 {
	if retry >= len(intervalArray) {
		return intervalArray[len(intervalArray)-1]
	}
	return intervalArray[retry]
}

// IsRetry 是否重试
func (r *RetryAction) IsRetry() bool {
	return r.retry
}

// IsNextEndpoint 是否切换端点
func (r *RetryAction) IsNextEndpoint() bool {
	return r.nextEndpoint
}
