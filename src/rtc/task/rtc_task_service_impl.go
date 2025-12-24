package task

import (
	"encoding/json"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	. "github.com/netease-im/yunxin-server-sdk-golang/src/rtc/result"
)

// rtcTaskService 云端播放任务服务实现
type rtcTaskService struct {
	httpClient core.YunxinApiHttpClient
}

// NewRtcTaskService 创建云端播放任务服务实例
func NewRtcTaskService(httpClient core.YunxinApiHttpClient) RtcTaskService {
	return &rtcTaskService{
		httpClient: httpClient,
	}
}

// CreateTask 创建云端播放任务
func (s *rtcTaskService) CreateTask(request *TaskCreateRequest) (*RtcResult[*TaskCreateResponse], error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.POST, CreateTaskV2, nil, nil, string(body))
	if err != nil {
		return nil, err
	}

	return convertToRtcResult[*TaskCreateResponse](apiResponse)
}

// UpdateTaskV2 V2版本更新云端播放任务
func (s *rtcTaskService) UpdateTaskV2(request *TaskUpdateRequestV2) (*RtcResult[*TaskUpdateResponse], error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.POST, UpdateTaskV2, nil, nil, string(body))
	if err != nil {
		return nil, err
	}

	return convertToRtcResult[*TaskUpdateResponse](apiResponse)
}

// UpdateTaskV3 V3版本更新云端播放任务
func (s *rtcTaskService) UpdateTaskV3(request *TaskUpdateRequestV3) (*RtcResult[*TaskUpdateResponse], error) {
	queryString := map[string]string{
		"cname": request.Cname,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.POST, UpdateTaskV3, nil, queryString, string(body))
	if err != nil {
		return nil, err
	}

	return convertToRtcResult[*TaskUpdateResponse](apiResponse)
}

// QueryTaskV2 V2版本查询云端播放任务
func (s *rtcTaskService) QueryTaskV2(request *TaskQueryRequestV2) (*RtcResult[*TaskQueryResponse], error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.POST, QueryTaskV2, nil, nil, string(body))
	if err != nil {
		return nil, err
	}

	return convertToRtcResult[*TaskQueryResponse](apiResponse)
}

// QueryTaskV3 V3版本查询云端播放任务
func (s *rtcTaskService) QueryTaskV3(request *TaskQueryRequestV3) (*RtcResult[*TaskQueryResponse], error) {
	queryString := map[string]string{
		"cname": request.Cname,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.POST, QueryTaskV3, nil, queryString, string(body))
	if err != nil {
		return nil, err
	}

	return convertToRtcResult[*TaskQueryResponse](apiResponse)
}

// PauseTaskV2 V2版本暂停云端播放任务
func (s *rtcTaskService) PauseTaskV2(request *TaskPauseRequestV2) (*RtcResult[*TaskPauseResponse], error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.POST, PauseTaskV2, nil, nil, string(body))
	if err != nil {
		return nil, err
	}

	return convertToRtcResult[*TaskPauseResponse](apiResponse)
}

// PauseTaskV3 V3版本暂停云端播放任务
func (s *rtcTaskService) PauseTaskV3(request *TaskPauseRequestV3) (*RtcResult[*TaskPauseResponse], error) {
	queryString := map[string]string{
		"cname": request.Cname,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.POST, PauseTaskV3, nil, queryString, string(body))
	if err != nil {
		return nil, err
	}

	return convertToRtcResult[*TaskPauseResponse](apiResponse)
}

// ResumeTaskV2 V2版本恢复云端播放任务
func (s *rtcTaskService) ResumeTaskV2(request *TaskResumeRequestV2) (*RtcResult[*TaskResumeResponse], error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.POST, ResumeTaskV2, nil, nil, string(body))
	if err != nil {
		return nil, err
	}

	return convertToRtcResult[*TaskResumeResponse](apiResponse)
}

// ResumeTaskV3 V3版本恢复云端播放任务
func (s *rtcTaskService) ResumeTaskV3(request *TaskResumeRequestV3) (*RtcResult[*TaskResumeResponse], error) {
	queryString := map[string]string{
		"cname": request.Cname,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.POST, ResumeTaskV3, nil, queryString, string(body))
	if err != nil {
		return nil, err
	}

	return convertToRtcResult[*TaskResumeResponse](apiResponse)
}

// DestroyTaskV2 V2版本销毁云端播放任务
func (s *rtcTaskService) DestroyTaskV2(request *TaskDestroyRequestV2) (*RtcResult[*TaskDestroyResponse], error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.POST, DestroyTaskV2, nil, nil, string(body))
	if err != nil {
		return nil, err
	}

	return convertToRtcResult[*TaskDestroyResponse](apiResponse)
}

// DestroyTaskV3 V3版本销毁云端播放任务
func (s *rtcTaskService) DestroyTaskV3(request *TaskDestroyRequestV3) (*RtcResult[*TaskDestroyResponse], error) {
	queryString := map[string]string{
		"cname": request.Cname,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.POST, DestroyTaskV3, nil, queryString, string(body))
	if err != nil {
		return nil, err
	}

	return convertToRtcResult[*TaskDestroyResponse](apiResponse)
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
