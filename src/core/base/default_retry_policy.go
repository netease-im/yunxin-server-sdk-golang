package base

import (
	"errors"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core/utils"
)

// DefaultRetryPolicy 默认重试策略
type DefaultRetryPolicy struct {
	maxRetry   int
	retryOn502 bool
}

var (
	retryPolicy   = DefaultRetryPolicy{maxRetry: 1, retryOn502: true}
	retryOn502    = DefaultRetryPolicy{maxRetry: 1, retryOn502: true}
	notRetryOn502 = DefaultRetryPolicy{maxRetry: 1, retryOn502: false}
)

// NewDefaultRetryPolicyWithParams 创建默认重试策略（带参数）
func NewDefaultRetryPolicyWithParams(maxRetry int, retryOn502 bool) *DefaultRetryPolicy {
	return &DefaultRetryPolicy{
		maxRetry:   maxRetry,
		retryOn502: retryOn502,
	}
}

// NewDefaultRetryPolicy 创建默认重试策略（使用默认参数）
func NewDefaultRetryPolicy(maxRetry int, retryOn502 bool) *DefaultRetryPolicy {
	return &DefaultRetryPolicy{
		maxRetry:   maxRetry,
		retryOn502: retryOn502,
	}
}

// MaxRetry 最大重试次数
func (d *DefaultRetryPolicy) MaxRetry() int {
	return d.maxRetry
}

// RetryInterval 重试间隔
func (d *DefaultRetryPolicy) RetryInterval(retryContext ExecuteContext, retry int) int64 {
	return GetRetryInterval(retry)
}

// OnError 错误重试策略
func (d *DefaultRetryPolicy) OnError(retryContext ExecuteContext, retry int, err error) *RetryAction {
	var httpErr *HttpCodeError
	if errors.As(err, &httpErr) {
		if httpErr.Code == 502 && d.retryOn502 {
			return RetryNext
		}
	}
	if utils.IsConnectError(err) {
		return RetryNext
	}
	return NoRetry
}
