package util

import (
	"encoding/json"
	"fmt"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
)

// ResultUtils 结果转换工具类
type ResultUtils struct{}

// Convert 将YunxinApiResponse转换为Result[T]
func Convert[T any](apiResponse *core.YunxinApiResponse, responseType *T) (*core.Result[T], error) {
	// 解析响应数据
	var responseData map[string]interface{}
	if err := json.Unmarshal([]byte(apiResponse.GetData()), &responseData); err != nil {
		return nil, fmt.Errorf("failed to parse response data: %v", err)
	}

	// 提取code字段
	code := 0
	if codeValue, exists := responseData["code"]; exists {
		if codeFloat, ok := codeValue.(float64); ok {
			code = int(codeFloat)
		}
	}

	// 提取msg字段
	msg := ""
	if msgValue, exists := responseData["msg"]; exists {
		if msgStr, ok := msgValue.(string); ok {
			msg = msgStr
		}
	}

	// 提取data字段并转换为目标类型
	dataDescs := []string{"data", "ret"}
	var result T
	for _, dataDesc := range dataDescs {
		if dataValue, exists := responseData[dataDesc]; exists && dataValue != nil {
			// 将data字段重新序列化再反序列化到目标类型
			dataBytes, err := json.Marshal(dataValue)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal data field: %v", err)
			}

			if err := json.Unmarshal(dataBytes, &result); err != nil {
				return nil, fmt.Errorf("failed to unmarshal to target type: %v", err)
			}
			break
		}
	}

	// 构建Result对象
	return &core.Result[T]{
		Endpoint: apiResponse.GetEndpoint(),
		Code:     code,
		TraceId:  apiResponse.GetTraceId(),
		Msg:      msg,
		Response: result,
	}, nil
}

// ConvertWithClass 将YunxinApiResponse转换为Result[T]的便捷方法
func ConvertWithClass[T any](apiResponse *core.YunxinApiResponse) (*core.Result[T], error) {
	var responseType T
	return Convert(apiResponse, &responseType)
}
