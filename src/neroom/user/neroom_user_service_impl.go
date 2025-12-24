package user

import (
	"encoding/json"
	"fmt"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	. "github.com/netease-im/yunxin-server-sdk-golang/src/neroom/result"
)

// neroomUserService Neroom用户服务实现
type neroomUserService struct {
	httpClient core.YunxinApiHttpClient
}

// NewNeroomUserService 创建Neroom用户服务
func NewNeroomUserService(httpClient core.YunxinApiHttpClient) NeroomUserService {
	return &neroomUserService{
		httpClient: httpClient,
	}
}

// CreateAccount 创建Neroom账号
func (s *neroomUserService) CreateAccount(req *CreateNeroomAccountRequest) (*NeroomResult[*CreateNeroomAccountResponse], error) {
	// 序列化请求体
	requestData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	// 使用HTTP客户端发送POST请求到Neroom创建账号API
	apiResponse, err := s.httpClient.ExecuteJson(http.POST, CreateUser, nil, nil, string(requestData))
	if err != nil {
		return nil, err
	}

	// 转换响应为NeroomResult类型
	return convertToNeroomResult[*CreateNeroomAccountResponse](apiResponse)
}

// convertToNeroomResult 将YunxinApiResponse转换为NeroomResult[T]
func convertToNeroomResult[T any](apiResponse *core.YunxinApiResponse) (*NeroomResult[T], error) {
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

	// 提取ts字段
	ts := int64(0)
	if tsValue, exists := responseData["ts"]; exists {
		if tsFloat, ok := tsValue.(float64); ok {
			ts = int64(tsFloat)
		}
	}

	// 提取cost字段
	cost := ""
	if costValue, exists := responseData["cost"]; exists {
		if costStr, ok := costValue.(string); ok {
			cost = costStr
		}
	}

	// 提取msg字段
	msg := ""
	if msgValue, exists := responseData["msg"]; exists {
		if msgStr, ok := msgValue.(string); ok {
			msg = msgStr
		}
	}

	// 提取requestId字段
	requestId := ""
	if requestIdValue, exists := responseData["request_id"]; exists {
		if requestIdStr, ok := requestIdValue.(string); ok {
			requestId = requestIdStr
		}
	}

	// 提取data字段并转换为目标类型
	var result T
	if dataValue, exists := responseData["data"]; exists && dataValue != nil {
		// 将data字段重新序列化再反序列化到目标类型
		dataBytes, err := json.Marshal(dataValue)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal data field: %v", err)
		}

		if err := json.Unmarshal(dataBytes, &result); err != nil {
			return nil, fmt.Errorf("failed to unmarshal to target type: %v", err)
		}
	}

	// 构建NeroomResult对象
	return NewNeroomResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), requestId, ts, cost, msg, result), nil
}
