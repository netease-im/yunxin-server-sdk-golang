package result

import (
	"encoding/json"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
)

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

// ConvertToRtcResult 将YunxinApiResponse转换为RtcResult[T]
func ConvertToRtcResult[T any](apiResponse *core.YunxinApiResponse) (*RtcResult[T], error) {
	httpCode := apiResponse.GetHttpCode()
	code := 0
	requestId := ""
	msg := ""
	var rtcResponse T

	defer func() {
		if r := recover(); r != nil {
			msg = apiResponse.GetData()
		}
	}()

	var jsonObj map[string]interface{}
	if err := json.Unmarshal([]byte(apiResponse.GetData()), &jsonObj); err == nil {
		if codeVal, ok := jsonObj["code"]; ok {
			if codeFloat, ok := codeVal.(float64); ok {
				code = int(codeFloat)
			}
		}

		if requestIdVal, ok := jsonObj["requestId"]; ok {
			if requestIdStr, ok := requestIdVal.(string); ok {
				requestId = requestIdStr
			}
		}

		if errmsgVal, ok := jsonObj["errmsg"]; ok {
			if errmsgStr, ok := errmsgVal.(string); ok {
				msg = errmsgStr
			}
		}

		// 解析响应对象
		if err := json.Unmarshal([]byte(apiResponse.GetData()), &rtcResponse); err != nil {
			msg = apiResponse.GetData()
		}
	} else {
		msg = apiResponse.GetData()
	}

	return NewRtcResult(apiResponse.GetEndpoint(), code, httpCode, requestId, apiResponse.GetTraceId(), msg, rtcResponse), nil
}
