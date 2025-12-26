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

	return ConvertToRtcResult[*CreateTaskResponse](apiResponse)
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

	return ConvertToRtcResult[*StopTaskResponse](apiResponse)
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

	return ConvertToRtcResult[*QueryImageResponse](apiResponse)
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

	return ConvertToRtcResult[*QueryAudioTaskResponse](apiResponse)
}
