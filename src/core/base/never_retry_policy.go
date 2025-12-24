package base

// NeverRetryPolicy 不重试策略
type NeverRetryPolicy struct{}

// NewNeverRetryPolicy 创建不重试策略
func NewNeverRetryPolicy() *NeverRetryPolicy {
	return &NeverRetryPolicy{}
}

// MaxRetry 最大重试次数（返回0表示不重试）
func (n *NeverRetryPolicy) MaxRetry() int {
	return 0
}

// RetryInterval 重试间隔（不重试时返回0）
func (n *NeverRetryPolicy) RetryInterval(retryContext ExecuteContext, retry int) int64 {
	return 0
}

// OnError 错误重试策略（总是返回不重试）
func (n *NeverRetryPolicy) OnError(retryContext ExecuteContext, retry int, err error) *RetryAction {
	return NoRetry
}
