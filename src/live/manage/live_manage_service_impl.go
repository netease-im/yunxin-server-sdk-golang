package manage

import (
	"encoding/json"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/util"
)

// LiveManageServiceImpl 直播管理服务实现
type LiveManageServiceImpl struct {
	httpClient core.YunxinApiHttpClient
}

// NewLiveManageService 创建直播管理服务
func NewLiveManageService(httpClient core.YunxinApiHttpClient) LiveManageService {
	return &LiveManageServiceImpl{
		httpClient: httpClient,
	}
}

// CreateChannel 创建直播频道
func (s *LiveManageServiceImpl) CreateChannel(req *LiveCreateChannelRequest) (*core.Result[*LiveCreateChannelResponse], error) {
	// 序列化请求体
	requestData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	// 使用HTTP客户端发送POST请求到Live创建频道API
	apiResponse, err := s.httpClient.ExecuteJson(http.POST, CreateChannel, nil, nil, string(requestData))
	if err != nil {
		return nil, err
	}

	// 转换响应为Result类型
	return util.ConvertWithClass[*LiveCreateChannelResponse](apiResponse)
}
