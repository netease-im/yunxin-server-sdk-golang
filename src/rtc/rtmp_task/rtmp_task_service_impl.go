package rtmp_task

import (
	"encoding/json"
	"strconv"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	. "github.com/netease-im/yunxin-server-sdk-golang/src/rtc/result"
)

// rtmpTaskService 旁路推流任务服务实现
type rtmpTaskService struct {
	httpClient core.YunxinApiHttpClient
}

// NewRtmpTaskService 创建旁路推流任务服务实例
func NewRtmpTaskService(httpClient core.YunxinApiHttpClient) RtmpTaskService {
	return &rtmpTaskService{
		httpClient: httpClient,
	}
}

// CreateRtmpTaskV2 V2版本创建旁路推流任务
func (s *rtmpTaskService) CreateRtmpTaskV2(request *RtmpTaskCreateRequestV2) (*RtcResult[*RtmpTaskCreateResponse], error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	pathParams := map[string]string{
		"cid": strconv.FormatInt(request.Cid, 10),
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.POST, CreateRtmpTaskV2, pathParams, nil, string(body))
	if err != nil {
		return nil, err
	}

	return ConvertToRtcResult[*RtmpTaskCreateResponse](apiResponse)
}

// CreateRtmpTaskV3 V3版本创建旁路推流任务
func (s *rtmpTaskService) CreateRtmpTaskV3(request *RtmpTaskCreateRequestV3) (*RtcResult[*RtmpTaskCreateResponse], error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	// V3版本: cname作为查询参数
	queryString := map[string]string{
		"cname": request.Cname,
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.POST, CreateRtmpTaskV3, nil, queryString, string(body))
	if err != nil {
		return nil, err
	}

	return ConvertToRtcResult[*RtmpTaskCreateResponse](apiResponse)
}

// UpdateRtmpTaskV2 V2版本更新旁路推流任务
func (s *rtmpTaskService) UpdateRtmpTaskV2(request *RtmpTaskUpdateRequestV2) (*RtcResult[*RtmpTaskUpdateResponse], error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	pathParams := map[string]string{
		"cid": strconv.FormatInt(request.Cid, 10),
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.POST, UpdateRtmpTaskV2, pathParams, nil, string(body))
	if err != nil {
		return nil, err
	}

	return ConvertToRtcResult[*RtmpTaskUpdateResponse](apiResponse)
}

// UpdateRtmpTaskV3 V3版本更新旁路推流任务
func (s *rtmpTaskService) UpdateRtmpTaskV3(request *RtmpTaskUpdateRequestV3) (*RtcResult[*RtmpTaskUpdateResponse], error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	// V3版本: cname作为查询参数
	queryString := map[string]string{
		"cname": request.Cname,
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.POST, UpdateRtmpTaskV3, nil, queryString, string(body))
	if err != nil {
		return nil, err
	}

	return ConvertToRtcResult[*RtmpTaskUpdateResponse](apiResponse)
}

// QueryRtmpTaskV2 V2版本查询指定旁路推流任务
func (s *rtmpTaskService) QueryRtmpTaskV2(request *RtmpTaskQueryRequestV2) (*RtcResult[*RtmpTaskQueryResponse], error) {
	pathParams := map[string]string{
		"cid":    strconv.FormatInt(request.Cid, 10),
		"taskId": request.TaskId,
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.GET, QueryRtmpTaskV2, pathParams, nil, "")
	if err != nil {
		return nil, err
	}

	return ConvertToRtcResult[*RtmpTaskQueryResponse](apiResponse)
}

// QueryRtmpTaskV3 V3版本查询指定旁路推流任务
func (s *rtmpTaskService) QueryRtmpTaskV3(request *RtmpTaskQueryRequestV3) (*RtcResult[*RtmpTaskQueryResponse], error) {
	// V3版本: cname和taskId作为查询参数
	queryString := map[string]string{
		"cname":  request.Cname,
		"taskId": request.TaskId,
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.GET, QueryRtmpTaskV3, nil, queryString, "")
	if err != nil {
		return nil, err
	}

	return ConvertToRtcResult[*RtmpTaskQueryResponse](apiResponse)
}

// QueryAllRtmpTasksV2 V2版本查询所有旁路推流任务
func (s *rtmpTaskService) QueryAllRtmpTasksV2(request *RtmpTaskQueryAllRequestV2) (*RtcResult[*RtmpTaskQueryAllResponse], error) {
	pathParams := map[string]string{
		"cid": strconv.FormatInt(request.Cid, 10),
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.GET, QueryAllRtmpTasksV2, pathParams, nil, "")
	if err != nil {
		return nil, err
	}

	return ConvertToRtcResult[*RtmpTaskQueryAllResponse](apiResponse)
}

// QueryAllRtmpTasksV3 V3版本查询所有旁路推流任务
func (s *rtmpTaskService) QueryAllRtmpTasksV3(request *RtmpTaskQueryAllRequestV3) (*RtcResult[*RtmpTaskQueryAllResponse], error) {
	// V3版本: cname作为查询参数
	queryString := map[string]string{
		"cname": request.Cname,
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.GET, QueryAllRtmpTasksV3, nil, queryString, "")
	if err != nil {
		return nil, err
	}

	return ConvertToRtcResult[*RtmpTaskQueryAllResponse](apiResponse)
}

// DeleteRtmpTaskV2 V2版本停止旁路推流任务
func (s *rtmpTaskService) DeleteRtmpTaskV2(request *RtmpTaskDeleteRequestV2) (*RtcResult[*RtmpTaskDeleteResponse], error) {
	pathParams := map[string]string{
		"cid": strconv.FormatInt(request.Cid, 10),
	}

	requestData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.DELETE, DeleteRtmpTaskV2, pathParams, nil, string(requestData))
	if err != nil {
		return nil, err
	}

	return ConvertToRtcResult[*RtmpTaskDeleteResponse](apiResponse)
}

// DeleteRtmpTaskV3 V3版本停止旁路推流任务
func (s *rtmpTaskService) DeleteRtmpTaskV3(request *RtmpTaskDeleteRequestV3) (*RtcResult[*RtmpTaskDeleteResponse], error) {
	// V3版本: cname作为查询参数
	queryString := map[string]string{
		"cname": request.Cname,
	}

	requestData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.DELETE, DeleteRtmpTaskV3, nil, queryString, string(requestData))
	if err != nil {
		return nil, err
	}

	return ConvertToRtcResult[*RtmpTaskDeleteResponse](apiResponse)
}
