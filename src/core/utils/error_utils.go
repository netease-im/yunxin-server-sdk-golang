package utils

import (
	"context"
	"errors"
	"net"
	"strings"
	"syscall"
)

// IsConnectError 判断是否为连接错误
func IsConnectError(err error) bool {
	if err == nil {
		return false
	}

	// 检查context错误
	if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) {
		return true
	}

	// 检查网络错误
	var netErr net.Error
	if errors.As(err, &netErr) && (netErr.Timeout() || netErr.Temporary()) {
		return true
	}

	// 检查DNS错误
	var dnsErr *net.DNSError
	if errors.As(err, &dnsErr) {
		return true
	}

	// 检查连接被拒绝错误
	var opErr *net.OpError
	if errors.As(err, &opErr) {
		if opErr.Op == "dial" {
			return true
		}
	}

	// 检查底层系统错误
	var syscallErr syscall.Errno
	if errors.As(err, &syscallErr) {
		if errors.Is(syscallErr, syscall.ECONNREFUSED) ||
			errors.Is(syscallErr, syscall.ECONNRESET) ||
			errors.Is(syscallErr, syscall.ETIMEDOUT) {
			return true
		}
	}

	// 检查字符串错误信息
	errStr := strings.ToLower(err.Error())
	return strings.Contains(errStr, "connection refused") ||
		strings.Contains(errStr, "connection reset") ||
		strings.Contains(errStr, "timeout") ||
		strings.Contains(errStr, "no such host") ||
		strings.Contains(errStr, "network unreachable")
}

// IsTimeoutErr 判断是否为超时错误
func IsTimeoutErr(err error) bool {
	if err == nil {
		return false
	}

	// 检查context超时
	if errors.Is(err, context.DeadlineExceeded) {
		return true
	}

	// 检查网络超时
	var netErr net.Error
	if errors.As(err, &netErr) && netErr.Timeout() {
		return true
	}

	// 检查系统调用超时
	var syscallErr syscall.Errno
	if errors.As(err, &syscallErr) && errors.Is(syscallErr, syscall.ETIMEDOUT) {
		return true
	}

	// 检查字符串错误信息
	errStr := strings.ToLower(err.Error())
	return strings.Contains(errStr, "timeout") ||
		strings.Contains(errStr, "deadline exceeded")
}
