package user

import (
	"encoding/json"
	"fmt"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
)

// meetingUserService 会议用户服务实现
type meetingUserService struct {
	httpClient core.YunxinApiHttpClient
}

// NewMeetingUserService 创建会议用户服务
func NewMeetingUserService(httpClient core.YunxinApiHttpClient) MeetingUserService {
	return &meetingUserService{
		httpClient: httpClient,
	}
}

// CreateAccount 创建会议账号
func (s *meetingUserService) CreateAccount(req *CreateMeetingAccountRequest) (*MeetingResult[*CreateMeetingAccountResponse], error) {
	// 序列化请求体
	requestData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	// 使用HTTP客户端发送POST请求到Meeting创建账号API
	apiResponse, err := s.httpClient.ExecuteJson(http.POST, CreateUser, nil, nil, string(requestData))
	if err != nil {
		return nil, err
	}

	// 转换响应为MeetingResult类型
	return convertToMeetingResult[*CreateMeetingAccountResponse](apiResponse)
}

// convertToMeetingResult 将YunxinApiResponse转换为MeetingResult[T]
func convertToMeetingResult[T any](apiResponse *core.YunxinApiResponse) (*MeetingResult[T], error) {
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

	// 提取requestId字段
	requestId := ""
	if requestIdValue, exists := responseData["requestId"]; exists {
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

	// 构建MeetingResult对象
	return NewMeetingResult(
		apiResponse.GetEndpoint(),
		code,
		apiResponse.GetHttpCode(),
		requestId,
		apiResponse.GetTraceId(),
		msg,
		result,
	), nil
}
