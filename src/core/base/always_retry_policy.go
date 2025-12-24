package base

// AlwaysRetryPolicy 总是重试策略
type AlwaysRetryPolicy struct {
	maxRetry int
}

// NewAlwaysRetryPolicy 创建总是重试策略
func NewAlwaysRetryPolicy(maxRetry int) *AlwaysRetryPolicy {
	return &AlwaysRetryPolicy{
		maxRetry: maxRetry,
	}
}

// MaxRetry 最大重试次数
func (a *AlwaysRetryPolicy) MaxRetry() int {
	return a.maxRetry
}

// RetryInterval 重试间隔
func (a *AlwaysRetryPolicy) RetryInterval(retryContext ExecuteContext, retry int) int64 {
	return GetRetryInterval(retry)
}

// OnError 错误重试策略
func (a *AlwaysRetryPolicy) OnError(retryContext ExecuteContext, retry int, err error) *RetryAction {
	return RetryNext
}
