package livewallsolution

import (
	"encoding/json"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	. "github.com/netease-im/yunxin-server-sdk-golang/src/rtc/result"
)

// livewallSolutionServiceImpl 安全通服务实现
type livewallSolutionServiceImpl struct {
	httpClient core.YunxinApiHttpClient
}

// NewLivewallSolutionService 创建安全通服务实例
func NewLivewallSolutionService(httpClient core.YunxinApiHttpClient) LivewallSolutionService {
	return &livewallSolutionServiceImpl{
		httpClient: httpClient,
	}
}

// CreateTask 创建安全通审核任务
func (s *livewallSolutionServiceImpl) CreateTask(req *CreateTaskRequest) (*RtcResult[*CreateTaskResponse], error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.POST, CreateTask, nil, nil, string(body))
	if err != nil {
		return nil, err
	}

	return convertToRtcResult[*CreateTaskResponse](apiResponse)
}

// StopTask 停止安全通审核任务
func (s *livewallSolutionServiceImpl) StopTask(req *StopTaskRequest) (*RtcResult[*StopTaskResponse], error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.POST, StopTask, nil, nil, string(body))
	if err != nil {
		return nil, err
	}

	return convertToRtcResult[*StopTaskResponse](apiResponse)
}

// QueryImage 查询审核视频截图
func (s *livewallSolutionServiceImpl) QueryImage(req *QueryImageRequest) (*RtcResult[*QueryImageResponse], error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.POST, QueryImage, nil, nil, string(body))
	if err != nil {
		return nil, err
	}

	return convertToRtcResult[*QueryImageResponse](apiResponse)
}

// QueryAudioTask 查询审核音频断句
func (s *livewallSolutionServiceImpl) QueryAudioTask(req *QueryAudioTaskRequest) (*RtcResult[*QueryAudioTaskResponse], error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.POST, QueryAudioTask, nil, nil, string(body))
	if err != nil {
		return nil, err
	}

	return convertToRtcResult[*QueryAudioTaskResponse](apiResponse)
}

// convertToRtcResult 将YunxinApiResponse转换为RtcResult[T]
func convertToRtcResult[T any](apiResponse *core.YunxinApiResponse) (*RtcResult[T], error) {
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
