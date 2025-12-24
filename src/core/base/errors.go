package base

import (
	"fmt"
)

// YunxinSdkError 云信SDK错误
type YunxinSdkError struct {
	Biz      BizName
	Endpoint string
	TraceId  string
	Err      error
}

func (e *YunxinSdkError) Error() string {
	return fmt.Sprintf("biz:%s, endpoint: %s, trace-id: %s, message: %s",
		e.Biz.Name, e.Endpoint, e.TraceId, e.Err.Error())
}

// NewYunxinSdkError 创建云信SDK错误
func NewYunxinSdkError(biz BizName, endpoint, traceId string, err error) *YunxinSdkError {
	return &YunxinSdkError{
		Biz:      biz,
		Endpoint: endpoint,
		TraceId:  traceId,
		Err:      err,
	}
}

// GetTraceId 获取跟踪ID
func (e *YunxinSdkError) GetTraceId() string {
	return e.TraceId
}

// HttpCodeError HTTP状态码错误
type HttpCodeError struct {
	Biz      BizName
	Endpoint string
	Code     int
	Data     string
}

func (e *HttpCodeError) Error() string {
	return fmt.Sprintf("biz:%s, endpoint: %s, http.code: %d, data: %s",
		e.Biz.Name, e.Endpoint, e.Code, e.Data)
}

// NewHttpCodeError 创建HTTP状态码错误
func NewHttpCodeError(biz BizName, endpoint string, code int, data string) *HttpCodeError {
	return &HttpCodeError{
		Biz:      biz,
		Endpoint: endpoint,
		Code:     code,
		Data:     data,
	}
}

// EndpointFetchError 端点获取错误
type EndpointFetchError struct {
	Message string
	Cause   error
}

func (e *EndpointFetchError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("EndpointFetchError: %s, caused by: %s", e.Message, e.Cause.Error())
	}
	return fmt.Sprintf("EndpointFetchError: %s", e.Message)
}

// NewEndpointFetchErrorWithMessage 创建端点获取错误（带消息）
func NewEndpointFetchErrorWithMessage(message string) *EndpointFetchError {
	return &EndpointFetchError{
		Message: message,
	}
}

// NewEndpointFetchErrorWithCause 创建端点获取错误（带原因）
func NewEndpointFetchErrorWithCause(cause error) *EndpointFetchError {
	return &EndpointFetchError{
		Cause: cause,
	}
}
