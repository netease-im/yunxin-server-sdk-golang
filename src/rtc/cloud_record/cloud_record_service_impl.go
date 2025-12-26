package cloud_record

import (
	"encoding/json"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	. "github.com/netease-im/yunxin-server-sdk-golang/src/rtc/result"
)

// cloudRecordService 云端录制服务实现
type cloudRecordService struct {
	httpClient core.YunxinApiHttpClient
}

// NewCloudRecordService 创建云端录制服务实例
func NewCloudRecordService(httpClient core.YunxinApiHttpClient) CloudRecordService {
	return &cloudRecordService{
		httpClient: httpClient,
	}
}

// CreateCloudRecordV2 V2版本创建云端录制任务
func (s *cloudRecordService) CreateCloudRecordV2(request *CloudRecordCreateRequestV2) (*RtcResult[*CloudRecordCreateResponse], error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.POST, CreateCloudRecordV2, nil, nil, string(body))
	if err != nil {
		return nil, err
	}

	return ConvertToRtcResult[*CloudRecordCreateResponse](apiResponse)
}

// CreateCloudRecordV3 V3版本创建云端录制任务
func (s *cloudRecordService) CreateCloudRecordV3(request *CloudRecordCreateRequestV3) (*RtcResult[*CloudRecordCreateResponse], error) {
	queryString := map[string]string{
		"cname": request.Cname,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.POST, CreateCloudRecordV3, nil, queryString, string(body))
	if err != nil {
		return nil, err
	}

	return ConvertToRtcResult[*CloudRecordCreateResponse](apiResponse)
}

// UpdateLayoutV2 V2版本更新录制布局
func (s *cloudRecordService) UpdateLayoutV2(request *CloudRecordUpdateLayoutRequestV2) (*RtcResult[*CloudRecordUpdateLayoutResponse], error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.POST, UpdateLayoutV2, nil, nil, string(body))
	if err != nil {
		return nil, err
	}

	return ConvertToRtcResult[*CloudRecordUpdateLayoutResponse](apiResponse)
}

// UpdateLayoutV3 V3版本更新录制布局
func (s *cloudRecordService) UpdateLayoutV3(request *CloudRecordUpdateLayoutRequestV3) (*RtcResult[*CloudRecordUpdateLayoutResponse], error) {
	queryString := map[string]string{
		"cname": request.Cname,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.POST, UpdateLayoutV3, nil, queryString, string(body))
	if err != nil {
		return nil, err
	}

	return ConvertToRtcResult[*CloudRecordUpdateLayoutResponse](apiResponse)
}

// QueryTaskV2 V2版本查询云端录制任务
func (s *cloudRecordService) QueryTaskV2(request *CloudRecordQueryTaskRequestV2) (*RtcResult[*CloudRecordQueryTaskResponse], error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.POST, QueryTaskV2, nil, nil, string(body))
	if err != nil {
		return nil, err
	}

	return ConvertToRtcResult[*CloudRecordQueryTaskResponse](apiResponse)
}

// QueryTaskV3 V3版本查询云端录制任务
func (s *cloudRecordService) QueryTaskV3(request *CloudRecordQueryTaskRequestV3) (*RtcResult[*CloudRecordQueryTaskResponse], error) {
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

	return ConvertToRtcResult[*CloudRecordQueryTaskResponse](apiResponse)
}

// UpdateSubscriptionV2 V2版本更新订阅名单
func (s *cloudRecordService) UpdateSubscriptionV2(request *CloudRecordUpdateSubscriptionRequestV2) (*RtcResult[*CloudRecordUpdateSubscriptionResponse], error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.POST, UpdateSubscriptionV2, nil, nil, string(body))
	if err != nil {
		return nil, err
	}

	return ConvertToRtcResult[*CloudRecordUpdateSubscriptionResponse](apiResponse)
}

// UpdateSubscriptionV3 V3版本更新订阅名单
func (s *cloudRecordService) UpdateSubscriptionV3(request *CloudRecordUpdateSubscriptionRequestV3) (*RtcResult[*CloudRecordUpdateSubscriptionResponse], error) {
	queryString := map[string]string{
		"cname": request.Cname,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.POST, UpdateSubscriptionV3, nil, queryString, string(body))
	if err != nil {
		return nil, err
	}

	return ConvertToRtcResult[*CloudRecordUpdateSubscriptionResponse](apiResponse)
}

// StopRecordV2 V2版本停止录制任务
func (s *cloudRecordService) StopRecordV2(request *CloudRecordStopRequestV2) (*RtcResult[*CloudRecordStopResponse], error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.POST, StopRecordV2, nil, nil, string(body))
	if err != nil {
		return nil, err
	}

	return ConvertToRtcResult[*CloudRecordStopResponse](apiResponse)
}

// StopRecordV3 V3版本停止录制任务
func (s *cloudRecordService) StopRecordV3(request *CloudRecordStopRequestV3) (*RtcResult[*CloudRecordStopResponse], error) {
	queryString := map[string]string{
		"cname": request.Cname,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.POST, StopRecordV3, nil, queryString, string(body))
	if err != nil {
		return nil, err
	}

	return ConvertToRtcResult[*CloudRecordStopResponse](apiResponse)
}
